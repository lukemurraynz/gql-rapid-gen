
data "aws_route53_zone" "default" {
  name = var.cert-domain
}