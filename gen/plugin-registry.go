package gen

import (
	"errors"
	"github.com/mjdrgn/gql-rapid-gen/parser"
	"golang.org/x/exp/slices"
	"log"
)

var plugins = make(map[string]Plugin, 32)

// RegisterPlugin adds a new plugin to the registry, validating it does not already exist
func RegisterPlugin(name string, p Plugin) {
	_, pres := plugins[name]
	if pres {
		panic("Double registration of Plugin: " + name)
	}
	plugins[name] = p
	log.Printf("Plugin Registered: %s", name)
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
		return a.Order() - b.Order()
	})

	for _, p := range plugins {
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
