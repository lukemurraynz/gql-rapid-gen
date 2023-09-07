// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package parser

import (
	"fmt"
	"github.com/vektah/gqlparser/v2/ast"
)

type ParsedArgument struct {
	Name  string
	Value Value
}

func (pa *ParsedArgument) Validate() error {
	if pa.Name == "" {
		return fmt.Errorf("name is required")
	}
	return nil
}

func parseArgument(v *ast.Argument) *ParsedArgument {
	return &ParsedArgument{
		Name:  v.Name,
		Value: parseValue(v.Value),
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

// ArgumentsFromMap generates a ParsedArgument map based on a Go map - this should be used only for testing purposes
func ArgumentsFromMap(in map[string]any) (args map[string]*ParsedArgument) {
	args = make(map[string]*ParsedArgument, len(in))
	for k, v := range in {
		var tv Value
		switch v.(type) {
		case string:
			tv = &valMarshallable{val: v.(string)}
		case []byte:
			tv = &valMarshallable{val: string(v.([]byte))}
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			tv = &valRaw{val: fmt.Sprintf("%d", v)}
		case float32, float64:
			tv = &valRaw{val: fmt.Sprintf("%f", v)}
		case bool:
			tv = &valBool{val: v.(bool)}
		default:
			panic("unexpected argument type for " + k)
		}

		args[k] = &ParsedArgument{
			Name:  k,
			Value: tv,
		}
	}
	return args
}
