
resource "aws_iam_role" "sns" {
  name = "cognito_sns"

  assume_role_policy = data.aws_iam_policy_document.sns-assume.json
}

resource "aws_iam_role_policy" "sns" {
  role   = aws_iam_role.sns.id
  policy = data.aws_iam_policy_document.sns.json
}

data "aws_iam_policy_document" "sns-assume" {
  statement {
    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["cognito-idp.amazonaws.com"]
    }

    actions = ["sts:AssumeRole"]

    condition {
      test     = "StringEquals"
      values   = [data.aws_caller_identity.me.account_id]
      variable = "aws:SourceAccount"
    }

    #    condition {
    #      test     = "StringEquals"
    #      values   = [aws_cognito_user_pool.users.arn]
    #      variable = "aws:SourceArn"
    #    }
  }
}

data "aws_iam_policy_document" "sns" {
  statement {
    effect    = "Allow"
    actions   = ["sns:publish"]
    resources = ["*"]
  }
}