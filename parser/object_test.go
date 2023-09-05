package parser

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_ParsedObject_Field(t *testing.T) {
	obj := &ParsedObject{
		Name:        "test",
		Directives:  nil,
		Description: "test",
		Fields: []*ParsedField{
			{
				Name:        "f1",
				Directives:  nil,
				Description: "f1",
				Arguments:   nil,
				Type:        nil,
			},
		},
		Interfaces: nil,
	}

	assert.NotNil(t, obj.Field("f1"))
	assert.Nil(t, obj.Field("other"))
}

func Test_ParsedObject_HasDirective(t *testing.T) {
	obj := &ParsedObject{
		Name: "test",
		Directives: map[string][]*ParsedDirective{
			"t1": {
				{
					Name:      "t1",
					Arguments: nil,
				},
			},
		},
		Description: "test",
		Fields:      nil,
		Interfaces:  nil,
	}

	assert.True(t, obj.HasDirective("t1"))
	assert.False(t, obj.HasDirective("other"))
}

func Test_ParsedObject_SingleDirective(t *testing.T) {
	obj := &ParsedObject{
		Name: "test",
		Directives: map[string][]*ParsedDirective{
			"t1": {
				{
					Name:      "t1",
					Arguments: nil,
				},
			},
		},
		Description: "test",
		Fields:      nil,
		Interfaces:  nil,
	}

	assert.True(t, obj.HasDirective("t1"))
	dir := obj.SingleDirective("t1")
	require.NotNil(t, dir)
	assert.Equal(t, "t1", dir.Name)
}
