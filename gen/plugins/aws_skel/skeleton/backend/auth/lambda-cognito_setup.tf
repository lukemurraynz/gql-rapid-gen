resource "aws_iam_role" "lambda-cognito-setup" {
  name = "lambda-cognito-setup-${terraform.workspace}"

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

resource "aws_iam_policy" "lambda-cognito-setup" {
  name = "lambda-cognito-setup-${terraform.workspace}"
  path = "/"

  policy = data.aws_iam_policy_document.lambda-cognito-setup.json
}

data "aws_iam_policy_document" "lambda-cognito-setup" {
  statement {
    actions = [
      "dynamodb:GetItem",
      "dynamodb:PutItem"
    ]
    effect = "Allow"
    resources = [
      var.user-table-arn
    ]
  }
}

resource "aws_iam_role_policy_attachment" "lambda-cognito-setup" {
  policy_arn = aws_iam_policy.lambda-cognito-setup.arn
  role       = aws_iam_role.lambda-cognito-setup.name
}

resource "aws_iam_role_policy_attachment" "lambda-cognito-setup-basic" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  role       = aws_iam_role.lambda-cognito-setup.name
}

data "archive_file" "lambda-cognito-setup-source" {
  type        = "zip"
  output_path = "${path.module}/lambda/build/src/cognito_setup.zip"
  source_dir  = "${path.module}/lambda/cognito_setup/"
}

resource "null_resource" "lambda-cognito-setup-compile" {
  triggers = {
    source_hash = data.archive_file.lambda-cognito-setup-source.output_base64sha256
    lib_hash    = var.lib-hash
    always_run  = timestamp()
  }

  provisioner "local-exec" {
    working_dir = "${path.module}/lambda/cognito_setup"
    environment = {
      GOOS    = "linux"
      GOARCH  = "arm64"
      GOFLAGS = "-trimpath"
    }
    command = "go build -ldflags \"-X lib/config.Env=${terraform.workspace}\" -o ../build/${terraform.workspace}/cognito_setup/bootstrap -trimpath ."
  }
}

data "archive_file" "lambda-cognito-setup" {
  type        = "zip"
  output_path = "${path.module}/lambda/build/${terraform.workspace}/cognito_setup_${terraform.workspace}.zip"
  source_file = "${path.module}/lambda/build/${terraform.workspace}/cognito_setup/bootstrap"

  depends_on = [
    null_resource.lambda-cognito-setup-compile
  ]
}


resource "aws_lambda_function" "cognito-setup" {
  function_name = "cognito_setup_${terraform.workspace}"
  role          = aws_iam_role.lambda-cognito-setup.arn

  runtime       = "provided.al2"
  architectures = ["arm64"]

  handler = "cognitoSetup"

  memory_size = 128
  timeout     = 4

  environment {
    variables = {
      ENVIRONMENT = terraform.workspace
    }
  }

  filename         = data.archive_file.lambda-cognito-setup.output_path
  source_code_hash = data.archive_file.lambda-cognito-setup.output_base64sha256
}

resource "aws_lambda_permission" "cognito-setup" {
  statement_id  = "AllowExecutionFromCognito"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.cognito-setup.function_name
  principal     = "cognito-idp.amazonaws.com"
  source_arn    = aws_cognito_user_pool.users.arn
}
