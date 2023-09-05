// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package go_dynamodb

import "github.com/mjdrgn/gql-rapid-gen/parser"

func (p *Plugin) Qualify(schema *parser.Schema) bool {
	if len(schema.Objects) == 0 {
		return false
	}
	for _, o := range schema.Objects {
		if o.HasDirective("dynamodb") {
			return true
		}
	}
	return false
}
