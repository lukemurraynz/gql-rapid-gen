// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package parser

import (
	_ "embed"
	"errors"
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
	Unions       map[string]*ParsedUnion

	Query        map[string]*ParsedField
	Mutation     map[string]*ParsedField
	Subscription map[string]*ParsedField
}

func (s *Schema) Validate() error {
	errs := make([]error, 0)

	for k, v := range s.Objects {
		err := v.Validate()
		if err != nil {
			errs = append(errs, fmt.Errorf("failed validating Object '%s': %w", k, err))
		}
	}

	for k, v := range s.InputObjects {
		err := v.Validate()
		if err != nil {
			errs = append(errs, fmt.Errorf("failed validating InputObject '%s': %w", k, err))
		}
	}

	for k, v := range s.Enums {
		err := v.Validate()
		if err != nil {
			errs = append(errs, fmt.Errorf("failed validating Enum '%s': %w", k, err))
		}
	}

	// TODO unions

	for k, v := range s.Query {
		err := v.Validate()
		if err != nil {
			errs = append(errs, fmt.Errorf("failed validating Query '%s': %w", k, err))
		}
	}

	for k, v := range s.Mutation {
		err := v.Validate()
		if err != nil {
			errs = append(errs, fmt.Errorf("failed validating Mutation '%s': %w", k, err))
		}
	}

	for k, v := range s.Subscription {
		err := v.Validate()
		if err != nil {
			errs = append(errs, fmt.Errorf("failed validating Subscription '%s': %w", k, err))
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

//go:embed def.graphql
var defSchema string
var defSource = &ast.Source{
	Name:    "def.graphql",
	Input:   defSchema,
	BuiltIn: false,
}

func Parse(schemaFiles []string) (output *Schema, err error) {
	sources := make([]*ast.Source, 0, len(schemaFiles)+1)
	sources = append(sources, defSource)
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
		if t.BuiltIn {
			continue
		}
		if name == "Query" || name == "Subscription" || name == "Mutation" {
			continue
		}
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

	if schema.Query != nil {
		for _, t := range schema.Query.Fields {
			output.Query[t.Name] = parseField(t)
		}
	}

	if schema.Mutation != nil {
		for _, t := range schema.Mutation.Fields {
			output.Mutation[t.Name] = parseField(t)
		}
	}

	if schema.Subscription != nil {
		for _, t := range schema.Subscription.Fields {
			output.Subscription[t.Name] = parseField(t)
		}
	}

	return output, nil
}
