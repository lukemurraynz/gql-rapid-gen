// Copyright (c) 2024 under the MIT license per gql-rapid-gen/LICENSE.MD

package azure_skel

import "github.com/mjdrgn/gql-rapid-gen/gen"

// Plugin struct represents the Azure skeleton plugin.
type Plugin struct {
}

// Name returns the name of the plugin.
func (p *Plugin) Name() string {
    return "azure_skel"
}

// Order returns the order in which the plugin should be executed.
// Lower numbers indicate higher priority.
func (p *Plugin) Order() int {
    return 0
}

// Tags returns a list of tags associated with the plugin.
func (p *Plugin) Tags() []string {
    return []string{
        "azure",
    }
}

// init function registers the Azure skeleton plugin with the gen package.
func init() {
    gen.RegisterPlugin("azure_skel", &Plugin{})
}