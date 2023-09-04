// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package go_lambda

import "github.com/mjdrgn/gql-rapid-gen/parser"

func (p *Plugin) Qualify(schema *parser.Schema) bool {
	for _, m := range schema.Mutation {
		if m.HasDirective("appsync_lambda") {
			return true
		}
	}
	for _, m := range schema.Query {
		if m.HasDirective("appsync_lambda") {
			return true
		}
	}
	return false
}
