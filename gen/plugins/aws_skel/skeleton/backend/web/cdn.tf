
resource "aws_s3_bucket" "cdn" {
  bucket = "${var.client-code}-cdn-${terraform.workspace}"
}

resource "aws_s3_bucket_versioning" "cdn" {
  bucket = aws_s3_bucket.cdn.bucket

  versioning_configuration {
    status = "Enabled"
  }
}

resource "aws_s3_bucket_website_configuration" "cdn" {
  bucket = aws_s3_bucket.cdn.bucket

  index_document {
    suffix = "index.html"
  }
}

resource "aws_s3_bucket_cors_configuration" "cdn" {
  bucket = aws_s3_bucket.cdn.bucket

  cors_rule {
    allowed_methods = [
      "PUT",
      "POST",
      "DELETE",
      "GET",
    ]
    allowed_origins = [
      "http://localhost:3000",
      "https://${var.web-domain}",
    ]
    allowed_headers = [
      "*"
    ]
    max_age_seconds = 300
  }
}

resource "aws_cloudfront_origin_access_identity" "cdn" {

}

data "aws_iam_policy_document" "cdn_s3_policy" {
  statement {
    actions   = ["s3:GetObject"]
    resources = ["${aws_s3_bucket.cdn.arn}/*"]

    principals {
      type        = "AWS"
      identifiers = [aws_cloudfront_origin_access_identity.cdn.iam_arn]
    }
  }
  statement {
    sid     = "AllowSSLRequestsOnly"
    actions = ["s3:*"]
    condition {
      test     = "Bool"
      values   = ["false"]
      variable = "aws:SecureTransport"
    }
    effect = "Deny"
    principals {
      type        = "*"
      identifiers = ["*"]
    }
    resources = [aws_s3_bucket.cdn.arn, "${aws_s3_bucket.cdn.arn}/*"]
  }
}

resource "aws_s3_bucket_policy" "cdn" {
  bucket = aws_s3_bucket.cdn.id
  policy = data.aws_iam_policy_document.cdn_s3_policy.json
}

resource "aws_cloudfront_response_headers_policy" "cdn" {
  name    = "cdn-${terraform.workspace}"
  comment = ""

  security_headers_config {
    content_security_policy {
      content_security_policy = "default-src https://*.${var.cert-domain}; font-src 'none'; img-src 'self'; script-src 'self'; style-src 'none'; object-src 'none'; manifest-src 'self'"
      override                = true
    }

    content_type_options {
      override = true
    }

    frame_options {
      frame_option = "DENY"
      override     = true
    }

    referrer_policy {
      override        = true
      referrer_policy = "no-referrer"
    }

    strict_transport_security {
      access_control_max_age_sec = 63072000
      include_subdomains         = true
      preload                    = true
      override                   = true
    }

    xss_protection {
      override   = true
      protection = true
    }

  }
}

resource "aws_cloudfront_distribution" "cdn" {
  origin {
    domain_name = aws_s3_bucket.cdn.bucket_regional_domain_name
    origin_id   = "s3-cdn-${terraform.workspace}"

    s3_origin_config {
      origin_access_identity = aws_cloudfront_origin_access_identity.cdn.cloudfront_access_identity_path
    }
  }

  enabled             = true
  is_ipv6_enabled     = true
  default_root_object = "index.html"

  aliases = [var.media-domain]

  default_cache_behavior {
    allowed_methods        = ["GET", "HEAD"]
    cached_methods         = ["GET", "HEAD"]
    target_origin_id       = "s3-cdn-${terraform.workspace}"
    viewer_protocol_policy = "redirect-to-https"

    response_headers_policy_id = aws_cloudfront_response_headers_policy.cdn.id

    forwarded_values {
      query_string = false
      cookies {
        forward = "none"
      }
    }

    compress = true

    min_ttl     = "3600"
    default_ttl = "3600"
    max_ttl     = "86400"
  }

  price_class = "PriceClass_All"

  viewer_certificate {
    acm_certificate_arn      = var.acm-cert-arn-us
    minimum_protocol_version = "TLSv1.2_2021"
    ssl_support_method       = "sni-only"
  }

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

}

resource "aws_route53_record" "cdn-A" {
  zone_id = var.dns-zone-id
  name    = var.media-domain
  type    = "A"
  alias {
    evaluate_target_health = false
    name                   = aws_cloudfront_distribution.cdn.domain_name
    zone_id                = aws_cloudfront_distribution.cdn.hosted_zone_id
  }
}

resource "aws_route53_record" "cdn-AAAA" {
  zone_id = var.dns-zone-id
  name    = var.media-domain
  type    = "AAAA"
  alias {
    evaluate_target_health = false
    name                   = aws_cloudfront_distribution.cdn.domain_name
    zone_id                = aws_cloudfront_distribution.cdn.hosted_zone_id
  }
}
