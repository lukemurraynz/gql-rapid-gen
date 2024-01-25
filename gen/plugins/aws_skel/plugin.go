// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package aws_skel

import "github.com/mjdrgn/gql-rapid-gen/gen"

type Plugin struct {
}

func (p *Plugin) Name() string {
	return "aws_skel"
}

func (p *Plugin) Order() int {
	return 0
}

func (p *Plugin) Tags() []string {
	return []string{
		"aws",
	}
}

func init() {
	gen.RegisterPlugin("aws_skel", &Plugin{})
}
