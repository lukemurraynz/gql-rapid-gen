// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package go_common

import (
	"github.com/mjdrgn/gql-rapid-gen/gen"
	"github.com/mjdrgn/gql-rapid-gen/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPlugin_Execute(t *testing.T) {
	plugin := &Plugin{}

	schema := &parser.Schema{
		Objects: map[string]*parser.ParsedObject{
			"TestObjectName": {
				Name: "TestObjectName",
				Directives: map[string][]*parser.ParsedDirective{
					"dynamodb": {
						&parser.ParsedDirective{
							Name:      "dynamodb",
							Arguments: nil,
						},
					},
				},
				Description: "Description",
				Fields:      []*parser.ParsedField{},
				Interfaces:  nil,
			},
			"Test2": {
				Name: "Test2",
				Directives: map[string][]*parser.ParsedDirective{
					"dynamodb": {
						&parser.ParsedDirective{
							Name:      "dynamodb",
							Arguments: nil,
						},
					},
				},
				Description: "Description",
				Fields:      []*parser.ParsedField{},
				Interfaces:  nil,
			},
			"Test3": {
				Name:        "Test3",
				Directives:  nil,
				Description: "Description",
				Fields:      []*parser.ParsedField{},
				Interfaces:  nil,
			},
		},
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

	assert.True(t, len(rendered) > 100)
	assert.Contains(t, rendered, "initTestObjectNameProvider")
	assert.Contains(t, rendered, "initTestObjectNameMock")
	assert.Contains(t, rendered, "initTest2Provider")
	assert.NotContains(t, rendered, "initTest3Provider")
}
