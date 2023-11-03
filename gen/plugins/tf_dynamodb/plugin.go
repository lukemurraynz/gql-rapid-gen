// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package tf_dynamodb

import "github.com/mjdrgn/gql-rapid-gen/gen"

type Plugin struct {
}

func (p *Plugin) Name() string {
	return "tf_dynamodb"
}

func (p *Plugin) Order() int {
	return -1
}

func (p *Plugin) Tags() []string {
	return []string{
		"aws",
	}
}

func init() {
	gen.RegisterPlugin("tf_dynamodb", &Plugin{})
}
