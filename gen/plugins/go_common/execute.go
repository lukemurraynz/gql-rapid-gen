// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package go_common

import (
	"fmt"
	"github.com/mjdrgn/gql-rapid-gen/gen"
	"github.com/mjdrgn/gql-rapid-gen/parser"
	"golang.org/x/exp/slices"
)

type data struct {
	ProviderNames []string
}

func (p *Plugin) Generate(schema *parser.Schema, output *gen.Output) error {

	names := make([]string, 0, len(schema.Objects))

	for _, o := range schema.Objects {
		if o.HasDirective("go_ignore") {
			continue
		}

		if o.HasDirective("dynamodb") {
			names = append(names, o.NameTitle())
		}
	}

	slices.Sort(names)

	rendered, err := gen.ExecuteTemplate("plugins/go_common/templates/providers.tmpl", data{
		ProviderNames: names,
	})
	if err != nil {
		return fmt.Errorf("failed rendering providers: %w", err)
	}

	of, err := output.AppendOrCreate(gen.GO_DATA_GEN, "providers", rendered)
	if err != nil {
		return fmt.Errorf("failed appending providers: %w", err)
	}
	of.AddExtraData("context", "github.com/aws/aws-sdk-go-v2/config", "github.com/aws/aws-sdk-go-v2/service/dynamodb")

	return nil
}
