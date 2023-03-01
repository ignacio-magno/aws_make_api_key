# provider aws
provider "aws" {
  region = "us-west-2"
}

data "aws_caller_identity" "current" {}

# resource iam role with policy to invoke lambda
# form many methods, use the same role
resource "aws_iam_role" "role" {
  name               = var.ROLE_NAME
  // assume role policy lambda
  assume_role_policy = jsonencode({
    Version   = "2012-10-17"
    Statement = [
      {
        Action    = "sts:AssumeRole",
        Effect    = "Allow"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      },
    ]
  })
}

// policy to logs
resource "aws_iam_role_policy" "logs" {
  name   = "logs"
  role   = aws_iam_role.role.id
  policy = jsonencode({
    Version   = "2012-10-17"
    Statement = [
      {
        Action = [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents",
        ]
        Effect   = "Allow"
        Resource = "arn:aws:logs:*:*:*"
      },
    ]
  })
}

# deploye lambda function
resource "aws_lambda_function" "lambda_function" {
  function_name    = var.LAMBDA_FUNCTION_NAME
  handler          = var.HANDLER
  filename         = data.archive_file.lambda_zip.output_path
  role             = aws_iam_role.role.arn
  runtime          = "go1.x"
  source_code_hash = data.archive_file.lambda_zip.output_base64sha256
  depends_on       = [data.archive_file.lambda_zip]

  timeouts {
    create = "1m"
  }

  environment {
    variables = {
      "DynamoDBTable"           = var.DynamoDBTable
      "USAGE_PLAN_ID"           = var.USAGE_PLAN_ID
      "EmailToNotifyErrorsApi"  = var.EmailToNotifyErrorsApi
      "EMAIL_SOURCE"            = var.EMAIL_SOURCE
      "EMAIL_SUBJECT"           = var.EMAIL_SUBJECT
      "EMAIL_TEST"              = var.EMAIL_TEST
      "TEMPLATE_EMAIL_SES_NAME" = var.TEMPLATE_EMAIL_SES_NAME
    }
  }
}

# add permission to read and write to dynamodb
resource "aws_iam_role_policy" "policy_lambda_dynamodb" {
  name = format("%s-policy-dynamodb", var.LAMBDA_FUNCTION_NAME)
  role = aws_iam_role.role.id

  policy = jsonencode({
    Version   = "2012-10-17"
    Statement = [
      {
        Action = [
          "dynamodb:PutItem",
          "dynamodb:GetItem",
          "dynamodb:Scan",
          "dynamodb:Query"
        ]
        Effect   = "Allow"
        Resource = [
          "arn:aws:dynamodb:us-west-2:${data.aws_caller_identity.current.account_id}:table/${var.DynamoDBTable}",
          "arn:aws:dynamodb:us-west-2:${data.aws_caller_identity.current.account_id}:table/${var.DynamoDBTable}/*",
        ]
      }
    ]
  })
}

# add permission to send email by ses
resource "aws_iam_role_policy" "policy_lambda_ses" {
  name = format("%s-policy-ses", var.LAMBDA_FUNCTION_NAME)
  role = aws_iam_role.role.id

  policy = jsonencode({
    Version   = "2012-10-17"
    Statement = [
      {
        Effect = "Allow",
        Action = [
          "ses:SendEmail",
          "ses:SendRawEmail",
          "ses:SendTemplatedEmail"
        ]
        Resource = "*"
      }
    ]
  })
}


# add permission to make api key usage plan
resource "aws_iam_role_policy" "policy_lambda_api_gateway" {
  name = format("%s-policy-api-gateway", var.LAMBDA_FUNCTION_NAME)
  role = aws_iam_role.role.id

  policy = jsonencode({
    Version   = "2012-10-17"
    Statement = [
      {
        Action = [
          "apigateway:GET",
          "apigateway:POST",
          "apigateway:PUT",
          "apigateway:PATCH",
          "apigateway:DELETE",
        ]
        Effect   = "Allow"
        Resource = "*"
      }
    ]
  })
}