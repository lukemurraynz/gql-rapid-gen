// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package parser

import (
	"github.com/mjdrgn/gql-rapid-gen/util"
	"github.com/vektah/gqlparser/v2/ast"
	"strings"
)

type ParsedArgumentDef struct {
	Name    string
	Type    *FieldType
	Default Value
}

func (pad *ParsedArgumentDef) NameTitle() string {
	return util.TitleCase(pad.Name)
}

func (pad *ParsedArgumentDef) NameCamel() string {
	return util.CamelCase(pad.Name)
}

func (pad *ParsedArgumentDef) NameUnder() string {
	return util.UnderCase(pad.Name)
}

func (pad *ParsedArgumentDef) NameDash() string {
	return util.DashCase(pad.Name)
}

func (pad *ParsedArgumentDef) GoStructTag() string {
	entries := make([]string, 0, 16)
	if !pad.Type.Required {
		entries = append(
			entries,
			`json:"`+pad.Name+`,omitempty"`,
		)
	} else if pad.Type.Collection {
		entries = append(
			entries,
			`json:"`+pad.Name+`,omitempty"`,
		)
	} else if pad.Type.Required {
		entries = append(
			entries,
			`json:"`+pad.Name+`"`,
		)
	}
	if len(entries) > 0 {
		return "`" + strings.Join(entries, " ") + "`"
	} else {
		return ""
	}
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
