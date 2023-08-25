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

func init() {
	gen.RegisterPlugin("go_objects", &Plugin{})
}
