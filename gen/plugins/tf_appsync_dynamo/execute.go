// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package tf_appsync_dynamo

import (
	"fmt"
	"github.com/mjdrgn/gql-rapid-gen/gen"
	"github.com/mjdrgn/gql-rapid-gen/parser"
	"github.com/mjdrgn/gql-rapid-gen/util"
)

type data struct {
	Object *parser.ParsedObject
}

type scanData struct {
	Object *parser.ParsedObject
	Plural string
}

type listData struct {
	Object    *parser.ParsedObject
	QueryName string
	Field     *parser.ParsedField
	Directive *parser.ParsedDirective
}

type fkeyData struct {
	Object                 *parser.ParsedObject
	Field                  *parser.ParsedField
	Directive              *parser.ParsedDirective
	Query                  bool
	QuerySingle            bool
	Batch                  bool
	FieldSource            string
	FieldForeign           string
	Table                  string
	Index                  string
	AdditionalFieldSource  string
	AdditionalFieldForeign string
}

type crudData struct {
	Object         *parser.ParsedObject
	HashKey        *parser.ParsedField
	SortKey        *parser.ParsedField
	HasSort        bool
	CreateHashType string
	CreateSortType string
}

func (p *Plugin) Generate(schema *parser.Schema, output *gen.Output) error {

	for _, o := range schema.Objects {
		dynamo := o.SingleDirective("dynamodb")
		if dynamo == nil {
			continue
		}

		hashKey := o.Field(dynamo.Arg("hash_key"))
		sortKey := o.Field(dynamo.Arg("sort_key"))

		{
			// Data Source block
			rendered, err := gen.ExecuteTemplate("plugins/tf_appsync_dynamo/templates/datasource.tmpl", data{
				Object: o,
			})
			if err != nil {
				return fmt.Errorf("failed rendering datasource %s: %w", o.Name, err)
			}

			_, err = output.AppendOrCreate(gen.TF_API_GEN, util.DashCase(o.Name), rendered)
			if err != nil {
				return fmt.Errorf("failed appending datasource %s: %w", o.Name, err)
			}
		}

		if d := o.SingleDirective("appsync_scan"); d != nil {
			// Scan block
			plural := d.Arg("plural")
			if plural == "" {
				plural = o.Name + "s"
			}

			rendered, err := gen.ExecuteTemplate("plugins/tf_appsync_dynamo/templates/scan.tmpl", scanData{
				Object: o,
				Plural: plural,
			})
			if err != nil {
				return fmt.Errorf("failed rendering Scan %s: %w", o.Name, err)
			}

			_, err = output.AppendOrCreate(gen.TF_API_GEN, util.DashCase(o.Name), rendered)
			if err != nil {
				return fmt.Errorf("failed appending Scan %s: %w", o.Name, err)
			}
		}

		if d := o.SingleDirective("appsync_crud"); d != nil {
			// CRUD blocks
			cd := crudData{
				Object:         o,
				HashKey:        hashKey,
				SortKey:        sortKey,
				HasSort:        dynamo.HasArg("sort_key"),
				CreateHashType: d.Arg("create_hash_type"),
				CreateSortType: d.Arg("create_sort_type"),
			}

			if !d.ArgBool("disable_create") {
				rendered, err := gen.ExecuteTemplate("plugins/tf_appsync_dynamo/templates/crud_create.tmpl", cd)
				if err != nil {
					return fmt.Errorf("failed rendering CRUD create %s: %w", o.Name, err)
				}

				_, err = output.AppendOrCreate(gen.TF_API_GEN, util.DashCase(o.Name), rendered)
				if err != nil {
					return fmt.Errorf("failed appending CRUD create %s: %w", o.Name, err)
				}
			}

			if !d.ArgBool("disable_read") {
				rendered, err := gen.ExecuteTemplate("plugins/tf_appsync_dynamo/templates/crud_read.tmpl", cd)
				if err != nil {
					return fmt.Errorf("failed rendering CRUD read %s: %w", o.Name, err)
				}

				_, err = output.AppendOrCreate(gen.TF_API_GEN, util.DashCase(o.Name), rendered)
				if err != nil {
					return fmt.Errorf("failed appending CRUD read %s: %w", o.Name, err)
				}
			}

			if !d.ArgBool("disable_update") {
				rendered, err := gen.ExecuteTemplate("plugins/tf_appsync_dynamo/templates/crud_update.tmpl", cd)
				if err != nil {
					return fmt.Errorf("failed rendering CRUD update %s: %w", o.Name, err)
				}

				_, err = output.AppendOrCreate(gen.TF_API_GEN, util.DashCase(o.Name), rendered)
				if err != nil {
					return fmt.Errorf("failed appending CRUD update %s: %w", o.Name, err)
				}
			}

			if !d.ArgBool("disable_delete") {
				rendered, err := gen.ExecuteTemplate("plugins/tf_appsync_dynamo/templates/crud_delete.tmpl", cd)
				if err != nil {
					return fmt.Errorf("failed rendering CRUD delete %s: %w", o.Name, err)
				}

				_, err = output.AppendOrCreate(gen.TF_API_GEN, util.DashCase(o.Name), rendered)
				if err != nil {
					return fmt.Errorf("failed appending CRUD delete %s: %w", o.Name, err)
				}
			}
		}

		for _, f := range o.Fields {
			if f.HasDirective("appsync_list") {
				for _, d := range f.Directives["appsync_list"] {
					plural := d.Arg("plural")
					if plural == "" {
						plural = o.Name + "s"
					}
					queryName := d.Arg("name")
					if queryName == "" {
						if d.Arg("plural") == "" {
							return fmt.Errorf("appsync_list on %s %s has no name or using", o.Name, f.Name)
						}
						if d.HasArg("using") {
							queryName = "list" + plural + "By" + util.TitleCase(d.Arg("using"))
						} else {
							queryName = "list" + plural + "By" + hashKey.NameTitle()
						}
					}

					rendered, err := gen.ExecuteTemplate("plugins/tf_appsync_dynamo/templates/list.tmpl", listData{
						Object:    o,
						QueryName: queryName,
						Field:     f,
						Directive: d,
					})
					if err != nil {
						return fmt.Errorf("failed rendering List %s %s: %w", o.Name, f.Name, err)
					}

					_, err = output.AppendOrCreate(gen.TF_API_GEN, util.DashCase(o.Name), rendered)
					if err != nil {
						return fmt.Errorf("failed appending List %s %s: %w", o.Name, f.Name, err)
					}
				}
			}

			if f.HasDirective("appsync_foreign_key") {
				for _, d := range f.Directives["appsync_foreign_key"] {

					tmpl := "foreign_key_"
					if d.ArgBool("query_single") {
						tmpl += "query_single"
					} else if d.ArgBool("query") {
						tmpl += "query"
					} else if d.ArgBool("batch") {
						tmpl += "batch"
					} else {
						tmpl += "single"
					}

					rendered, err := gen.ExecuteTemplate("plugins/tf_appsync_dynamo/templates/"+tmpl+".tmpl", fkeyData{
						Object:                 o,
						Field:                  f,
						Directive:              d,
						Query:                  d.ArgBool("query"),
						QuerySingle:            d.ArgBool("query_single"),
						Batch:                  d.ArgBool("batch"),
						FieldSource:            d.Arg("field_source"),
						FieldForeign:           d.Arg("field_foreign"),
						Table:                  d.Arg("table"),
						Index:                  d.Arg("index"),
						AdditionalFieldSource:  d.Arg("additional_field_source"),
						AdditionalFieldForeign: d.Arg("additional_field_foreign"),
					})
					if err != nil {
						return fmt.Errorf("failed rendering Foreign Key %s %s: %w", o.Name, f.Name, err)
					}

					_, err = output.AppendOrCreate(gen.TF_API_GEN, util.DashCase(o.Name), rendered)
					if err != nil {
						return fmt.Errorf("failed appending Foreign Key %s %s: %w", o.Name, f.Name, err)
					}
				}
			}
		}
	}

	return nil
}
