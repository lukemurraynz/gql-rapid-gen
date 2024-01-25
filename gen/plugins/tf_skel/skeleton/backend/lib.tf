data "archive_file" "lambda-lib-source" {
  type        = "zip"
  output_path = "${path.module}/build/src/lib.zip"
  source_dir  = "${path.module}/../lib/"
}