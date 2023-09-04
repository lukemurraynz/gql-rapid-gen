// Skeleton Code Copyright (c) 2023 under the MIT license from gql-rapid-gen

//go:build !testing

package config

// if you're encountering this file from a stack trace in your logs when trying to run go:test
// you're probably missing the --tags=testing flag on your run configuration

// Env will be filled at compile time with the environment code being deployed into
// e.g.: go build -ldflags "-X lib/config.Env=dev"
var Env = "invalid"

func init() {
	if Env == "invalid" {
		panic("Env is not set: go build -ldflags \"-X lib/config.Env=dev\"")
	}
}
