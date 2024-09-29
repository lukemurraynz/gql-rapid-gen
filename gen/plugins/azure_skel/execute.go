// Copyright (c) 2024 under the MIT license per gql-rapid-gen/LICENSE.MD

package azure_skel

import (
    "embed"
    "fmt"
    "github.com/mjdrgn/gql-rapid-gen/gen"
    "github.com/mjdrgn/gql-rapid-gen/parser"
)

// skeletonFiles includes all plugin skeleton files
//
//go:embed skeleton/*
var skeletonFiles embed.FS

// Generate is a method of the Plugin struct that generates the necessary files
// for the Azure skeleton plugin. It takes a schema and an output as parameters.
func (p *Plugin) Generate(schema *parser.Schema, output *gen.Output) error {
    // Write the plugin skeleton files to the output directory
    err := gen.WritePluginSkeleton(skeletonFiles, output)
    if err != nil {
        // Return an error if writing the skeleton files fails
        return fmt.Errorf("failed writing skeleton: %w", err)
    }
    // Return nil if the operation is successful
    return nil
}