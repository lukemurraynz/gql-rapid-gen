// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package go_objects

import "github.com/mjdrgn/gql-rapid-gen/gen"

type Plugin struct {
}

func (p *Plugin) Name() string {
	return "tf_dynamodb"
}

func (p *Plugin) Order() int {
	return -100
}

func init() {
	gen.RegisterPlugin("tf_dynamodb", &Plugin{})
}
