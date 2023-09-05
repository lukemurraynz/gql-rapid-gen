// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package go_objects

import (
	"github.com/mjdrgn/gql-rapid-gen/gen"
	"github.com/mjdrgn/gql-rapid-gen/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPlugin_Qualify(t *testing.T) {
	plugin := &Plugin{}

	valid := &parser.Schema{
		Objects:      map[string]*parser.ParsedObject{"test": nil},
		InputObjects: map[string]*parser.ParsedObject{"test": nil},
	}

	assert.True(t, plugin.Qualify(valid))

	invalid := &parser.Schema{
		Objects:      map[string]*parser.ParsedObject{},
		InputObjects: map[string]*parser.ParsedObject{},
	}

	assert.False(t, plugin.Qualify(invalid))
}

func TestPlugin_Qualify_Input(t *testing.T) {
	plugin := &Plugin{}

	valid := &parser.Schema{
		InputObjects: map[string]*parser.ParsedObject{"test": nil},
	}

	assert.True(t, plugin.Qualify(valid))
}

func TestPlugin_Qualify_Type(t *testing.T) {
	plugin := &Plugin{}

	valid := &parser.Schema{
		Objects: map[string]*parser.ParsedObject{"test": nil},
	}

	assert.True(t, plugin.Qualify(valid))
}

func TestPlugin_Execute(t *testing.T) {
	plugin := &Plugin{}

	schema := &parser.Schema{
		InputObjects: map[string]*parser.ParsedObject{"test": &parser.ParsedObject{
			Name:        "TestObjectName",
			Directives:  nil,
			Description: "Description",
			Fields: []*parser.ParsedField{
				{
					Name:        "FieldName",
					Directives:  nil,
					Description: "F1 Desc",
					Arguments:   nil,
					Type: &parser.FieldType{
						Kind:              "String",
						Required:          true,
						Collection:        false,
						CollectionSubtype: nil,
					},
				},
			},
			Interfaces: nil,
		}},
	}

	output := &gen.Output{}

	err := plugin.Generate(schema, output)
	require.Nil(t, err)

	files := output.GetFiles()

	require.Equal(t, 1, len(files))

	var file *gen.OutputFile
	for _, f := range files {
		file = f
		break
	}

	rendered := file.String()

	assert.True(t, len(rendered) > 300)
	assert.Contains(t, rendered, "TestObjectName")
	assert.Contains(t, rendered, "FieldName")
	assert.Contains(t, rendered, "string")
}

func TestPlugin_Execute_Format(t *testing.T) {
	plugin := &Plugin{}

	schema := &parser.Schema{
		InputObjects: map[string]*parser.ParsedObject{"test": &parser.ParsedObject{
			Name:        "TestObjectName",
			Directives:  nil,
			Description: "Description",
			Fields: []*parser.ParsedField{
				{
					Name:        "FieldName",
					Directives:  nil,
					Description: "F1 Desc",
					Arguments:   nil,
					Type: &parser.FieldType{
						Kind:              "String",
						Required:          true,
						Collection:        false,
						CollectionSubtype: nil,
					},
				},
			},
			Interfaces: nil,
		}},
	}

	output := &gen.Output{}

	err := plugin.Generate(schema, output)
	require.Nil(t, err)

	files := output.GetFiles()

	require.Equal(t, 1, len(files))

	var file *gen.OutputFile
	for _, f := range files {
		file = f
		break
	}

	rendered := file.String()

	_, err = gen.Types[gen.GO_DATA_GEN].Format(rendered)
	require.Nil(t, err)
}

func TestPlugin_Execute_Fields(t *testing.T) {
	plugin := &Plugin{}

	schema := &parser.Schema{
		InputObjects: map[string]*parser.ParsedObject{"test": &parser.ParsedObject{
			Name:        "TestObjectName",
			Directives:  nil,
			Description: "Description",
			Fields: []*parser.ParsedField{
				{
					Name:        "str_field",
					Directives:  nil,
					Description: "F1 Desc",
					Arguments:   nil,
					Type: &parser.FieldType{
						Kind:              "String",
						Required:          true,
						Collection:        false,
						CollectionSubtype: nil,
					},
				},
				{
					Name:        "int_field",
					Directives:  nil,
					Description: "F1 Desc",
					Arguments:   nil,
					Type: &parser.FieldType{
						Kind:              "Int",
						Required:          true,
						Collection:        false,
						CollectionSubtype: nil,
					},
				},
				{
					Name:        "opt_str_field",
					Directives:  nil,
					Description: "F1 Desc",
					Arguments:   nil,
					Type: &parser.FieldType{
						Kind:              "String",
						Required:          false,
						Collection:        false,
						CollectionSubtype: nil,
					},
				},
				{
					Name:        "str_coll",
					Directives:  nil,
					Description: "F1 Desc",
					Arguments:   nil,
					Type: &parser.FieldType{
						Kind:       "",
						Required:   true,
						Collection: true,
						CollectionSubtype: &parser.FieldType{
							Kind:              "String",
							Required:          true,
							Collection:        false,
							CollectionSubtype: nil,
						},
					},
				},
			},
			Interfaces: nil,
		}},
	}

	output := &gen.Output{}

	err := plugin.Generate(schema, output)
	require.Nil(t, err)

	files := output.GetFiles()

	require.Equal(t, 1, len(files))

	var file *gen.OutputFile
	for _, f := range files {
		file = f
		break
	}

	rendered := file.String()

	_, err = gen.Types[gen.GO_DATA_GEN].Format(rendered)
	require.Nil(t, err)

	assert.Contains(t, rendered, "StrField")
	assert.Contains(t, rendered, "IntField")
	assert.Contains(t, rendered, "OptStrField")
	assert.Contains(t, rendered, "StrColl")
	assert.Contains(t, rendered, "[]string")
	assert.Contains(t, rendered, "*string")
	assert.Contains(t, rendered, "int")
	assert.NotContains(t, rendered, "*int")
	assert.NotContains(t, rendered, "[]*string")
	assert.NotContains(t, rendered, "*[]")
}
