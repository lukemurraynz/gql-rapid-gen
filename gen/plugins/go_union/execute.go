// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package go_union

import (
	"fmt"
	"github.com/mjdrgn/gql-rapid-gen/gen"
	"github.com/mjdrgn/gql-rapid-gen/parser"
	"github.com/mjdrgn/gql-rapid-gen/util"
)

type data struct {
	Union *parser.ParsedUnion
}

func (p *Plugin) Generate(schema *parser.Schema, output *gen.Output) error {

	for _, o := range schema.Unions {
		d := data{
			Union: o,
		}

		rendered, err := gen.ExecuteTemplate("plugins/go_union/templates/template.tmpl", d)
		if err != nil {
			return fmt.Errorf("failed rendering Union template %s: %w", o.Name, err)
		}

		_, err = output.AppendOrCreate(gen.GO_DATA_GEN, util.DashCase(o.Name), rendered)
		if err != nil {
			return fmt.Errorf("failed appending Union template %s: %w", o.Name, err)
		}

		rendered, err = gen.ExecuteTemplate("plugins/go_union/templates/skel.tmpl", d)
		if err != nil {
			return fmt.Errorf("failed rendering Union skel %s: %w", o.Name, err)
		}

		_, err = output.AppendOrCreate(gen.GO_DATA_SKEL, util.DashCase(o.Name), rendered)
		if err != nil {
			return fmt.Errorf("failed appending Union skel %s: %w", o.Name, err)
		}
	}

	return nil
}
