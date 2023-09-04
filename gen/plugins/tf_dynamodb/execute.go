// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package go_objects

import (
	"fmt"
	"github.com/mjdrgn/gql-rapid-gen/gen"
	"github.com/mjdrgn/gql-rapid-gen/parser"
	"github.com/mjdrgn/gql-rapid-gen/util"
)

type data struct {
	Object *parser.ParsedObject
	Input  bool
}

func (p *Plugin) Generate(schema *parser.Schema, output *gen.Output) error {

	for _, o := range schema.Objects {
		if !o.HasDirective("dynamodb") {

		}
		rendered, err := gen.ExecuteTemplate("plugins/tf_dynamodb/templates/ddb.tmpl", data{
			Object: o,
			Input:  false,
		})
		if err != nil {
			return fmt.Errorf("failed rendering Object %s: %w", o.Name, err)
		}

		_, err = output.AppendOrCreate(gen.GO_GEN, util.DashCase(o.Name), rendered)
		if err != nil {
			return fmt.Errorf("failed appending Object %s: %w", o.Name, err)
		}
	}

	return nil
}
