// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package parser

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Parse_Object(t *testing.T) {
	files := []string{"testdata/obj.graphql"}

	output, err := Parse(files)

	require.Nilf(t, err, "Parse returned error: %s", err)

	require.NotNil(t, output.Objects)
	require.NotNil(t, output.Objects["Test"])
}

func Test_Parse_Types(t *testing.T) {
	files := []string{"testdata/obj.graphql"}

	output, err := Parse(files)

	require.Nilf(t, err, "Parse returned error: %s", err)

	require.NotNil(t, output.Objects)
	require.NotNil(t, output.Objects["Test"])

	obj := output.Objects["Test"]
	require.NotNil(t, obj)

	one := obj.Field("One")
	two := obj.Field("Two")
	three := obj.Field("Three")
	four := obj.Field("Four")
	oneReq := obj.Field("OneReq")
	twoReq := obj.Field("TwoReq")
	threeReq := obj.Field("ThreeReq")
	fourReq := obj.Field("FourReq")
	loop := obj.Field("Loop")

	assert.NotNil(t, one)
	assert.NotNil(t, two)
	assert.NotNil(t, three)
	assert.NotNil(t, four)
	assert.NotNil(t, oneReq)
	assert.NotNil(t, twoReq)
	assert.NotNil(t, threeReq)
	assert.NotNil(t, fourReq)
	assert.NotNil(t, loop)

	assert.Equal(t, "String", one.Type.Kind)
	assert.Equal(t, "Int", two.Type.Kind)
	assert.Equal(t, "Float", three.Type.Kind)
	assert.Equal(t, "Boolean", four.Type.Kind)
	assert.False(t, one.Type.Required)
	assert.False(t, two.Type.Required)
	assert.False(t, three.Type.Required)
	assert.False(t, four.Type.Required)
}

func Test_Parse_RepeatedDirectives(t *testing.T) {
	files := []string{"testdata/dir_repeat.graphql"}

	output, err := Parse(files)

	require.Nilf(t, err, "Parse returned error: %s", err)

	require.NotNil(t, output.Objects)
	require.NotNil(t, output.Objects["T2"])
	obj := output.Objects["T2"]
	require.True(t, obj.HasDirective("dynamodb"))
	require.True(t, obj.HasDirective("dynamodb_gsi"))
}
