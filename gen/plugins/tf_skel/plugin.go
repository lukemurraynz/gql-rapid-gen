// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package tf_skel

import "github.com/mjdrgn/gql-rapid-gen/gen"

type Plugin struct {
}

func (p *Plugin) Name() string {
	return "tf_skel"
}

func (p *Plugin) Order() int {
	return 0
}

func (p *Plugin) Tags() []string {
	return []string{}
}

func init() {
	gen.RegisterPlugin("tf_skel", &Plugin{})
}
