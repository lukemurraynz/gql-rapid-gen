// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package tf_appsync_lambda

import "github.com/mjdrgn/gql-rapid-gen/gen"

type Plugin struct {
}

func (p *Plugin) Name() string {
	return "tf_appsync_lambda"
}

func (p *Plugin) Order() int {
	return -1
}

func init() {
	gen.RegisterPlugin("tf_appsync_lambda", &Plugin{})
}
