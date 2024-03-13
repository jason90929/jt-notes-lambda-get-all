resource "aws_iam_role" "login_info_query_role" {
  name = "login_info_query_role"
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
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

resource "aws_iam_role_policy_attachment" "ddb-read-only-policy-attachment" {
    role = "${aws_iam_role.login_info_query_role.name}"
    policy_arn = "arn:aws:iam::aws:policy/AmazonDynamoDBReadOnlyAccess"
}

resource "aws_lambda_function" "login_info_query_lambda" {
  # If the file is not in the current working directory you will need to include a
  # path.module in the filename.
  image_uri     = "306698408315.dkr.ecr.ap-northeast-1.amazonaws.com/login-info-lambda-query:latest"
  function_name = "login_info_query_lambda"
  role          = aws_iam_role.login_info_query_role.arn
  package_type  = "Image"
}