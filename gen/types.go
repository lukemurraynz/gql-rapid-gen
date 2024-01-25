// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package gen

const (
	RAW_SKEL        = "raw_skel"
	GO_DATA_GEN     = "go_data_gen"
	GO_DATA_SKEL    = "go_data_skel"
	GO_LAMBDA_SKEL  = "go_lambda_skel"
	GO_LAMBDA_MOD   = "go_lambda_mod"
	TS_FRONTEND_GEN = "ts_frontend_gen"
	TF_API_GEN      = "tf_api_gen"
	TF_API_SKEL     = "tf_api_skel"
)

var Types = map[string]Type{
	RAW_SKEL: {
		Header:          ``,
		Extension:       "",
		Format:          nil,
		GenerateIfEmpty: false,
		Overwrite:       false,
		DirectoryPrefix: "",
		ExtraDataRender: nil,
	},
	GO_DATA_GEN: {
		Header: `package data
// Code generated by gql-rapid-gen DO NOT EDIT.
// This file is derivative of gql-rapid-gen templates.
// Template Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD.
`,
		Extension:       ".gen.go",
		Format:          formatGo,
		GenerateIfEmpty: false,
		Overwrite:       true,
		DirectoryPrefix: "lib/data/",
		ExtraDataRender: goImports,
	},
	GO_DATA_SKEL: {
		Header: `package data
// Skeleton generated by gql-rapid-gen.
// This file is derivative of gql-rapid-gen templates.
// Template Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD.
`,
		Extension:       ".go",
		Format:          formatGo,
		GenerateIfEmpty: true,
		Overwrite:       false,
		DirectoryPrefix: "lib/data/",
		ExtraDataRender: goImports,
	},
	GO_LAMBDA_SKEL: {
		Header: `package main
// Skeleton generated by gql-rapid-gen.
// This file is derivative of gql-rapid-gen templates.
// Template Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD.
`,
		Extension:       ".go",
		Format:          formatGo,
		GenerateIfEmpty: true,
		Overwrite:       false,
		DirectoryPrefix: "backend/api/lambda/",
	},
	GO_LAMBDA_MOD: {
		Header:          "",
		Extension:       ".mod",
		Format:          nil,
		GenerateIfEmpty: false,
		Overwrite:       false,
		DirectoryPrefix: "backend/api/lambda/",
	},
	TS_FRONTEND_GEN: {
		Header: `
// Code generated by gql-rapid-gen DO NOT EDIT.
// This file is derivative of gql-rapid-gen templates.
// Template Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD.
`,
		Extension:       ".gen.tsx",
		GenerateIfEmpty: false,
		Overwrite:       true,
		DirectoryPrefix: "frontend/gen/",
		ExtraDataRender: tsImports,
	},
	TF_API_GEN: {
		Header: `
// Code generated by gql-rapid-gen DO NOT EDIT.
// This file is derivative of gql-rapid-gen templates.
// Template Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD.
`,
		Extension:       ".gen.tf",
		GenerateIfEmpty: false,
		Overwrite:       true,
		DirectoryPrefix: "backend/api/",
	},
	TF_API_SKEL: {
		Header: `
// Skeleton generated by gql-rapid-gen.
// This file is derivative of gql-rapid-gen templates.
// Template Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD.
`,
		Extension:       ".tf",
		GenerateIfEmpty: true,
		Overwrite:       false,
		DirectoryPrefix: "backend/api/",
	},
}

type Type struct {
	Header          string
	Extension       string
	DirectoryPrefix string
	Format          func(string) (string, error)
	ExtraDataRender func(file *OutputFile) error
	GenerateIfEmpty bool
	Overwrite       bool
}
