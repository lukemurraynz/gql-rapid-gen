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
	Fields []*parser.ParsedField
	Input  bool
}

func (p *Plugin) Generate(schema *parser.Schema, output *gen.Output) error {

	for _, o := range schema.Objects {
		if o.HasDirective("go_ignore") {
			continue
		}

		fields := make([]*parser.ParsedField, 0, len(o.Fields))
		for _, f := range o.Fields {
			if f.HasDirective("go_ignore") {
				continue
			}
			fields = append(fields, f)
		}
		if len(fields) == 0 {
			return fmt.Errorf("object '%s' has no valid fields", o.Name)
		}

		rendered, err := gen.ExecuteTemplate("plugins/go_objects/templates/struct.tmpl", data{
			Object: o,
			Fields: fields,
			Input:  false,
		})
		if err != nil {
			return fmt.Errorf("failed rendering Object %s: %w", o.Name, err)
		}

		of, err := output.AppendOrCreate(gen.GO_DATA_GEN, util.DashCase(o.Name), rendered)
		if err != nil {
			return fmt.Errorf("failed appending Object %s: %w", o.Name, err)
		}
		of.AddExtraData("fmt")
	}

	for _, o := range schema.InputObjects {
		if o.HasDirective("go_ignore") {
			continue
		}

		fields := make([]*parser.ParsedField, 0, len(o.Fields))
		for _, f := range o.Fields {
			if f.HasDirective("go_ignore") {
				continue
			}
			fields = append(fields, f)
		}

		rendered, err := gen.ExecuteTemplate("plugins/go_objects/templates/struct.tmpl", data{
			Object: o,
			Fields: fields,
			Input:  true,
		})
		if err != nil {
			return fmt.Errorf("failed rendering Input Object %s: %w", o.Name, err)
		}

		of, err := output.AppendOrCreate(gen.GO_DATA_GEN, util.DashCase(o.Name), rendered)
		if err != nil {
			return fmt.Errorf("failed appending Input Object %s: %w", o.Name, err)
		}
		of.AddExtraData("fmt")
	}

	return nil
}
