// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package tf_dynamodb

import (
	"fmt"
	"github.com/mjdrgn/gql-rapid-gen/gen"
	"github.com/mjdrgn/gql-rapid-gen/parser"
	"github.com/mjdrgn/gql-rapid-gen/util"
)

type data struct {
	Object     *parser.ParsedObject
	Dynamo     *parser.ParsedDirective
	HashKey    *parser.ParsedField
	SortKey    *parser.ParsedField
	HasSort    bool
	Attributes map[string]string
	GSIs       []gsiData
}

type gsiData struct {
	Name    string
	HashKey *parser.ParsedField
	SortKey *parser.ParsedField
	HasSort bool
}

func (p *Plugin) Generate(schema *parser.Schema, output *gen.Output) error {

	for _, o := range schema.Objects {
		dynamo := o.SingleDirective("dynamodb")
		if dynamo == nil {
			continue
		}

		hashKey := o.Field(dynamo.Arg("hash_key"))
		sortKey := o.Field(dynamo.Arg("sort_key"))

		atts := make(map[string]string, len(o.Fields))
		atts[hashKey.Name] = hashKey.Type.DynamoType()
		if sortKey != nil {
			atts[sortKey.Name] = sortKey.Type.DynamoType()
		}

		var GSIs []gsiData
		for _, v := range o.Directives["dynamodb_gsi"] {
			gsiHashKey := o.Field(v.Arg("hash_key"))
			gsiSortKey := o.Field(v.Arg("sort_key"))
			GSIs = append(GSIs, gsiData{
				Name:    v.Arg("name"),
				HashKey: gsiHashKey,
				SortKey: gsiSortKey,
				HasSort: v.HasArg("sort_key"),
			})

			atts[gsiHashKey.Name] = gsiHashKey.Type.DynamoType()
			if gsiSortKey != nil {
				atts[gsiSortKey.Name] = gsiSortKey.Type.DynamoType()
			}
		}

		// Fix up attributes which are type M, when they should be S.
		// This may mean we deploy code that doesn't work, if someone puts an Object as a key, however it fixes the more common case of Enums.
		for k, v := range atts {
			if v == "M" {
				atts[k] = "S"
			}
		}

		rendered, err := gen.ExecuteTemplate("plugins/tf_dynamodb/templates/ddb.tmpl", data{
			Object:     o,
			Dynamo:     dynamo,
			HashKey:    hashKey,
			SortKey:    sortKey,
			HasSort:    dynamo.HasArg("sort_key"),
			Attributes: atts,
			GSIs:       GSIs,
		})
		if err != nil {
			return fmt.Errorf("failed rendering Object %s: %w", o.Name, err)
		}

		_, err = output.AppendOrCreate(gen.TF_API_GEN, "ddb-"+util.DashCase(o.Name), rendered)
		if err != nil {
			return fmt.Errorf("failed appending Object %s: %w", o.Name, err)
		}
	}

	return nil
}
