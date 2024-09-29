resource "aws_ses_configuration_set" "emails" {
  name = "${var.client-code}-${terraform.workspace}"

  reputation_metrics_enabled = true

  delivery_options {
    tls_policy = "Require"
  }
}