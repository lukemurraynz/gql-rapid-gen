// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package tf_skel

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

func (p *Plugin) Generate(schema *parser.Schema, output *gen.Output) error {

	err := gen.WritePluginSkeleton(skeletonFiles, output)
	if err != nil {
		return fmt.Errorf("failed writing skeleton: %w", err)
	}

	return nil
}
