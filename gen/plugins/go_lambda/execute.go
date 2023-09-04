// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package go_lambda

import (
	"fmt"
	"github.com/mjdrgn/gql-rapid-gen/gen"
	"github.com/mjdrgn/gql-rapid-gen/parser"
	"github.com/mjdrgn/gql-rapid-gen/util"
)

type data struct {
	Field  *parser.ParsedField
	Parent string
}

func (p *Plugin) Generate(schema *parser.Schema, output *gen.Output) error {

	for _, mut := range schema.Mutation {
		if !mut.HasDirective("appsync_lambda") {
			continue
		}

		rendered, err := gen.ExecuteTemplate("plugins/go_lambda/templates/event.tmpl", data{
			Parent: "Mutation",
			Field:  mut,
		})
		if err != nil {
			return fmt.Errorf("failed rendering Mutation %s: %w", mut.Name, err)
		}

		_, err = output.AppendOrCreate(gen.GO_DATA_GEN, util.DashCase(mut.Name)+"-event", rendered)
		if err != nil {
			return fmt.Errorf("failed appending Mutation %s: %w", mut.Name, err)
		}
	}

	for _, mut := range schema.Query {
		if !mut.HasDirective("appsync_lambda") {
			continue
		}

		rendered, err := gen.ExecuteTemplate("plugins/go_lambda/templates/event.tmpl", data{
			Parent: "Query",
			Field:  mut,
		})
		if err != nil {
			return fmt.Errorf("failed rendering Query %s: %w", mut.Name, err)
		}

		_, err = output.AppendOrCreate(gen.GO_DATA_GEN, util.DashCase(mut.Name)+"-event", rendered)
		if err != nil {
			return fmt.Errorf("failed appending Query %s: %w", mut.Name, err)
		}
	}

	return nil
}
