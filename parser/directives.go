// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package parser

import "github.com/vektah/gqlparser/v2/ast"

type ParsedDirective struct {
	Name      string
	Arguments map[string]*ParsedArgument
}

func (pd *ParsedDirective) ArgIsNull(key string) bool {
	arg, ok := pd.Arguments[key]
	if !ok {
		return true
	}
	_, ok = arg.Value.(*valNull)
	return ok
}

func (pd *ParsedDirective) ArgGo(key string) string {
	arg, ok := pd.Arguments[key]
	if !ok {
		return "nil"
	}
	return arg.Value.GoString()
}

func (pd *ParsedDirective) ArgJS(key string) string {
	arg, ok := pd.Arguments[key]
	if !ok {
		return "undefined"
	}
	return arg.Value.JSString()
}

func (pd *ParsedDirective) ArgHCL(key string) string {
	arg, ok := pd.Arguments[key]
	if !ok {
		return "null"
	}
	return arg.Value.HCLString()
}

func parseDirective(d *ast.Directive) *ParsedDirective {
	return &ParsedDirective{
		Name:      d.Name,
		Arguments: parseArguments(d.Arguments),
	}
}

func parseDirectives(ds ast.DirectiveList) (ret map[string]*ParsedDirective) {
	if ds == nil {
		return nil
	}
	ret = make(map[string]*ParsedDirective, len(ds))
	for _, d := range ds {
		ret[d.Name] = parseDirective(d)
	}
	return ret
}
