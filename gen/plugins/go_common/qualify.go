package go_objects

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
}
