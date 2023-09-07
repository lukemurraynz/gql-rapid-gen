// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package tf_appsync_lambda

import (
	"fmt"
	"github.com/mjdrgn/gql-rapid-gen/gen"
	"github.com/mjdrgn/gql-rapid-gen/parser"
)

type data struct {
	Field     *parser.ParsedField
	Parent    string
	Directive *parser.ParsedDirective
}

func (p *Plugin) Generate(schema *parser.Schema, output *gen.Output) error {

	for _, mut := range schema.Mutation {
		if !mut.HasDirective("appsync_lambda") {
			continue
		}

		dir := mut.SingleDirective("appsync_lambda")

		language := dir.Arg("language")

		switch language {
		case "go":
			rendered, err := gen.ExecuteTemplate("plugins/tf_appsync_lambda/templates/go.tmpl", data{
				Parent:    "Mutation",
				Field:     mut,
				Directive: dir,
			})
			if err != nil {
				return fmt.Errorf("failed rendering Mutation %s: %w", mut.Name, err)
			}

			_, err = output.AppendOrCreate(gen.TF_API_GEN, "lambda-"+mut.NameDash(), rendered)
			if err != nil {
				return fmt.Errorf("failed appending Mutation %s: %w", mut.Name, err)
			}
		}

		rendered, err := gen.ExecuteTemplate("plugins/tf_appsync_lambda/templates/skel.tmpl", data{
			Parent: "Mutation",
			Field:  mut,
		})
		if err != nil {
			return fmt.Errorf("failed rendering Mutation %s: %w", mut.Name, err)
		}

		_, err = output.AppendOrCreate(gen.TF_API_SKEL, "lambda-"+mut.NameDash(), rendered)
		if err != nil {
			return fmt.Errorf("failed appending Mutation %s: %w", mut.Name, err)
		}
	}

	return nil
}
