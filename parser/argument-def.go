// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package parser

import (
	"github.com/vektah/gqlparser/v2/ast"
)

type ParsedArgumentDef struct {
	Name    string
	Type    *FieldType
	Default Value
}

func parseArgumentDef(v *ast.ArgumentDefinition) *ParsedArgumentDef {
	ft := ParseFieldType(v.Type)
	return &ParsedArgumentDef{
		Name:    v.Name,
		Type:    ft,
		Default: parseValue(v.DefaultValue),
	}
}

func parseArgumentDefs(al ast.ArgumentDefinitionList) (args map[string]*ParsedArgumentDef) {
	if al == nil {
		return nil
	}
	args = make(map[string]*ParsedArgumentDef, len(al))
	for _, a := range al {
		args[a.Name] = parseArgumentDef(a)

	}
	return args
}
