
resource "aws_apigatewayv2_api" "webhooks" {
  name          = "${var.client-code}-webhooks-${terraform.workspace}"
  protocol_type = "HTTP"
  cors_configuration {
    allow_methods = ["GET", "POST", "PUT"]
    allow_origins = ["https://${var.web-domain}"]
    max_age       = 3600
  }
}

resource "aws_apigatewayv2_stage" "webhooks-default" {
  api_id      = aws_apigatewayv2_api.webhooks.id
  name        = "${var.client-code}-webhooks-${terraform.workspace}-default"
  auto_deploy = true

  lifecycle {
    ignore_changes = [
      deployment_id,
      default_route_settings
    ]
  }

}

resource "aws_apigatewayv2_api_mapping" "webhooks" {
  api_id      = aws_apigatewayv2_api.webhooks.id
  domain_name = aws_apigatewayv2_domain_name.webhooks.id
  stage       = aws_apigatewayv2_stage.webhooks-default.id
}

resource "aws_apigatewayv2_domain_name" "webhooks" {
  domain_name = var.webhook-domain

  domain_name_configuration {
    certificate_arn = var.acm-cert-arn-syd
    endpoint_type   = "REGIONAL"
    security_policy = "TLS_1_2"
  }
}

resource "aws_route53_record" "webhooks-A" {
  name    = aws_apigatewayv2_domain_name.webhooks.domain_name
  type    = "A"
  zone_id = var.dns-zone-id

  alias {
    name                   = aws_apigatewayv2_domain_name.webhooks.domain_name_configuration[0].target_domain_name
    zone_id                = aws_apigatewayv2_domain_name.webhooks.domain_name_configuration[0].hosted_zone_id
    evaluate_target_health = false
  }
}

resource "aws_route53_record" "webhooks-AAAA" {
  name    = aws_apigatewayv2_domain_name.webhooks.domain_name
  type    = "AAAA"
  zone_id = var.dns-zone-id

  alias {
    name                   = aws_apigatewayv2_domain_name.webhooks.domain_name_configuration[0].target_domain_name
    zone_id                = aws_apigatewayv2_domain_name.webhooks.domain_name_configuration[0].hosted_zone_id
    evaluate_target_health = false
  }
}

output "webhooks-url" {
  value = "https://${aws_apigatewayv2_domain_name.webhooks.domain_name}/"
}