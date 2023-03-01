data "archive_file" "lambda_zip" {
  type        = "zip"
  source_file = "${path.module}/.build/main"
  output_path = "${path.module}/.build/main.zip"
}