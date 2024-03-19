variable "lambda_function_name" {
  default = "jt-notes-lambda-get-all"
}

resource "aws_iam_role" "assume_role" {
  name               = "jt-notes-role-get-all"
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": [
          "lambda.amazonaws.com"
        ]
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "ddb_readonly_policy_attachment" {
  role       = aws_iam_role.assume_role.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonDynamoDBReadOnlyAccess"
}

resource "aws_lambda_function" "jt_notes_lambda_get_all" {
  # If the file is not in the current working directory you will need to include a
  # path.module in the filename.
  image_uri     = "306698408315.dkr.ecr.ap-northeast-1.amazonaws.com/jt-notes-lambda-get-all:latest"
  function_name = var.lambda_function_name
  role          = aws_iam_role.assume_role.arn
  package_type  = "Image"

  # Advanced logging controls (optional)
  logging_config {
    log_format = "Text"
  }

  # ... other configuration ...
  depends_on = [
    aws_iam_role_policy_attachment.lambda_logs,
    aws_cloudwatch_log_group.example,
  ]
}

# See also the following AWS managed policy: AWSLambdaBasicExecutionRole
data "aws_iam_policy_document" "lambda_logging" {
  statement {
    effect = "Allow"

    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents",
    ]

    resources = ["arn:aws:logs:*:*:*"]
  }
}

# This is to optionally manage the CloudWatch Log Group for the Lambda Function.
# If skipping this resource configuration, also add "logs:CreateLogGroup" to the IAM policy below.
resource "aws_cloudwatch_log_group" "example" {
  name              = "/aws/lambda/${var.lambda_function_name}"
  retention_in_days = 14
}

resource "aws_iam_policy" "lambda_logging" {
  name        = "lambda_logging"
  path        = "/"
  description = "IAM policy for logging from a lambda"
  policy      = data.aws_iam_policy_document.lambda_logging.json
}

resource "aws_iam_role_policy_attachment" "lambda_logs" {
  role       = aws_iam_role.assume_role.name
  policy_arn = aws_iam_policy.lambda_logging.arn
}
