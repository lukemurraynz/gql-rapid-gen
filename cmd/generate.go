// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package cmd

import (
	"github.com/mjdrgn/gql-rapid-gen/gen"
	"github.com/mjdrgn/gql-rapid-gen/parser"
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

		outputDir, err := cmd.Flags().GetString("output-dir")
		if err != nil {
			cmd.PrintErrln(err)
			_ = cmd.Help()
			return
		}

		if outputDir == "" && !dryRun && !validate {
			cmd.PrintErrln("Output Directory must be specified with --output-dir unless you are using --dry-run or --validate")
			os.Exit(1)
			return
		} else if !dryRun && !validate {
			// Clean up directory path and confirm access
			outputDir, err = filepath.Abs(outputDir)
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
		}

		enablePlugin, err := cmd.Flags().GetStringSlice("enable-plugin")
		if err != nil {
			cmd.PrintErrln(err)
			_ = cmd.Help()
			return
		}

		disablePlugin, err := cmd.Flags().GetStringSlice("disable-plugin")
		if err != nil {
			cmd.PrintErrln(err)
			_ = cmd.Help()
			return
		}

		if len(enablePlugin) > 0 && len(disablePlugin) > 0 {
			cmd.PrintErrln("You cannot use both --enable-plugin and --disable-plugin at the same time")
			os.Exit(1)
			return
		}

		schemaFiles, err := cmd.Flags().GetStringArray("schema")
		if err != nil {
			cmd.PrintErrln(err)
			_ = cmd.Help()
			return
		}

		if len(schemaFiles) == 0 {
			cmd.PrintErrln("At least one schema file must be specified")
			os.Exit(1)
			return
		}

		schema, err := parser.Parse(schemaFiles)
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

		filteredPlugins := make([]gen.Plugin, 0, len(plugins))

		if len(enablePlugin) > 0 {
			for _, p := range plugins {
				if slices.Contains(enablePlugin, p.Name()) {
					filteredPlugins = append(filteredPlugins, p)
				}
			}
		} else if len(disablePlugin) > 0 {
			for _, p := range plugins {
				if !slices.Contains(disablePlugin, p.Name()) {
					filteredPlugins = append(filteredPlugins, p)
				}
			}
		} else {
			filteredPlugins = plugins
		}

		if len(filteredPlugins) == 0 {
			if len(enablePlugin) > 0 {
				cmd.Println("No available plugins within enabled list")
			} else {
				cmd.Println("No available plugins are removing disabled")
			}
			return
		}

		for _, p := range filteredPlugins {
			log.Printf("Using plugin: %s", p.Name())
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

		cmd.Println("Generation Successful")
		return
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.PersistentFlags().Bool("validate", false, "Validate input only (do not output files)")
	generateCmd.PersistentFlags().Bool("dry-run", false, "Dry run only (do not output files)")
	generateCmd.PersistentFlags().Bool("list", false, "Output list of plugins and exit")

	generateCmd.PersistentFlags().String("output-dir", "", "Output Directory")
	_ = generateCmd.MarkFlagDirname("output-dir")

	generateCmd.PersistentFlags().StringSlice("enable-plugin", nil, "Plugins to enable - if specified, all other plugins will be disabled")
	generateCmd.PersistentFlags().StringSlice("disable-plugin", nil, "Plugins to disable. Conflicts with enable-plugin.")

	generateCmd.PersistentFlags().StringArray("schema", nil, "Schema file to load (can be specified multiple times)")
	_ = generateCmd.MarkFlagFilename("schema", "graphql")

}
