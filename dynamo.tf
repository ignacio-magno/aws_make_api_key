resource "aws_dynamodb_table" "dynamodb" {
  name         = var.DynamoDBTable
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "email"

  attribute {
    name = "email"
    type = "S"
  }
}