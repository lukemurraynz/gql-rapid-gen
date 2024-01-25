
resource "aws_cognito_user_pool" "users" {
  name = "${var.client-code}-${terraform.workspace}"

  auto_verified_attributes = ["email", "phone_number"]

  alias_attributes = ["email", "phone_number", "preferred_username"]

  mfa_configuration          = "OPTIONAL"
  sms_authentication_message = "Your code is {####}"

  sms_configuration {
    external_id    = "${var.client-code}"
    sns_caller_arn = aws_iam_role.sns.arn
    sns_region     = "ap-southeast-2"
  }

  software_token_mfa_configuration {
    enabled = true
  }

  schema {
    name                = "name"
    attribute_data_type = "String"
    mutable             = true
    required            = false

    string_attribute_constraints {
      min_length = "0"
      max_length = "2048"
    }
  }
  schema {
    name                = "preferred_username"
    attribute_data_type = "String"
    mutable             = true
    required            = false

    string_attribute_constraints {
      min_length = "0"
      max_length = "2048"
    }
  }
  schema {
    name                = "email"
    attribute_data_type = "String"
    mutable             = true
    required            = false

    string_attribute_constraints {
      min_length = "0"
      max_length = "2048"
    }
  }
  schema {
    name                = "phone_number"
    attribute_data_type = "String"
    mutable             = true
    required            = false

    string_attribute_constraints {
      min_length = "0"
      max_length = "2048"
    }
  }

  lambda_config {
    //pre_sign_up = aws_lambda_function.cognito-setup.arn
    post_authentication = aws_lambda_function.cognito-login.arn
  }
}

resource "aws_cognito_user_group" "admins" {
  name         = "Admins"
  user_pool_id = aws_cognito_user_pool.users.id
}

resource "aws_cognito_user_pool_domain" "auth" {
  user_pool_id    = aws_cognito_user_pool.users.id
  domain          = var.auth-domain
  certificate_arn = var.acm-cert-arn-us
}

resource "aws_route53_record" "auth-A" {
  zone_id = var.dns-zone-id
  name    = var.auth-domain
  type    = "A"
  alias {
    evaluate_target_health = false
    name                   = aws_cognito_user_pool_domain.auth.cloudfront_distribution_arn
    zone_id                = "Z2FDTNDATAQYW2"
  }
}

resource "aws_route53_record" "auth-AAAA" {
  zone_id = var.dns-zone-id
  name    = var.auth-domain
  type    = "AAAA"
  alias {
    evaluate_target_health = false
    name                   = aws_cognito_user_pool_domain.auth.cloudfront_distribution_arn
    zone_id                = "Z2FDTNDATAQYW2"
  }
}

resource "aws_cognito_user_pool_client" "web" {
  name         = "${var.client-code}-web-${terraform.workspace}"
  user_pool_id = aws_cognito_user_pool.users.id

  generate_secret = false

  callback_urls = ["https://${var.web-domain}/", "http://localhost:3000/", "http://localhost:5173/"]
  logout_urls   = ["https://${var.web-domain}/", "http://localhost:3000/", "http://localhost:5173/"]

  allowed_oauth_flows_user_pool_client = true

  allowed_oauth_flows  = ["code", "implicit"]
  allowed_oauth_scopes = ["email", "phone", "openid", "profile", "aws.cognito.signin.user.admin"]
  explicit_auth_flows  = ["ALLOW_CUSTOM_AUTH", "ALLOW_REFRESH_TOKEN_AUTH", "ALLOW_USER_SRP_AUTH", "ALLOW_USER_PASSWORD_AUTH", "ALLOW_ADMIN_USER_PASSWORD_AUTH"]

  access_token_validity  = 24
  id_token_validity      = 24
  refresh_token_validity = 30

  supported_identity_providers = [
    "COGNITO"
  ]
}

output "web-cognito-app-client-id" {
  value = aws_cognito_user_pool_client.web.id
}

output "web-cognito-user-pool-id" {
  value = aws_cognito_user_pool.users.id
}