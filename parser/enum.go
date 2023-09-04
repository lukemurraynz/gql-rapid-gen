// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package parser

import "github.com/vektah/gqlparser/v2/ast"

type ParsedEnum struct {
	Name        string
	Description string
	Values      []*EnumValue
}

func (e ParsedEnum) ValueString() []string {
	ret := make([]string, 0, len(e.Values))
	for _, v := range e.Values {
		ret = append(ret, v.Name)
	}
	return ret
}

type EnumValue struct {
	Name        string
	Description string
	Directives  map[string]*ParsedDirective
}

func parseEnumValue(def *ast.EnumValueDefinition) (ret *EnumValue) {
	return &EnumValue{
		Name:        def.Name,
		Description: def.Description,
		Directives:  parseDirectives(def.Directives),
	}
}

func parseEnum(def *ast.Definition) (ret *ParsedEnum) {
	vals := make([]*EnumValue, 0, len(def.EnumValues))
	for _, d := range def.EnumValues {
		vals = append(vals, parseEnumValue(d))
	}
	return &ParsedEnum{
		Name:        def.Name,
		Description: def.Description,
		Values:      vals,
	}
}
