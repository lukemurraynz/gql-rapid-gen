// Skeleton Code Copyright (c) 2023 under the MIT license from gql-rapid-gen

//go:build testing

package config

// Env will be filled at compile time with the environment code being deployed into
// e.g.: go build -ldflags "-X lib/config.Env=dev"
var Env = "dev"

func init() {
	// No Env check on testing
}
