// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package aws_skel

import (
	"github.com/mjdrgn/gql-rapid-gen/gen"
	"github.com/mjdrgn/gql-rapid-gen/parser"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPlugin_Execute(t *testing.T) {
	plugin := &Plugin{}

	schema := &parser.Schema{}

	output := &gen.Output{}

	err := plugin.Generate(schema, output)
	require.Nil(t, err)

	files := output.GetFiles()

	require.Greater(t, len(files), 1)
}
