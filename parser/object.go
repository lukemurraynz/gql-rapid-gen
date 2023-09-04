// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package parser

import (
	"fmt"
	"github.com/mjdrgn/gql-rapid-gen/util"
	"github.com/vektah/gqlparser/v2/ast"
)

type ParsedObject struct {
	Name        string
	Directives  map[string][]*ParsedDirective
	Description string
	Fields      []*ParsedField
	Interfaces  []string
}

func (po *ParsedObject) NameTitle() string {
	return util.TitleCase(po.Name)
}

func (po *ParsedObject) NameCamel() string {
	return util.CamelCase(po.Name)
}

func (po *ParsedObject) NameUnder() string {
	return util.UnderCase(po.Name)
}

func (po *ParsedObject) NameDash() string {
	return util.DashCase(po.Name)
}

func (po *ParsedObject) Field(key string) *ParsedField {
	if key == "" {
		return nil
	}
	for _, f := range po.Fields {
		if f.Name == key {
			return f
		}
	}
	return nil
}

func (po *ParsedObject) SingleDirective(key string) *ParsedDirective {
	v, ok := po.Directives[key]
	if !ok {
		return nil
	}
	if len(v) == 0 {
		panic(fmt.Errorf("'%s' had '%s' directive found but len 0", po.Name, key))
	}
	if len(v) > 1 {
		panic(fmt.Errorf("object '%s' had duplicate '%s' directives", po.Name, key))
	}
	return v[0]
}

func (po *ParsedObject) HasDirective(key string) bool {
	_, ok := po.Directives[key]
	return ok
}

func parseObject(def *ast.Definition) (ret *ParsedObject) {
	return &ParsedObject{
		Name:        def.Name,
		Directives:  parseDirectives(def.Directives),
		Description: def.Description,
		Fields:      parseFields(def.Fields),
		Interfaces:  def.Interfaces,
	}
}
