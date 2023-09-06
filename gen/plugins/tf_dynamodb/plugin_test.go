// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package tf_dynamodb

import (
	"github.com/mjdrgn/gql-rapid-gen/gen"
	"github.com/mjdrgn/gql-rapid-gen/gen/testdata"
	"github.com/mjdrgn/gql-rapid-gen/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPlugin_Execute_Simple(t *testing.T) {
	plugin := &Plugin{}

	schema := &parser.Schema{
		Objects: map[string]*parser.ParsedObject{
			"TestObjectName": testdata.DynamoDBSimple("TestObjectName"),
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
	}
	require.NotNil(t, file)

	rendered := file.String()

	assert.True(t, len(rendered) > 100)

	// TODO detailed checks
}

func TestPlugin_Execute_Composite(t *testing.T) {
	plugin := &Plugin{}

	schema := &parser.Schema{
		Objects: map[string]*parser.ParsedObject{
			"TestObjectName": testdata.DynamoDBComposite("TestObjectName"),
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
	}
	require.NotNil(t, file)

	rendered := file.String()

	assert.True(t, len(rendered) > 100)

	// TODO detailed checks
}

func TestPlugin_Execute_CompositeNumbers(t *testing.T) {
	plugin := &Plugin{}

	schema := &parser.Schema{
		Objects: map[string]*parser.ParsedObject{
			"TestObjectName": testdata.DynamoDBCompositeNumbers("TestObjectName"),
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
	}
	require.NotNil(t, file)

	rendered := file.String()

	assert.True(t, len(rendered) > 100)

	// TODO detailed checks
}

func TestPlugin_Execute_GSI(t *testing.T) {
	plugin := &Plugin{}

	schema := &parser.Schema{
		Objects: map[string]*parser.ParsedObject{
			"TestObjectName": testdata.DynamoDBGSI("TestObjectName"),
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
	}
	require.NotNil(t, file)

	rendered := file.String()

	assert.True(t, len(rendered) > 100)

	// TODO detailed checks
}

func TestPlugin_Execute_GSIOrdered(t *testing.T) {
	plugin := &Plugin{}

	schema := &parser.Schema{
		Objects: map[string]*parser.ParsedObject{
			"TestObjectName": testdata.DynamoDBGSIOrdered("TestObjectName"),
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
	}
	require.NotNil(t, file)

	rendered := file.String()

	assert.True(t, len(rendered) > 100)

	// TODO detailed checks
}

func TestPlugin_Execute_GSIComposite(t *testing.T) {
	plugin := &Plugin{}

	schema := &parser.Schema{
		Objects: map[string]*parser.ParsedObject{
			"TestObjectName": testdata.DynamoDBGSIComposite("TestObjectName"),
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
	}
	require.NotNil(t, file)

	rendered := file.String()

	assert.True(t, len(rendered) > 100)

	// TODO detailed checks
}
