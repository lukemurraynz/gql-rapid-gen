// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package parser

import (
	"fmt"
	"github.com/mjdrgn/gql-rapid-gen/util"
	"github.com/vektah/gqlparser/v2/ast"
	"strings"
)

type ParsedField struct {
	Name        string
	Directives  map[string][]*ParsedDirective
	Description string
	Arguments   map[string]*ParsedArgumentDef
	Type        *FieldType
}

func (pf *ParsedField) NameTitle() string {
	return util.TitleCase(pf.Name)
}

func (pf *ParsedField) NameCamel() string {
	return util.CamelCase(pf.Name)
}

func (pf *ParsedField) NameUnder() string {
	return util.UnderCase(pf.Name)
}

func (pf *ParsedField) NameDash() string {
	return util.DashCase(pf.Name)
}

func (pf *ParsedField) HasDirective(key string) bool {
	_, ok := pf.Directives[key]
	return ok
}

func (pf *ParsedField) SingleDirective(key string) *ParsedDirective {
	v, ok := pf.Directives[key]
	if !ok {
		return nil
	}
	if len(v) == 0 {
		panic(fmt.Errorf("'%s' had '%s' directive found but len 0", pf.Name, key))
	}
	if len(v) > 1 {
		panic(fmt.Errorf("field '%s' had duplicate '%s' directives", pf.Name, key))
	}
	return v[0]
}

func (pf *ParsedField) GoStructTag() string {
	entries := make([]string, 0, 16)
	if !pf.Type.Required {
		entries = append(
			entries,
			`json:"`+pf.Name+`,omitempty"`,
			`dynamodbav:"`+pf.Name+`,omitempty"`,
		)
	} else if pf.Type.Collection {
		entries = append(
			entries,
			`json:"`+pf.Name+`,omitempty"`,
			`dynamodbav:"`+pf.Name+`,omitempty"`,
		)
	} else if pf.Type.Required {
		entries = append(
			entries,
			`json:"`+pf.Name+`"`,
			`dynamodbav:"`+pf.Name+`"`,
		)
	}
	if len(entries) > 0 {
		return "`" + strings.Join(entries, " ") + "`"
	} else {
		return ""
	}
}

func parseField(f *ast.FieldDefinition) *ParsedField {
	return &ParsedField{
		Name:        f.Name,
		Directives:  parseDirectives(f.Directives),
		Description: f.Description,
		Arguments:   parseArgumentDefs(f.Arguments),
		Type:        ParseFieldType(f.Type),
	}
}

func parseFields(fl ast.FieldList) []*ParsedField {
	if fl == nil {
		return nil
	}
	fields := make([]*ParsedField, 0, len(fl))
	for _, f := range fl {
		fields = append(fields, parseField(f))
	}
	return fields
}
