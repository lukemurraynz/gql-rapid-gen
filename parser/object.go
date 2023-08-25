package parser

import (
	"github.com/mjdrgn/gql-rapid-gen/util"
	"github.com/vektah/gqlparser/v2/ast"
)

type ParsedObject struct {
	Name        string
	Directives  map[string]*ParsedDirective
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
	for _, f := range po.Fields {
		if f.Name == key {
			return f
		}
	}
	return nil
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
