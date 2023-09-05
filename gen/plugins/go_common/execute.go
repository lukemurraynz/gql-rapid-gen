// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package go_objects

import (
	"fmt"
	"github.com/mjdrgn/gql-rapid-gen/gen"
	"github.com/mjdrgn/gql-rapid-gen/parser"
)

type data struct {
	ProviderNames []string
}

func (p *Plugin) Generate(schema *parser.Schema, output *gen.Output) error {

	names := make([]string, 0, len(schema.Objects))

	for _, o := range schema.Objects {
		if o.HasDirective("dynamodb") {
			names = append(names, o.NameTitle())
		}
	}

	rendered, err := gen.ExecuteTemplate("plugins/go_common/templates/providers.tmpl", data{
		ProviderNames: names,
	})
	if err != nil {
		return fmt.Errorf("failed rendering providers: %w", err)
	}

	_, err = output.AppendOrCreate(gen.GO_DATA_GEN, "providers", rendered)
	if err != nil {
		return fmt.Errorf("failed appending providers: %w", err)
	}

	return nil
}
