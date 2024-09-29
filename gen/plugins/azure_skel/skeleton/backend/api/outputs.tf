output "graphql-api-url" {
  value = aws_appsync_graphql_api.backend.uris["GRAPHQL"]
}

output "user_table_arn" {
  value = aws_dynamodb_table.user.arn
}