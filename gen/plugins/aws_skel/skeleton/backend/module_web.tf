module "web" {
  source = "./web"

  client-code     = var.client-code
  web-domain      = var.web-domain
  cert-domain     = var.cert-domain
  media-domain    = var.media-domain
  dns-zone-id     = data.aws_route53_zone.default.zone_id
  acm-cert-arn-us = data.aws_acm_certificate.wildcard-us-east-1.arn
  graphql-api-url = module.api.graphql-api-url
}