
resource "aws_backup_vault" "default" {
  name = "${var.client-code}-${terraform.workspace}"
}

resource "aws_backup_plan" "standard" {
  name = "${var.client-code}-${terraform.workspace}-daily"

  rule {
    rule_name         = "daily"
    target_vault_name = aws_backup_vault.default.name
    schedule          = "cron(0 12 * * ? *)"
    lifecycle {
      delete_after = 90
    }
  }

  rule {
    rule_name         = "weekly"
    target_vault_name = aws_backup_vault.default.name
    schedule          = "cron(0 12 ? * 1 *)"
    lifecycle {
      cold_storage_after = 14
      delete_after       = 365
    }
  }
}

resource "aws_iam_role" "backup" {
  name               = "backup-${var.client-code}-${terraform.workspace}"
  assume_role_policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": ["sts:AssumeRole"],
      "Effect": "allow",
      "Principal": {
        "Service": ["backup.amazonaws.com"]
      }
    }
  ]
}
POLICY
}

resource "aws_iam_role_policy_attachment" "backup" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSBackupServiceRolePolicyForBackup"
  role       = aws_iam_role.backup.name
}