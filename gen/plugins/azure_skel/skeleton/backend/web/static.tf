
resource "aws_s3_bucket" "static" {
  bucket = "${var.client-code}-static-${terraform.workspace}"
}

resource "aws_s3_bucket_versioning" "static" {
  bucket = aws_s3_bucket.static.bucket

  versioning_configuration {
    status = "Enabled"
  }
}

resource "aws_s3_bucket_website_configuration" "static" {
  bucket = aws_s3_bucket.static.bucket

  index_document {
    suffix = "index.html"
  }

  error_document {
    key = "index.html"
  }
}

resource "aws_cloudfront_origin_access_identity" "static" {

}

data "aws_iam_policy_document" "static_s3_policy" {
  statement {
    actions   = ["s3:GetObject"]
    resources = ["${aws_s3_bucket.static.arn}/*"]

    principals {
      type        = "AWS"
      identifiers = [aws_cloudfront_origin_access_identity.static.iam_arn]
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
    resources = [aws_s3_bucket.static.arn, "${aws_s3_bucket.static.arn}/*"]
  }
}

resource "aws_s3_bucket_policy" "static" {
  bucket = aws_s3_bucket.static.id
  policy = data.aws_iam_policy_document.static_s3_policy.json
}

resource "aws_cloudfront_response_headers_policy" "static" {
  name    = "static-${terraform.workspace}"
  comment = ""

  security_headers_config {
    content_security_policy {
      content_security_policy = "default-src https://*.${var.cert-domain} ${var.graphql-api-url} https://cognito-idp.ap-southeast-2.amazonaws.com https://cognito-identity.ap-southeast-2.amazonaws.com; font-src data: https://use.fontawesome.com https://fonts.gstatic.com; img-src 'self' https://${var.media-domain}; script-src 'self' 'unsafe-inline' 'unsafe-eval' https://cdnjs.cloudflare.com; style-src 'self' 'unsafe-inline' https://use.fontawesome.com https://fonts.googleapis.com; object-src 'none'; manifest-src 'self'"
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

  custom_headers_config {
    items {
      header   = "permissions-policy"
      value    = "autoplay=(self)"
      override = true
    }
    items {
      header   = "feature-policy"
      value    = "autoplay self;"
      override = true
    }
  }
}

resource "aws_cloudfront_distribution" "static" {
  origin {
    domain_name = aws_s3_bucket.static.bucket_regional_domain_name
    origin_id   = "s3-static-${terraform.workspace}"

    s3_origin_config {
      origin_access_identity = aws_cloudfront_origin_access_identity.static.cloudfront_access_identity_path
    }
  }

  enabled             = true
  is_ipv6_enabled     = true
  default_root_object = "index.html"

  aliases = [var.web-domain]

  default_cache_behavior {
    allowed_methods        = ["GET", "HEAD"]
    cached_methods         = ["GET", "HEAD"]
    target_origin_id       = "s3-static-${terraform.workspace}"
    viewer_protocol_policy = "redirect-to-https"

    response_headers_policy_id = aws_cloudfront_response_headers_policy.static.id

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

  custom_error_response {
    error_code            = 403
    response_code         = 200
    response_page_path    = "/index.html"
    error_caching_min_ttl = 30
  }

  custom_error_response {
    error_code            = 404
    response_code         = 200
    response_page_path    = "/index.html"
    error_caching_min_ttl = 30
  }

}

resource "aws_route53_record" "static-A" {
  zone_id = var.dns-zone-id
  name    = var.web-domain
  type    = "A"
  alias {
    evaluate_target_health = false
    name                   = aws_cloudfront_distribution.static.domain_name
    zone_id                = aws_cloudfront_distribution.static.hosted_zone_id
  }
}

resource "aws_route53_record" "static-AAAA" {
  zone_id = var.dns-zone-id
  name    = var.web-domain
  type    = "AAAA"
  alias {
    evaluate_target_health = false
    name                   = aws_cloudfront_distribution.static.domain_name
    zone_id                = aws_cloudfront_distribution.static.hosted_zone_id
  }
}
