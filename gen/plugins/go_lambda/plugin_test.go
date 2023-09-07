// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package go_lambda

import (
	"github.com/mjdrgn/gql-rapid-gen/gen"
	"github.com/mjdrgn/gql-rapid-gen/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPlugin_Execute_Mutation(t *testing.T) {
	plugin := &Plugin{}

	schema := &parser.Schema{
		Mutation: map[string]*parser.ParsedField{
			"updateTest": {
				Name: "updateTest",
				Directives: map[string][]*parser.ParsedDirective{
					"appsync_lambda": {
						{
							Name: "appsync_lambda",
							Arguments: parser.ArgumentsFromMap(map[string]any{
								"language": "go",
								"path":     "update-test",
								"timeout":  10,
								"memory":   256,
							}),
						},
					},
				},
				Description: "Update Test",
				Arguments: map[string]*parser.ParsedArgumentDef{
					"s1": {
						Name: "s1",
						Type: parser.TypeStringReq,
					},
					"i1": {
						Name: "i1",
						Type: parser.TypeIntReq,
					},
					"sl1": {
						Name: "sl1",
						Type: parser.TypeStringReqCollection,
					},
				},
				Type: parser.TypeStringReqCollection,
			},
		},
	}

	output := &gen.Output{}

	err := plugin.Generate(schema, output)
	require.Nil(t, err)

	files := output.GetFiles()

	require.Equal(t, 3, len(files))

	//var file *gen.OutputFile
	//for _, f := range files {
	//	file = f
	//	break
	//}
	//
	//rendered := file.String()
	//
	//assert.True(t, len(rendered) > 100)

	// TODO more detailed validation
}

func TestPlugin_Execute_Query(t *testing.T) {
	plugin := &Plugin{}

	schema := &parser.Schema{
		Query: map[string]*parser.ParsedField{
			"getTest": {
				Name: "getTest",
				Directives: map[string][]*parser.ParsedDirective{
					"appsync_lambda": {
						{
							Name: "appsync_lambda",
							Arguments: parser.ArgumentsFromMap(map[string]any{
								"language": "go",
								"path":     "update-test",
								"timeout":  10,
								"memory":   256,
							}),
						},
					},
				},
				Description: "Update Test",
				Arguments: map[string]*parser.ParsedArgumentDef{
					"s1": {
						Name: "s1",
						Type: parser.TypeStringReq,
					},
					"i1": {
						Name: "i1",
						Type: parser.TypeIntReq,
					},
					"sl1": {
						Name: "sl1",
						Type: parser.TypeStringReqCollection,
					},
				},
				Type: parser.TypeStringReqCollection,
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

	// TODO more detailed validation
}
