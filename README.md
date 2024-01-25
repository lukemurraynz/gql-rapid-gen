# GraphQL Rapid Application Generator

This tool is designed as a framework for rapid cloud-native application development using a schema first approach.
While currently only AWS is supported, the intention is to support multiple clouds.

A plugin and tag system is used allowing customisation of exactly which platforms and features are desired for your use case.

## System Requirements
To run the generator, the following is required:
- Linux or Mac OS X (WSL2 is recommended for use on Windows)
- Go >1.21
- Node.JS >v18
- pnpm
- bash or zsh

## Configuration
Primary configuration is through a gql-rapid-gen.json file located in the root of your project.
An example default config that should get most projects started is:
```json
{
	"OutputDirectory": ".",
	"TagEnable": [
		"aws",
		"go"
	],
	"SchemaFiles": [
		"./backend/schema.graphql"
	]
}
```
See [CONFIG.md](./CONFIG.md) for more details.

## Plugins & Tags
See [PLUGINS.md](./PLUGINS.md) for a list of all plugins and tags available.

## Quick Start
1. Create a new empty directory for your project
2. Initialise a Git repository using git init - this guide assumes a basic knowledge of Git already
3. Create a gql-rapid-gen.json file as specified in the Configuration section
4. Commit to git
5. Run "`go run github.com/mjdrgn/gql-rapid-gen bootstrap`" to initialise the skeleton structure
6. Commit to Git
7. Run "`pnpm install`" to set up GraphQL/TypeScript related dependencies
8. Run "`go get`" to set up Go related dependencies
9. Commit to git
10. Run "`go generate`" to complete initial generation and setup
11. Commit to git
12. From the `backend` directory, run `./bootstrap-providers.sh ap-southeast-2 terraform-state-my-project-name my-project-name` - replacing the parameters as required (AWS Region, S3 Bucket Name, Terraform key prefix)
13. Run `terraform init` to install required terraform modules
14. Run `terraform workspace select dev` to ensure you are on the dev environment workspace

At this point, you can deploy your project using standard Terraform commands in the backend folder.
The React frontend can be built and deployed from the Frontend folder.

## Deployment
TODO

## Schema Directives
This tool is based around schema driven development - the majority of infrastructure and code are created from the GraphQL Schema files.

As such, we use a whole set of custom GraphQL directives to inform the tool.

These are documented in [SCHEMA.md](./SCHEMA.md) along with all available options.

## Non-Schema Generation
TODO

## Generating using a local version of the tool
If you want to clone the tool locally, for instance to customise it or for development purposes, a slightly different process is required.
1. After creating your empty project directory, create a go.work file that has use statements for the current directory and the directory you have cloned gql-rapid-gen to
2. Compile the tool using a command similar to the following: `go build -C ../gql-rapid-gen -o ../MY-PROJECT-DIRECTORY/gql-rapid-gen .`
3. Run the tool from the project directory, for example: `./gql-rapid-gen bootstrap`
4. Update generate.go to use the local version
4. You will need to re-run the compile line if you update the code

## Extending the Tool
Additional functionality can be built with new Plugin modules within the gen/plugins directory.
Further documentation on this development process is planned in future.

