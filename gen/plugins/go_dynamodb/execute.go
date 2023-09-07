// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package go_dynamodb

import (
	"fmt"
	"github.com/mjdrgn/gql-rapid-gen/gen"
	"github.com/mjdrgn/gql-rapid-gen/parser"
	"github.com/mjdrgn/gql-rapid-gen/util"
)

type data struct {
	Object    *parser.ParsedObject
	HashKey   *parser.ParsedField
	SortKey   *parser.ParsedField
	HasSort   bool
	TableName string
	Fields    []*parser.ParsedField
	GSIs      []gsiData
}

type gsiData struct {
	Name      string
	NameTitle string
	HashKey   *parser.ParsedField
	SortKey   *parser.ParsedField
	HasSort   bool
}

func (p *Plugin) Generate(schema *parser.Schema, output *gen.Output) error {

	for _, o := range schema.Objects {
		if o.HasDirective("go_ignore") {
			continue
		}

		dynamo := o.SingleDirective("dynamodb")
		if dynamo == nil {
			continue
		}

		hashKey := o.Field(dynamo.Arg("hash_key"))
		sortKey := o.Field(dynamo.Arg("sort_key"))

		var GSIs []gsiData
		for _, v := range o.Directives["dynamodb_gsi"] {
			gsiHashKey := o.Field(v.Arg("hash_key"))
			gsiSortKey := o.Field(v.Arg("sort_key"))
			if gsiHashKey == nil {
				return fmt.Errorf("GSI '%s' refers to hash_key '%s' which is nil", v.Name, v.Arg("hash_key"))
			}
			if gsiSortKey == nil && v.HasArg("sort_key") {
				return fmt.Errorf("GSI '%s' refers to sort_key '%s' which is nil", v.Name, v.Arg("sort_key"))
			}
			GSIs = append(GSIs, gsiData{
				Name:      v.Arg("name"),
				NameTitle: util.TitleCase(v.Arg("name")),
				HashKey:   gsiHashKey,
				SortKey:   gsiSortKey,
				HasSort:   v.HasArg("sort_key"),
			})
		}

		{
			rendered, err := gen.ExecuteTemplate("plugins/go_dynamodb/templates/access.tmpl", data{
				Object:    o,
				HashKey:   hashKey,
				SortKey:   sortKey,
				HasSort:   dynamo.HasArg("sort_key"),
				TableName: dynamo.Arg("name"),
				Fields:    o.Fields,
				GSIs:      GSIs,
			})
			if err != nil {
				return fmt.Errorf("failed rendering access: %w", err)
			}

			_, err = output.AppendOrCreate(gen.GO_DATA_GEN, o.NameDash(), rendered)
			if err != nil {
				return fmt.Errorf("failed appending access: %w", err)
			}
		}

		{
			rendered, err := gen.ExecuteTemplate("plugins/go_dynamodb/templates/mock.tmpl", data{
				Object:  o,
				HashKey: hashKey,
				SortKey: sortKey,
				HasSort: dynamo.HasArg("sort_key"),
				GSIs:    GSIs,
			})
			if err != nil {
				return fmt.Errorf("failed rendering mock: %w", err)
			}

			_, err = output.AppendOrCreate(gen.GO_DATA_GEN, o.NameDash()+".mock", rendered)
			if err != nil {
				return fmt.Errorf("failed appending mock: %w", err)
			}
		}

		{
			rendered, err := gen.ExecuteTemplate("plugins/go_dynamodb/templates/skel.tmpl", data{
				Object: o,
			})
			if err != nil {
				return fmt.Errorf("failed rendering skeleton: %w", err)
			}

			_, err = output.AppendOrCreate(gen.GO_DATA_SKEL, o.NameDash(), rendered)
			if err != nil {
				return fmt.Errorf("failed appending skeleton: %w", err)
			}
		}

	}

	return nil
}
