variable "DynamoDBTable" {
  type = string
}

// set values
variable "LAMBDA_FUNCTION_NAME" {
  type = string
}

variable "ROLE_NAME" {
  type = string
}

variable "HANDLER" {
  type    = string
  default = "main"
}

variable "USAGE_PLAN_ID" {
  type = string
}

variable "EmailToNotifyErrorsApi" {
  type = string
}

variable "EMAIL_SOURCE" {
  type = string
}
variable "EMAIL_SUBJECT" {
  type = string
}
variable "EMAIL_TEST" {
  type = string
}

variable "TEMPLATE_EMAIL_SES_NAME" {
  type = string
}