// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package tf_appsync_lambda

import "github.com/mjdrgn/gql-rapid-gen/parser"

func (p *Plugin) Qualify(schema *parser.Schema) bool {
	for _, o := range schema.Mutation {
		if o.HasDirective("appsync_lambda") {
			return true
		}
	}
	for _, o := range schema.Query {
		if o.HasDirective("appsync_lambda") {
			return true
		}
	}
	return false
}
