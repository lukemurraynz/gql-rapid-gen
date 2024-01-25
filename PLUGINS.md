# GraphQL Rapid Application Generator - Plugins

Plugins represent specific generation capability, and are used to customise the generated code based on what infrastructure elements you are using.

They are organised into Tags - a plugin will only run if all it's tags are in the TagEnable list in configuration.
See [CONFIG.md](./CONFIG.md) for more info on enabling or disabling plugins.

## AWS

### aws_skel
Tags: aws
Provides basic AWS infrastructure and Terraform configuration, required for an AppSync based deployment model.

### tf_appsync_dynamo
Tags: aws
Generates AppSync resolvers for DynamoDB tables, including scan, list, create, read, update, and delete.

### tf_appsync_lambda
Tags: aws
Generates AppSync resolvers and AWS Lambda definitions for mutations and queries that are annotated. Supports multiple languages.

### tf_dynamodb
Tags: aws
Generates AWS DynamoDB table definitions including keys and indexes

## Go

### go_common
Tags: go
Generates Go wrapper code used by other plugins.

### go_dynamodb
Tags: go, aws
Generates a Go data layer to access DynamoDB tables, using the same annotations as the AWS plugins.

### go_enum
Tags: go
Generates Go versions of all GraphQL Enums

### go_lambda
Tags: go
Generates Go data layer elements for GraphQL mutation and query arguments and responses, as well as skeletons the Go repositories for each Lambda function including main and go.mod

### go_objects
Tags: go
Generates Go structs for GraphQL Types and Input Types, including basic validation and nesting.

### go_union
Tags: go
NOTE: Currently broken and disabled.
Generates Go types for GraphQL Unions including helper logic.

## Terraform

### tf_skel
Tags: NONE (always enabled)

Provides the basic Terraform skeleton required for all projects. Should not be disabled without a very good reason.

