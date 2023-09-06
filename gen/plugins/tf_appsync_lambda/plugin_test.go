// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package tf_appsync_lambda

import (
	"github.com/mjdrgn/gql-rapid-gen/gen"
	"github.com/mjdrgn/gql-rapid-gen/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestPlugin_Execute_Go(t *testing.T) {
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
				Arguments:   nil,
				Type:        parser.TypeStringReqCollection,
			},
		},
	}

	output := &gen.Output{}

	err := plugin.Generate(schema, output)
	require.Nil(t, err)

	files := output.GetFiles()

	require.Equal(t, 2, len(files))

	var def *gen.OutputFile
	var skel *gen.OutputFile
	for name, f := range files {
		if strings.Contains(name, ".gen.") {
			def = f
		} else {
			skel = f
		}
	}
	require.NotNil(t, def)
	require.NotNil(t, skel)

	accessRendered := def.String()
	skelRendered := skel.String()

	assert.True(t, len(accessRendered) > 100)
	assert.True(t, len(skelRendered) > 100)

	// TODO more detailed validation
}
