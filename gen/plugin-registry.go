// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package gen

import (
	"errors"
	"fmt"
	"github.com/mjdrgn/gql-rapid-gen/parser"
	"golang.org/x/exp/slices"
)

var plugins = make(map[string]Plugin, 32)

// RegisterPlugin adds a new plugin to the registry, validating it does not already exist
func RegisterPlugin(name string, p Plugin) {
	_, pres := plugins[name]
	if pres {
		panic("Double registration of Plugin: " + name)
	}
	if p.Name() != name {
		panic(fmt.Sprintf("Name mismatch in Plugin: %s vs %s", p.Name(), name))
	}
	plugins[name] = p
}

// ListPlugins returns the names of all plugins registered
func ListPlugins() []string {
	ret := make([]string, 0, len(plugins))
	for k, _ := range plugins {
		ret = append(ret, k)
	}
	return ret
}

// QualifySchema runs all registered Plugin entities against a given Schema and returns those that qualify
func QualifySchema(schema *parser.Schema) (ret []Plugin) {
	for _, p := range plugins {
		if p.Qualify(schema) {
			ret = append(ret, p)
		}
	}
	return ret
}

// ExecuteSchema runs a list of Plugin entities against a given Schema. It assumes they have already been qualified.
func ExecuteSchema(list []Plugin, schema *parser.Schema, output *Output) error {
	var errs []error

	slices.SortFunc(list, func(a Plugin, b Plugin) int {
		if a.Order() == b.Order() {
			if a.Name() < b.Name() {
				return -1
			} else {
				return 1
			}
		}
		return b.Order() - a.Order()
	})

	for _, p := range list {
		err := p.Generate(schema, output)
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	} else {
		return nil
	}
}
