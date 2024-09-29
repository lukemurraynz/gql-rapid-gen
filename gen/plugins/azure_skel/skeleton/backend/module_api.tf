module "api" {
  source = "./api"

  client-code      = var.client-code
  web-domain       = var.web-domain
  webhook-domain   = var.webhook-domain
  dns-zone-id      = data.aws_route53_zone.default.zone_id
  acm-cert-arn-us  = data.aws_acm_certificate.wildcard-us-east-1.arn
  acm-cert-arn-syd = data.aws_acm_certificate.wildcard.arn
  cognito-pool-id  = module.auth.pool-id
  lib-hash         = data.archive_file.lambda-lib-source.output_base64sha256
}