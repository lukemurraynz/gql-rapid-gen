module "auth" {
  source = "./auth"

  web-domain      = var.web-domain
  auth-domain     = var.auth-domain
  dns-zone-id     = data.aws_route53_zone.default.zone_id
  acm-cert-arn-us = data.aws_acm_certificate.wildcard-us-east-1.arn
  lib-hash        = data.archive_file.lambda-lib-source.output_base64sha256
  user-table-arn  = module.api.user_table_arn
  client-code     = var.client-code
}