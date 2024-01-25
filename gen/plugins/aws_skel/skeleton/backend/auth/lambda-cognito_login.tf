resource "aws_iam_role" "lambda-cognito-login" {
  name = "lambda-cognito-login-${terraform.workspace}"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_policy" "lambda-cognito-login" {
  name = "lambda-cognito-login-${terraform.workspace}"
  path = "/"

  policy = data.aws_iam_policy_document.lambda-cognito-login.json
}

data "aws_iam_policy_document" "lambda-cognito-login" {
  statement {
    actions = [
      "dynamodb:GetItem",
      "dynamodb:PutItem",
      "dynamodb:UpdateItem"
    ]
    effect = "Allow"
    resources = [
      var.user-table-arn
    ]
  }
}

resource "aws_iam_role_policy_attachment" "lambda-cognito-login" {
  policy_arn = aws_iam_policy.lambda-cognito-login.arn
  role       = aws_iam_role.lambda-cognito-login.name
}

resource "aws_iam_role_policy_attachment" "lambda-cognito-login-basic" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  role       = aws_iam_role.lambda-cognito-login.name
}

data "archive_file" "lambda-cognito-login-source" {
  type        = "zip"
  output_path = "${path.module}/lambda/build/src/cognito_login.zip"
  source_dir  = "${path.module}/lambda/cognito_login/"
}

resource "null_resource" "lambda-cognito-login-compile" {
  triggers = {
    source_hash = data.archive_file.lambda-cognito-login-source.output_base64sha256
    lib_hash    = var.lib-hash
    always_run  = timestamp()
  }

  provisioner "local-exec" {
    working_dir = "${path.module}/lambda/cognito_login"
    environment = {
      GOOS    = "linux"
      GOARCH  = "arm64"
      GOFLAGS = "-trimpath"
    }
    command = "go build -ldflags \"-X lib/config.Env=${terraform.workspace}\" -o ../build/${terraform.workspace}/cognito_login/bootstrap -trimpath ."
  }
}

data "archive_file" "lambda-cognito-login" {
  type        = "zip"
  output_path = "${path.module}/lambda/build/${terraform.workspace}/cognito_login_${terraform.workspace}.zip"
  source_file = "${path.module}/lambda/build/${terraform.workspace}/cognito_login/bootstrap"

  depends_on = [
    null_resource.lambda-cognito-login-compile
  ]
}


resource "aws_lambda_function" "cognito-login" {
  function_name = "cognito_login_${terraform.workspace}"
  role          = aws_iam_role.lambda-cognito-login.arn

  runtime       = "provided.al2"
  architectures = ["arm64"]

  handler = "cognitoLogin"

  memory_size = 128
  timeout     = 4

  environment {
    variables = {
      ENVIRONMENT = terraform.workspace
    }
  }

  filename         = data.archive_file.lambda-cognito-login.output_path
  source_code_hash = data.archive_file.lambda-cognito-login.output_base64sha256
}

resource "aws_lambda_permission" "cognito-login" {
  statement_id  = "AllowExecutionFromCognito"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.cognito-login.function_name
  principal     = "cognito-idp.amazonaws.com"
  source_arn    = aws_cognito_user_pool.users.arn
}
