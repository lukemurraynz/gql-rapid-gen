// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gql-rapid-gen",
	Short: "Generate boilerplate and skeleton code from GraphQL Schema",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String("schema", "./schema.graphql", "Primary schema file to process")

	rootCmd.PersistentFlags().StringArray("extra-schema", []string{}, "Additional schema files to parse")
}
