
resource "aws_iam_role" "appsync-logs" {
  name = "appsync-logs-${terraform.workspace}"

  assume_role_policy = <<POLICY
{
    "Version": "2012-10-17",
    "Statement": [
        {
        "Effect": "Allow",
        "Principal": {
            "Service": "appsync.amazonaws.com"
        },
        "Action": "sts:AssumeRole"
        }
    ]
}
POLICY
}

resource "aws_iam_role_policy_attachment" "appsync-logs" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSAppSyncPushToCloudWatchLogs"
  role       = aws_iam_role.appsync-logs.name
}

locals {
  raw_schema                         = file("./schema.graphql")
  schema_without_directives_go       = replace(local.raw_schema, "/@go_[a-z+]+(?:\\(.*?\\))?/", "")
  schema_without_directives_appsync  = replace(local.schema_without_directives_go, "/@appsync_[a-z_+]+(?:\\([\\s\\S]*?\\))?/", "")
  schema_without_directives_dynamodb = replace(local.schema_without_directives_appsync, "/@dynamodb(?:_[a-z_+]+)?(?:\\(.*?\\))?/", "")
  schema_without_directives_crud     = replace(local.schema_without_directives_dynamodb, "/@crud(?:_[a-z_+]+)?(?:\\(.*?\\))?/", "")
  schema_without_comments            = replace(local.schema_without_directives_crud, "/(?m)##? .*$/", "")
}

#resource "local_file" "schema" {
#  content = local.schema_without_comments
#  filename = "${path.module}/schema_rendered.graphql"
#}

resource "aws_appsync_graphql_api" "backend" {
  authentication_type = "AMAZON_COGNITO_USER_POOLS"
  name                = "backend-${terraform.workspace}"

  user_pool_config {
    aws_region     = "ap-southeast-2"
    user_pool_id   = var.cognito-pool-id
    default_action = "ALLOW"
  }

  log_config {
    cloudwatch_logs_role_arn = aws_iam_role.appsync-logs.arn
    field_log_level          = "ERROR"
  }

  schema = local.schema_without_comments
}
