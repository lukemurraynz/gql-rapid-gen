// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package go_common

import "github.com/mjdrgn/gql-rapid-gen/gen"

type Plugin struct {
}

func (p *Plugin) Name() string {
	return "go_common"
}

func (p *Plugin) Order() int {
	return 0
}

func (p *Plugin) Tags() []string {
	return []string{
		"go",
	}
}

func init() {
	gen.RegisterPlugin("go_common", &Plugin{})
}
