package parser

import (
	"fmt"
	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
	"os"
	"path/filepath"
)

type Schema struct {
	Objects      map[string]*ParsedObject
	InputObjects map[string]*ParsedObject
	Enums        map[string]*ParsedEnum

	Query        map[string]*ParsedField
	Mutation     map[string]*ParsedField
	Subscription map[string]*ParsedField
}

func Parse(schemaFiles []string) (output *Schema, err error) {
	sources := make([]*ast.Source, 0, len(schemaFiles))
	for _, f := range schemaFiles {
		f := f
		abs, err := filepath.Abs(f)
		if err != nil {
			return nil, fmt.Errorf("failed resolving filepath for %s: %w", f, err)
		}
		rawSchema, err := os.ReadFile(abs)
		if err != nil {
			return nil, fmt.Errorf("failed reading %s: %w", f, err)
		}
		sources = append(sources, &ast.Source{
			Name:    f,
			Input:   string(rawSchema),
			BuiltIn: false,
		})
	}

	output = &Schema{
		Objects:      make(map[string]*ParsedObject),
		InputObjects: make(map[string]*ParsedObject),
		Enums:        make(map[string]*ParsedEnum),
		Query:        make(map[string]*ParsedField),
		Mutation:     make(map[string]*ParsedField),
		Subscription: make(map[string]*ParsedField),
	}

	schema, err := gqlparser.LoadSchema(sources...)
	if err != nil {
		return nil, fmt.Errorf("failed parsing schema: %w", err)
	}

	for name, t := range schema.Types {
		switch t.Kind {
		case ast.Object:
			output.Objects[name] = parseObject(t)
		case ast.InputObject:
			output.InputObjects[name] = parseObject(t)
		case ast.Enum:
			output.Enums[name] = parseEnum(t)
			// TODO Union
			// TODO Interface
			// TODO Scalar
		}
	}

	return output, nil
}
