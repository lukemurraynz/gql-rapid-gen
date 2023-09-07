// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package go_enum

import (
	"fmt"
	"github.com/mjdrgn/gql-rapid-gen/gen"
	"github.com/mjdrgn/gql-rapid-gen/parser"
	"github.com/mjdrgn/gql-rapid-gen/util"
)

type data struct {
	Enum *parser.ParsedEnum
}

func (p *Plugin) Generate(schema *parser.Schema, output *gen.Output) error {

	for _, o := range schema.Enums {
		rendered, err := gen.ExecuteTemplate("plugins/go_enum/templates/enum.tmpl", data{
			Enum: o,
		})
		if err != nil {
			return fmt.Errorf("failed rendering Enum %s: %w", o.Name, err)
		}

		of, err := output.AppendOrCreate(gen.GO_DATA_GEN, util.DashCase(o.Name), rendered)
		if err != nil {
			return fmt.Errorf("failed appending Enum %s: %w", o.Name, err)
		}
		of.AddExtraData("fmt")
	}

	return nil
}
