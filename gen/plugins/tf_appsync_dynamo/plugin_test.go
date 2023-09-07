// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package tf_appsync_dynamo

import (
	"github.com/mjdrgn/gql-rapid-gen/gen"
	"github.com/mjdrgn/gql-rapid-gen/gen/testdata"
	"github.com/mjdrgn/gql-rapid-gen/parser"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPlugin_Execute_List_Index(t *testing.T) {
	obj := &parser.ParsedObject{
		Name: "TestObjectName",
		Directives: map[string][]*parser.ParsedDirective{
			"dynamodb": {
				&parser.ParsedDirective{
					Name: "dynamodb",
					Arguments: parser.ArgumentsFromMap(map[string]any{
						"name":     "TestObjectName",
						"hash_key": "myhash",
					}),
				},
			},
			"dynamodb_gsi": {
				&parser.ParsedDirective{
					Name: "dynamodb_gsi",
					Arguments: parser.ArgumentsFromMap(map[string]any{
						"name":     "test1_test2",
						"hash_key": "test1",
						"sort_key": "test2",
					}),
				},
			},
		},
		Description: "Description",
		Fields: []*parser.ParsedField{
			{
				Name:        "myhash",
				Directives:  nil,
				Description: "Hash Key",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
			{
				Name: "test1",
				Directives: map[string][]*parser.ParsedDirective{
					"appsync_list": {
						&parser.ParsedDirective{
							Name: "appsync_list",
							Arguments: parser.ArgumentsFromMap(map[string]any{
								"plural": "Tests",
								"using":  "test1_test2",
							}),
						},
					},
				},
				Description: "Test 1",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
			{
				Name:        "test2",
				Directives:  nil,
				Description: "Test 2",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
		},
		Interfaces: nil,
	}

	plugin := &Plugin{}

	schema := &parser.Schema{
		Objects: map[string]*parser.ParsedObject{
			"TestObjectName": obj,
		},
	}

	output := &gen.Output{}

	err := plugin.Generate(schema, output)
	require.Nil(t, err)

	files := output.GetFiles()

	require.Equal(t, 2, len(files))

	// TODO detailed checks
}

func TestPlugin_Execute_List_Index_Reverse(t *testing.T) {
	obj := &parser.ParsedObject{
		Name: "TestObjectName",
		Directives: map[string][]*parser.ParsedDirective{
			"dynamodb": {
				&parser.ParsedDirective{
					Name: "dynamodb",
					Arguments: parser.ArgumentsFromMap(map[string]any{
						"name":     "TestObjectName",
						"hash_key": "myhash",
					}),
				},
			},
			"dynamodb_gsi": {
				&parser.ParsedDirective{
					Name: "dynamodb_gsi",
					Arguments: parser.ArgumentsFromMap(map[string]any{
						"name":     "test1_test2",
						"hash_key": "test1",
						"sort_key": "test2",
					}),
				},
			},
		},
		Description: "Description",
		Fields: []*parser.ParsedField{
			{
				Name:        "myhash",
				Directives:  nil,
				Description: "Hash Key",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
			{
				Name: "test1",
				Directives: map[string][]*parser.ParsedDirective{
					"appsync_list": {
						&parser.ParsedDirective{
							Name: "appsync_list",
							Arguments: parser.ArgumentsFromMap(map[string]any{
								"plural":  "Tests",
								"using":   "test1_test2",
								"forward": false,
							}),
						},
					},
				},
				Description: "Test 1",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
			{
				Name:        "test2",
				Directives:  nil,
				Description: "Test 2",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
		},
		Interfaces: nil,
	}

	plugin := &Plugin{}

	schema := &parser.Schema{
		Objects: map[string]*parser.ParsedObject{
			"TestObjectName": obj,
		},
	}

	output := &gen.Output{}

	err := plugin.Generate(schema, output)
	require.Nil(t, err)

	files := output.GetFiles()

	require.Equal(t, 2, len(files))

	// TODO detailed checks
}

func TestPlugin_Execute_List_Hash(t *testing.T) {
	obj := &parser.ParsedObject{
		Name: "TestObjectName",
		Directives: map[string][]*parser.ParsedDirective{
			"dynamodb": {
				&parser.ParsedDirective{
					Name: "dynamodb",
					Arguments: parser.ArgumentsFromMap(map[string]any{
						"name":     "TestObjectName",
						"hash_key": "myhash",
						"sort_key": "mysort",
					}),
				},
			},
		},
		Description: "Description",
		Fields: []*parser.ParsedField{
			{
				Name: "myhash",
				Directives: map[string][]*parser.ParsedDirective{
					"appsync_list": {
						&parser.ParsedDirective{
							Name: "appsync_list",
							Arguments: parser.ArgumentsFromMap(map[string]any{
								"plural": "Tests",
							}),
						},
					},
				},
				Description: "Hash Key",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
			{
				Name:        "mysort",
				Directives:  nil,
				Description: "Sort Key",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
		},
		Interfaces: nil,
	}

	plugin := &Plugin{}

	schema := &parser.Schema{
		Objects: map[string]*parser.ParsedObject{
			"TestObjectName": obj,
		},
	}

	output := &gen.Output{}

	err := plugin.Generate(schema, output)
	require.Nil(t, err)

	files := output.GetFiles()

	require.Equal(t, 2, len(files))

	// TODO detailed checks
}

func TestPlugin_Execute_List_Hash_Name(t *testing.T) {
	obj := &parser.ParsedObject{
		Name: "TestObjectName",
		Directives: map[string][]*parser.ParsedDirective{
			"dynamodb": {
				&parser.ParsedDirective{
					Name: "dynamodb",
					Arguments: parser.ArgumentsFromMap(map[string]any{
						"name":     "TestObjectName",
						"hash_key": "myhash",
						"sort_key": "mysort",
					}),
				},
			},
		},
		Description: "Description",
		Fields: []*parser.ParsedField{
			{
				Name: "myhash",
				Directives: map[string][]*parser.ParsedDirective{
					"appsync_list": {
						&parser.ParsedDirective{
							Name: "appsync_list",
							Arguments: parser.ArgumentsFromMap(map[string]any{
								"name": "listMyTests",
							}),
						},
					},
				},
				Description: "Hash Key",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
			{
				Name:        "mysort",
				Directives:  nil,
				Description: "Sort Key",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
		},
		Interfaces: nil,
	}

	plugin := &Plugin{}

	schema := &parser.Schema{
		Objects: map[string]*parser.ParsedObject{
			"TestObjectName": obj,
		},
	}

	output := &gen.Output{}

	err := plugin.Generate(schema, output)
	require.Nil(t, err)

	files := output.GetFiles()

	require.Equal(t, 2, len(files))

	// TODO detailed checks
}

func TestPlugin_Execute_Fkey_Single(t *testing.T) {
	obj := &parser.ParsedObject{
		Name: "TestObjectName",
		Directives: map[string][]*parser.ParsedDirective{
			"dynamodb": {
				&parser.ParsedDirective{
					Name: "dynamodb",
					Arguments: parser.ArgumentsFromMap(map[string]any{
						"name":     "TestObjectName",
						"hash_key": "myhash",
					}),
				},
			},
		},
		Description: "Description",
		Fields: []*parser.ParsedField{
			{
				Name:        "myhash",
				Directives:  nil,
				Description: "Hash Key",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
			{
				Name:        "ref",
				Directives:  nil,
				Description: "Ref Val",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
			{
				Name: "obj",
				Directives: map[string][]*parser.ParsedDirective{
					"appsync_foreign_key": {
						&parser.ParsedDirective{
							Name: "appsync_foreign_key",
							Arguments: parser.ArgumentsFromMap(map[string]any{
								"field_source":  "ref",
								"field_foreign": "id",
								"table":         "other",
							}),
						},
					},
				},
				Description: "Obj",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
		},
		Interfaces: nil,
	}

	plugin := &Plugin{}

	schema := &parser.Schema{
		Objects: map[string]*parser.ParsedObject{
			"TestObjectName": obj,
		},
	}

	output := &gen.Output{}

	err := plugin.Generate(schema, output)
	require.Nil(t, err)

	files := output.GetFiles()

	require.Equal(t, 2, len(files))

	// TODO detailed checks
}

func TestPlugin_Execute_Fkey_Query_Single(t *testing.T) {
	obj := &parser.ParsedObject{
		Name: "TestObjectName",
		Directives: map[string][]*parser.ParsedDirective{
			"dynamodb": {
				&parser.ParsedDirective{
					Name: "dynamodb",
					Arguments: parser.ArgumentsFromMap(map[string]any{
						"name":     "TestObjectName",
						"hash_key": "myhash",
					}),
				},
			},
		},
		Description: "Description",
		Fields: []*parser.ParsedField{
			{
				Name:        "myhash",
				Directives:  nil,
				Description: "Hash Key",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
			{
				Name:        "ref",
				Directives:  nil,
				Description: "Ref Val",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
			{
				Name: "obj",
				Directives: map[string][]*parser.ParsedDirective{
					"appsync_foreign_key": {
						&parser.ParsedDirective{
							Name: "appsync_foreign_key",
							Arguments: parser.ArgumentsFromMap(map[string]any{
								"field_source":  "ref",
								"field_foreign": "id",
								"table":         "other",
								"index":         "test",
								"query_single":  true,
							}),
						},
					},
				},
				Description: "Obj",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
		},
		Interfaces: nil,
	}

	plugin := &Plugin{}

	schema := &parser.Schema{
		Objects: map[string]*parser.ParsedObject{
			"TestObjectName": obj,
		},
	}

	output := &gen.Output{}

	err := plugin.Generate(schema, output)
	require.Nil(t, err)

	files := output.GetFiles()

	require.Equal(t, 2, len(files))

	// TODO detailed checks
}

func TestPlugin_Execute_Fkey_Single_Index(t *testing.T) {
	obj := &parser.ParsedObject{
		Name: "TestObjectName",
		Directives: map[string][]*parser.ParsedDirective{
			"dynamodb": {
				&parser.ParsedDirective{
					Name: "dynamodb",
					Arguments: parser.ArgumentsFromMap(map[string]any{
						"name":     "TestObjectName",
						"hash_key": "myhash",
					}),
				},
			},
		},
		Description: "Description",
		Fields: []*parser.ParsedField{
			{
				Name:        "myhash",
				Directives:  nil,
				Description: "Hash Key",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
			{
				Name:        "ref",
				Directives:  nil,
				Description: "Ref Val",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
			{
				Name: "obj",
				Directives: map[string][]*parser.ParsedDirective{
					"appsync_foreign_key": {
						&parser.ParsedDirective{
							Name: "appsync_foreign_key",
							Arguments: parser.ArgumentsFromMap(map[string]any{
								"field_source":  "ref",
								"field_foreign": "id",
								"table":         "other",
								"index":         "target_gsi",
							}),
						},
					},
				},
				Description: "Obj",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
		},
		Interfaces: nil,
	}

	plugin := &Plugin{}

	schema := &parser.Schema{
		Objects: map[string]*parser.ParsedObject{
			"TestObjectName": obj,
		},
	}

	output := &gen.Output{}

	err := plugin.Generate(schema, output)
	require.Nil(t, err)

	files := output.GetFiles()

	require.Equal(t, 2, len(files))

	// TODO detailed checks
}

func TestPlugin_Execute_Fkey_Single_Additional(t *testing.T) {
	obj := &parser.ParsedObject{
		Name: "TestObjectName",
		Directives: map[string][]*parser.ParsedDirective{
			"dynamodb": {
				&parser.ParsedDirective{
					Name: "dynamodb",
					Arguments: parser.ArgumentsFromMap(map[string]any{
						"name":     "TestObjectName",
						"hash_key": "myhash",
						"sort_key": "mysort",
					}),
				},
			},
		},
		Description: "Description",
		Fields: []*parser.ParsedField{
			{
				Name:        "myhash",
				Directives:  nil,
				Description: "Hash Key",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
			{
				Name:        "mysort",
				Directives:  nil,
				Description: "Sort Key",
				Arguments:   nil,
				Type:        parser.TypeIntReq,
			},
			{
				Name:        "ref",
				Directives:  nil,
				Description: "Ref Val",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
			{
				Name:        "ref2",
				Directives:  nil,
				Description: "Ref Val 2",
				Arguments:   nil,
				Type:        parser.TypeIntReq,
			},
			{
				Name: "obj",
				Directives: map[string][]*parser.ParsedDirective{
					"appsync_foreign_key": {
						&parser.ParsedDirective{
							Name: "appsync_foreign_key",
							Arguments: parser.ArgumentsFromMap(map[string]any{
								"field_source":             "ref",
								"field_foreign":            "id",
								"additional_field_source":  "ref2",
								"additional_field_foreign": "id2",
								"table":                    "other",
							}),
						},
					},
				},
				Description: "Obj",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
		},
		Interfaces: nil,
	}

	plugin := &Plugin{}

	schema := &parser.Schema{
		Objects: map[string]*parser.ParsedObject{
			"TestObjectName": obj,
		},
	}

	output := &gen.Output{}

	err := plugin.Generate(schema, output)
	require.Nil(t, err)

	files := output.GetFiles()

	require.Equal(t, 2, len(files))

	// TODO detailed checks
}

func TestPlugin_Execute_Fkey_Query(t *testing.T) {
	obj := &parser.ParsedObject{
		Name: "TestObjectName",
		Directives: map[string][]*parser.ParsedDirective{
			"dynamodb": {
				&parser.ParsedDirective{
					Name: "dynamodb",
					Arguments: parser.ArgumentsFromMap(map[string]any{
						"name":     "TestObjectName",
						"hash_key": "myhash",
					}),
				},
			},
		},
		Description: "Description",
		Fields: []*parser.ParsedField{
			{
				Name:        "myhash",
				Directives:  nil,
				Description: "Hash Key",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
			{
				Name:        "ref",
				Directives:  nil,
				Description: "Ref Val",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
			{
				Name: "obj",
				Directives: map[string][]*parser.ParsedDirective{
					"appsync_foreign_key": {
						&parser.ParsedDirective{
							Name: "appsync_foreign_key",
							Arguments: parser.ArgumentsFromMap(map[string]any{
								"field_source":  "ref",
								"field_foreign": "id",
								"table":         "other",
								"query":         true,
							}),
						},
					},
				},
				Description: "Obj",
				Arguments:   nil,
				Type:        parser.TypeStringReqCollection,
			},
		},
		Interfaces: nil,
	}

	plugin := &Plugin{}

	schema := &parser.Schema{
		Objects: map[string]*parser.ParsedObject{
			"TestObjectName": obj,
		},
	}

	output := &gen.Output{}

	err := plugin.Generate(schema, output)
	require.Nil(t, err)

	files := output.GetFiles()

	require.Equal(t, 2, len(files))

	// TODO detailed checks
}

func TestPlugin_Execute_Fkey_Batch(t *testing.T) {
	obj := &parser.ParsedObject{
		Name: "TestObjectName",
		Directives: map[string][]*parser.ParsedDirective{
			"dynamodb": {
				&parser.ParsedDirective{
					Name: "dynamodb",
					Arguments: parser.ArgumentsFromMap(map[string]any{
						"name":     "TestObjectName",
						"hash_key": "myhash",
					}),
				},
			},
		},
		Description: "Description",
		Fields: []*parser.ParsedField{
			{
				Name:        "myhash",
				Directives:  nil,
				Description: "Hash Key",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
			{
				Name:        "ref",
				Directives:  nil,
				Description: "Ref Val",
				Arguments:   nil,
				Type:        parser.TypeStringReqCollection,
			},
			{
				Name: "obj",
				Directives: map[string][]*parser.ParsedDirective{
					"appsync_foreign_key": {
						&parser.ParsedDirective{
							Name: "appsync_foreign_key",
							Arguments: parser.ArgumentsFromMap(map[string]any{
								"field_source":  "ref",
								"field_foreign": "id",
								"table":         "other",
								"batch":         true,
							}),
						},
					},
				},
				Description: "Obj",
				Arguments:   nil,
				Type:        parser.TypeStringReqCollection,
			},
		},
		Interfaces: nil,
	}

	plugin := &Plugin{}

	schema := &parser.Schema{
		Objects: map[string]*parser.ParsedObject{
			"TestObjectName": obj,
		},
	}

	output := &gen.Output{}

	err := plugin.Generate(schema, output)
	require.Nil(t, err)

	files := output.GetFiles()

	require.Equal(t, 2, len(files))

	// TODO detailed checks
}

func TestPlugin_Execute_Fkey_Batch_Additional(t *testing.T) {
	obj := &parser.ParsedObject{
		Name: "TestObjectName",
		Directives: map[string][]*parser.ParsedDirective{
			"dynamodb": {
				&parser.ParsedDirective{
					Name: "dynamodb",
					Arguments: parser.ArgumentsFromMap(map[string]any{
						"name":     "TestObjectName",
						"hash_key": "myhash",
						"sort_key": "mysort",
					}),
				},
			},
		},
		Description: "Description",
		Fields: []*parser.ParsedField{
			{
				Name:        "myhash",
				Directives:  nil,
				Description: "Hash Key",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
			{
				Name:        "mysort",
				Directives:  nil,
				Description: "Sort Key",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
			{
				Name:        "ref",
				Directives:  nil,
				Description: "Ref Val",
				Arguments:   nil,
				Type:        parser.TypeStringReqCollection,
			},
			{
				Name:        "ref2",
				Directives:  nil,
				Description: "Ref Val 2",
				Arguments:   nil,
				Type:        parser.TypeStringReqCollection,
			},
			{
				Name: "obj",
				Directives: map[string][]*parser.ParsedDirective{
					"appsync_foreign_key": {
						&parser.ParsedDirective{
							Name: "appsync_foreign_key",
							Arguments: parser.ArgumentsFromMap(map[string]any{
								"field_source":             "ref",
								"field_foreign":            "id",
								"additional_field_source":  "ref2",
								"additional_field_foreign": "id2",
								"table":                    "other",
								"batch":                    true,
							}),
						},
					},
				},
				Description: "Obj",
				Arguments:   nil,
				Type:        parser.TypeStringReqCollection,
			},
		},
		Interfaces: nil,
	}

	plugin := &Plugin{}

	schema := &parser.Schema{
		Objects: map[string]*parser.ParsedObject{
			"TestObjectName": obj,
		},
	}

	output := &gen.Output{}

	err := plugin.Generate(schema, output)
	require.Nil(t, err)

	files := output.GetFiles()

	require.Equal(t, 2, len(files))

	// TODO detailed checks
}

func TestPlugin_Execute_Simple(t *testing.T) {
	plugin := &Plugin{}

	schema := &parser.Schema{
		Objects: map[string]*parser.ParsedObject{
			"TestObjectName": testdata.DynamoDBSimple("TestObjectName"),
		},
	}
	schema.Objects["TestObjectName"].Directives["appsync_crud"] = []*parser.ParsedDirective{
		{
			Name:      "appsync_crud",
			Arguments: nil,
		},
	}
	schema.Objects["TestObjectName"].Directives["appsync_scan"] = []*parser.ParsedDirective{
		{
			Name:      "appsync_scan",
			Arguments: nil,
		},
	}

	output := &gen.Output{}

	err := plugin.Generate(schema, output)
	require.Nil(t, err)

	files := output.GetFiles()

	require.Equal(t, 3, len(files))

	// TODO detailed checks
}

func TestPlugin_Execute_Composite(t *testing.T) {
	plugin := &Plugin{}

	schema := &parser.Schema{
		Objects: map[string]*parser.ParsedObject{
			"TestObjectName": testdata.DynamoDBComposite("TestObjectName"),
		},
	}
	schema.Objects["TestObjectName"].Directives["appsync_crud"] = []*parser.ParsedDirective{
		{
			Name:      "appsync_crud",
			Arguments: nil,
		},
	}
	schema.Objects["TestObjectName"].Directives["appsync_scan"] = []*parser.ParsedDirective{
		{
			Name:      "appsync_scan",
			Arguments: nil,
		},
	}

	output := &gen.Output{}

	err := plugin.Generate(schema, output)
	require.Nil(t, err)

	files := output.GetFiles()

	require.Equal(t, 3, len(files))

	// TODO detailed checks
}

func TestPlugin_Execute_CompositeNumbers(t *testing.T) {
	plugin := &Plugin{}

	schema := &parser.Schema{
		Objects: map[string]*parser.ParsedObject{
			"TestObjectName": testdata.DynamoDBCompositeNumbers("TestObjectName"),
		},
	}
	schema.Objects["TestObjectName"].Directives["appsync_crud"] = []*parser.ParsedDirective{
		{
			Name:      "appsync_crud",
			Arguments: nil,
		},
	}
	schema.Objects["TestObjectName"].Directives["appsync_scan"] = []*parser.ParsedDirective{
		{
			Name:      "appsync_scan",
			Arguments: nil,
		},
	}

	output := &gen.Output{}

	err := plugin.Generate(schema, output)
	require.Nil(t, err)

	files := output.GetFiles()

	require.Equal(t, 3, len(files))

	// TODO detailed checks
}

func TestPlugin_Execute_GSI(t *testing.T) {
	plugin := &Plugin{}

	schema := &parser.Schema{
		Objects: map[string]*parser.ParsedObject{
			"TestObjectName": testdata.DynamoDBGSI("TestObjectName"),
		},
	}
	schema.Objects["TestObjectName"].Directives["appsync_crud"] = []*parser.ParsedDirective{
		{
			Name:      "appsync_crud",
			Arguments: nil,
		},
	}
	schema.Objects["TestObjectName"].Directives["appsync_scan"] = []*parser.ParsedDirective{
		{
			Name:      "appsync_scan",
			Arguments: nil,
		},
	}

	output := &gen.Output{}

	err := plugin.Generate(schema, output)
	require.Nil(t, err)

	files := output.GetFiles()

	require.Equal(t, 3, len(files))

	// TODO detailed checks
}

func TestPlugin_Execute_GSIOrdered(t *testing.T) {
	plugin := &Plugin{}

	schema := &parser.Schema{
		Objects: map[string]*parser.ParsedObject{
			"TestObjectName": testdata.DynamoDBGSIOrdered("TestObjectName"),
		},
	}
	schema.Objects["TestObjectName"].Directives["appsync_crud"] = []*parser.ParsedDirective{
		{
			Name:      "appsync_crud",
			Arguments: nil,
		},
	}
	schema.Objects["TestObjectName"].Directives["appsync_scan"] = []*parser.ParsedDirective{
		{
			Name:      "appsync_scan",
			Arguments: nil,
		},
	}

	output := &gen.Output{}

	err := plugin.Generate(schema, output)
	require.Nil(t, err)

	files := output.GetFiles()

	require.Equal(t, 3, len(files))

	// TODO detailed checks
}

func TestPlugin_Execute_GSIComposite(t *testing.T) {
	plugin := &Plugin{}

	schema := &parser.Schema{
		Objects: map[string]*parser.ParsedObject{
			"TestObjectName": testdata.DynamoDBGSIComposite("TestObjectName"),
		},
	}
	schema.Objects["TestObjectName"].Directives["appsync_crud"] = []*parser.ParsedDirective{
		{
			Name:      "appsync_crud",
			Arguments: nil,
		},
	}
	schema.Objects["TestObjectName"].Directives["appsync_scan"] = []*parser.ParsedDirective{
		{
			Name:      "appsync_scan",
			Arguments: nil,
		},
	}

	output := &gen.Output{}

	err := plugin.Generate(schema, output)
	require.Nil(t, err)

	files := output.GetFiles()

	require.Equal(t, 3, len(files))

	// TODO detailed checks
}
