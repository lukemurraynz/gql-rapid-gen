
data "aws_acm_certificate" "wildcard" {
  domain      = var.cert-domain
  statuses    = ["ISSUED"]
  most_recent = true
}

data "aws_acm_certificate" "wildcard-us-east-1" {
  provider    = aws.us-east-1
  domain      = var.cert-domain
  statuses    = ["ISSUED"]
  most_recent = true
}