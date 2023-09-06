// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package parser

import (
	"fmt"
	"github.com/vektah/gqlparser/v2/ast"
	"log"
)

// FieldType is a struct that represents a field type, generated by parsing a graphql schema
type FieldType struct {
	Kind              string
	Required          bool
	Collection        bool
	CollectionSubtype *FieldType
}

// TypeStringReq is a pre-filled FieldType for a Required String field. This is used primarily for unit testing.
var TypeStringReq = &FieldType{
	Kind:              "String",
	Required:          true,
	Collection:        false,
	CollectionSubtype: nil,
}

// TypeStringReqCollection is a pre-filled FieldType for a Required String field. This is used primarily for unit testing.
var TypeStringReqCollection = &FieldType{
	Required:          true,
	Collection:        true,
	CollectionSubtype: TypeStringReq,
}

// TypeIntReq is a pre-filled FieldType for a Required Int field. This is used primarily for unit testing.
var TypeIntReq = &FieldType{
	Kind:              "String",
	Required:          true,
	Collection:        false,
	CollectionSubtype: nil,
}

func ParseFieldType(typ *ast.Type) *FieldType {
	if typ.Elem == nil {
		return &FieldType{
			Kind:       typ.NamedType,
			Required:   typ.NonNull,
			Collection: false,
		}
	} else {
		subtype := ParseFieldType(typ.Elem)
		return &FieldType{
			Required:          typ.NonNull,
			Collection:        true,
			CollectionSubtype: subtype,
		}
	}
}

func (ft FieldType) rawGoType() string {
	switch ft.Kind {
	case "String":
		return "string"
	case "Int":
		return "int"
	case "Long":
		return "int64"
	case "Float":
		return "float64"
	case "Boolean":
		return "bool"
	case "ID":
		return "string"
	default:
		return ft.Kind
	}
}

func (ft FieldType) IsObject() bool {
	if ft.Collection {
		return false
	}
	switch ft.Kind {
	case "String":
		return false
	case "Int":
		return false
	case "Long":
		return false
	case "Float":
		return false
	case "Boolean":
		return false
	case "ID":
		return false
	default:
		return true
	}
}

func (ft FieldType) GoType() string {
	if ft.Collection {
		return "[]" + ft.CollectionSubtype.GoType()
	}

	if ft.Required {
		return ft.rawGoType()
	} else {
		return "*" + ft.rawGoType()
	}
}

func (ft FieldType) GoTypeRequired() string {
	if ft.Collection {
		return "[]" + ft.CollectionSubtype.GoType()
	}

	return ft.rawGoType()
}

func (ft FieldType) AppSyncType() string {
	return ft.Kind
}

func (ft FieldType) DynamoType() string {
	if ft.Collection {
		return "L"
	}
	switch ft.Kind {
	case "String":
		return "S"
	case "Int":
		return "N"
	case "Long":
		return "N"
	case "Float":
		return "N"
	case "Boolean":
		return "BOOL"
	case "ID":
		return "S"
	default:
		return "M"
	}
}

func (ft FieldType) DynamoPointerFunc() string {
	if ft.Collection {
		return "aws.List"
	}
	switch ft.Kind {
	case "String":
		return "aws.String"
	case "Int":
		return "aws.Int"
	case "Long":
		return "aws.Int64"
	case "Float":
		return "aws.Float"
	case "Boolean":
		return "aws.Bool"
	case "ID":
		return "aws.String"
	default:
		log.Printf("Unexpected Kind for DynamoPointerFunc: %s", ft.Kind)
		return ""
	}
}

func (ft FieldType) IsCollectionOfObjects() bool {
	return ft.Collection == true && !ft.CollectionSubtype.Collection && ft.CollectionSubtype.IsObject()
}

// Validate is a validation function for field type object
func (ft FieldType) Validate() error {
	if ft.Kind == "" {
		return fmt.Errorf("field type kind is required")
	}
	//validate that ft.Kind is one of the valid kinds

	if ft.Collection && ft.CollectionSubtype == nil {
		return fmt.Errorf("collection subtype is required for collection field type")
	}

	if ft.CollectionSubtype != nil && !ft.Collection {
		return fmt.Errorf("collection subtype is only valid for collection field type")
	}

	return nil
}
