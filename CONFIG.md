# GraphQL Rapid Application Generator - Configuration File

The per-project config file is called `gql-rapid-gen.json` and is located in the root of an individual project.

This must be created prior to running bootstrap or generate, and can be modified as needed thereafter.

## Field Reference
The included fields are as follows

### OutputDirectory
Type: String
Required: Yes

This should be the directory where generated content is output - used if you are nesting a generated project within a larger repo and want the config somewhere else.
This defaults to `.` and should almost always be this value.

### TagEnable
Type: Array of String
Required: Yes

This is the 'tags' of modules that you wish to use for code generation.
A full list of supported tags is in [PLUGINS.md](./PLUGINS.md).
You are almost always going to want at least two tags - a cloud platform, and a backend programming language.
For example, the recommended defaults are `aws` and `go`.
If you do not specify at least one cloud platform and at least one backend language, the generated code will likely not deploy.

### PluginEnable
Type: Array of String
Required: No

If specified, runs only the plugins specified in this list. Used if you need tighter customisation than Tags.
NOTE: This is not recommended unless you are very careful. TagEnable and PluginDisable are better choices for most users.

### PluginDisable
Type: Array of String
Required: No

If specified, disables the specified plugin names from running. Used if you want to manually control specific elements of generation, or don't use specific features.
Wherever possible Tags should be used to limit generation to only what is desired, but this option remains for extra customisation.

### SchemaFiles
Type: Array of String
Required: Yes

The list of .graphql files that contain schema - for use when you have multiple schemas that may feed into your overall project.
The default is `./backend/schema.graphql` which lines up with the default generated schema when bootstrap is used.
