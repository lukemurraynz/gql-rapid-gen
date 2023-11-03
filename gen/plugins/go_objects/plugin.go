// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package go_objects

import "github.com/mjdrgn/gql-rapid-gen/gen"

type Plugin struct {
}

func (p *Plugin) Name() string {
	return "go_objects"
}

func (p *Plugin) Order() int {
	return -100
}

func (p *Plugin) Tags() []string {
	return []string{
		"go",
	}
}

func init() {
	gen.RegisterPlugin("go_objects", &Plugin{})
}
