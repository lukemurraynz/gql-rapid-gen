package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFieldType_IsObject(t *testing.T) {
	assert.False(t, (&FieldType{Collection: true}).IsObject())

	assert.False(t, (&FieldType{Kind: "String"}).IsObject())
	assert.False(t, (&FieldType{Kind: "Int"}).IsObject())
	assert.False(t, (&FieldType{Kind: "Float"}).IsObject())
	assert.False(t, (&FieldType{Kind: "Long"}).IsObject())
	assert.False(t, (&FieldType{Kind: "Boolean"}).IsObject())
	assert.False(t, (&FieldType{Kind: "ID"}).IsObject())

	assert.True(t, (&FieldType{Kind: "Test"}).IsObject())
	assert.True(t, (&FieldType{Kind: "ABC"}).IsObject())
}

func TestFieldType_DynamoType(t *testing.T) {
	assert.Equal(t, "S", (&FieldType{Kind: "String"}).DynamoType())
	assert.Equal(t, "N", (&FieldType{Kind: "Int"}).DynamoType())
	assert.Equal(t, "N", (&FieldType{Kind: "Float"}).DynamoType())
	assert.Equal(t, "N", (&FieldType{Kind: "Long"}).DynamoType())
	assert.Equal(t, "BOOL", (&FieldType{Kind: "Boolean"}).DynamoType())
	assert.Equal(t, "S", (&FieldType{Kind: "ID"}).DynamoType())
	assert.Equal(t, "M", (&FieldType{Kind: "MyObject"}).DynamoType())
	assert.Equal(t, "L", (&FieldType{Collection: true}).DynamoType())
}

func TestFieldType_GoType_Simple(t *testing.T) {
	assert.Equal(t, "*string", (&FieldType{Kind: "String"}).GoType())
	assert.Equal(t, "*int", (&FieldType{Kind: "Int"}).GoType())
	assert.Equal(t, "*float64", (&FieldType{Kind: "Float"}).GoType())
	assert.Equal(t, "*int64", (&FieldType{Kind: "Long"}).GoType())
	assert.Equal(t, "*bool", (&FieldType{Kind: "Boolean"}).GoType())
	assert.Equal(t, "*string", (&FieldType{Kind: "ID"}).GoType())
	assert.Equal(t, "*MyObject", (&FieldType{Kind: "MyObject"}).GoType())
}

func TestFieldType_GoType_Required(t *testing.T) {
	assert.Equal(t, "string", (&FieldType{Kind: "String", Required: true}).GoType())
	assert.Equal(t, "int", (&FieldType{Kind: "Int", Required: true}).GoType())
	assert.Equal(t, "float64", (&FieldType{Kind: "Float", Required: true}).GoType())
	assert.Equal(t, "int64", (&FieldType{Kind: "Long", Required: true}).GoType())
	assert.Equal(t, "bool", (&FieldType{Kind: "Boolean", Required: true}).GoType())
	assert.Equal(t, "string", (&FieldType{Kind: "ID", Required: true}).GoType())
	assert.Equal(t, "MyObject", (&FieldType{Kind: "MyObject", Required: true}).GoType())
}

func TestFieldType_GoType_Collection(t *testing.T) {
	assert.Equal(t, "[]*string", (&FieldType{CollectionSubtype: &FieldType{Kind: "String"}, Collection: true}).GoType())
	assert.Equal(t, "[]string", (&FieldType{CollectionSubtype: &FieldType{Kind: "String", Required: true}, Collection: true}).GoType())

	assert.Equal(t, "[]*string", (&FieldType{CollectionSubtype: &FieldType{Kind: "String"}, Collection: true, Required: true}).GoType())
	assert.Equal(t, "[]string", (&FieldType{CollectionSubtype: &FieldType{Kind: "String", Required: true}, Collection: true, Required: true}).GoType())
}
