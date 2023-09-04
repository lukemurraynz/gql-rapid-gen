// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package parser

import "github.com/vektah/gqlparser/v2/ast"

type ParsedArgument struct {
	Name    string
	Value   Value
	Comment string
}

func parseArgument(v *ast.Argument) *ParsedArgument {
	return &ParsedArgument{
		Name:    v.Name,
		Value:   parseValue(v.Value),
		Comment: v.Comment.Dump(),
	}
}

func parseArguments(al ast.ArgumentList) (args map[string]*ParsedArgument) {
	if al == nil {
		return nil
	}
	args = make(map[string]*ParsedArgument, len(al))
	for _, a := range al {
		args[a.Name] = parseArgument(a)

	}
	return args
}
