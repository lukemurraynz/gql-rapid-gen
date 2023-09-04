// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package go_lambda

import "github.com/mjdrgn/gql-rapid-gen/gen"

type Plugin struct {
}

func (p *Plugin) Name() string {
	return "go_lambda"
}

func (p *Plugin) Order() int {
	return 0
}

func init() {
	gen.RegisterPlugin("go_lambda", &Plugin{})
}
