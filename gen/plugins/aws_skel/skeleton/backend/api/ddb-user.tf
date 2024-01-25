
resource "aws_appsync_resolver" "Query-getMe" {
  api_id      = aws_appsync_graphql_api.backend.id
  type        = "Query"
  field       = "getMe"
  data_source = aws_appsync_datasource.user.name

  request_template = <<EOF
{
    "version": "2017-02-28",
    "operation": "GetItem",
    "key": {
        "id": $util.dynamodb.toDynamoDBJson($util.str.toLower($ctx.identity.claims["cognito:username"])),
    }
}
EOF

  response_template = <<EOF
$util.toJson($ctx.result)
EOF
}