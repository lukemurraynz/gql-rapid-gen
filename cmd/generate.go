// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package cmd

import (
	"github.com/mjdrgn/gql-rapid-gen/gen"
	"github.com/mjdrgn/gql-rapid-gen/parser"
	"github.com/mjdrgn/gql-rapid-gen/state"
	"github.com/mjdrgn/gql-rapid-gen/util"
	"golang.org/x/exp/slices"
	"log"
	"os"
	"path/filepath"
	"sort"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate output code and configuration from GraphQL Schema",
	Long:  `Generate output code and configuration from GraphQL Schema`,
	Run: func(cmd *cobra.Command, args []string) {
		listMode, err := cmd.Flags().GetBool("list")
		if err != nil {
			cmd.PrintErrln(err)
			_ = cmd.Help()
			return
		}
		if listMode {
			// List plugins and exit
			plugins := gen.ListPlugins()
			sort.Strings(plugins)
			cmd.Println("Plugins installed:")
			for _, v := range plugins {
				cmd.Printf("\t- %s\n", v)
			}
			return
		}

		dryRun, err := cmd.Flags().GetBool("dry-run")
		if err != nil {
			cmd.PrintErrln(err)
			_ = cmd.Help()
			return
		}

		validate, err := cmd.Flags().GetBool("validate")
		if err != nil {
			cmd.PrintErrln(err)
			_ = cmd.Help()
			return
		}

		var config = &state.Config{}

		if util.FileExists("gql-rapid-gen.json") {
			config, err = state.LoadConfig("gql-rapid-gen.json")
			if err != nil {
				cmd.PrintErrln(err)
				return
			}
		} else {
			cmd.PrintErrln("gql-rapid-gen.json config file is missing")
			return
		}

		// Clean up directory path and confirm access
		outputDir, err := filepath.Abs(config.OutputDirectory)
		if err != nil {
			cmd.PrintErrln("Output Directory is invalid")
			os.Exit(1)
			return
		}
		_, err = os.ReadDir(outputDir)
		if err != nil {
			cmd.PrintErrln("Output Directory cannot be accessed")
			os.Exit(1)
			return
		}

		schema, err := parser.Parse(config.SchemaFiles)
		if err != nil {
			cmd.PrintErrf("Failed to parse schema: %s\n", err)
			os.Exit(2)
			return
		}

		err = schema.Validate()
		if err != nil {
			cmd.PrintErrf("Schema validation error: \n%s\n", err)
			os.Exit(2)
			return
		}

		if validate {
			cmd.Println("Validation Successful")
			return
		}

		plugins := gen.QualifySchema(schema)
		if len(plugins) == 0 {
			cmd.Println("No generation required for schema")
			return
		}

		taggedPlugins := make([]gen.Plugin, 0, len(plugins))

		// Filter by tags, if required
		if len(config.TagEnable) > 0 {
			for _, p := range plugins {
				enable := true
				for _, t := range p.Tags() {
					if !slices.Contains(config.TagEnable, t) {
						enable = false
						break
					}
				}
				if enable {
					taggedPlugins = append(taggedPlugins, p)
				}
			}
		} else {
			taggedPlugins = plugins
		}

		filteredPlugins := make([]gen.Plugin, 0, len(plugins))

		if len(config.PluginEnable) > 0 {
			for _, p := range taggedPlugins {
				if slices.Contains(config.PluginEnable, p.Name()) {
					filteredPlugins = append(filteredPlugins, p)
				}
			}
		} else if len(config.PluginDisable) > 0 {
			for _, p := range taggedPlugins {
				if !slices.Contains(config.PluginDisable, p.Name()) {
					filteredPlugins = append(filteredPlugins, p)
				}
			}
		} else {
			filteredPlugins = taggedPlugins
		}

		if len(filteredPlugins) == 0 {
			cmd.Println("No plugins able to run")
			return
		}

		for _, p := range filteredPlugins {
			log.Printf("Using plugin: %s", p.Name())
		}

		lock := &state.LockFile{
			Plugins: make([]string, 0, len(filteredPlugins)),
		}
		for _, p := range filteredPlugins {
			lock.Plugins = append(lock.Plugins, p.Name())
		}

		output := &gen.Output{}

		err = gen.ExecuteSchema(filteredPlugins, schema, output)
		if err != nil {
			cmd.PrintErrln(err)
			os.Exit(3)
			return
		}

		if dryRun {
			cmd.Println("Successfully executed all plugins against schema")
			files := output.GetFiles()
			fileNames := make([]string, 0, len(files))
			for k, f := range files {
				_, err = f.Render()
				fileNames = append(fileNames, k)
			}
			slices.Sort(fileNames)
			for _, k := range fileNames {
				cmd.Printf("\t- %s\n", k)
			}
			return
		}

		lock.Files = output.FileNames()

		err = gen.WriteSkeleton(outputDir)
		if err != nil {
			cmd.PrintErrln(err)
			os.Exit(4)
			return
		}

		err = output.Write(outputDir)
		if err != nil {
			cmd.PrintErrln(err)
			os.Exit(5)
			return
		}

		err = config.Save("gql-rapid-gen.json")
		if err != nil {
			cmd.PrintErrln(err)
			os.Exit(6)
			return
		}

		err = lock.Save("gql-rapid-gen.lock.json")
		if err != nil {
			cmd.PrintErrln(err)
			os.Exit(7)
			return
		}

		cmd.Println("Generation Successful")
		return
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.PersistentFlags().Bool("validate", false, "Validate input only (do not output files)")
	generateCmd.PersistentFlags().Bool("dry-run", false, "Dry run only (do not output files)")
	generateCmd.PersistentFlags().Bool("list", false, "Output list of plugins and exit")

}
