// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package tf_appsync_dynamo

import "github.com/mjdrgn/gql-rapid-gen/parser"

func (p *Plugin) Qualify(schema *parser.Schema) bool {
	for _, o := range schema.Objects {
		if o.HasDirective("dynamodb") {
			return true
		}
	}
	return false
}
