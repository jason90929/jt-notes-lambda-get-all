resource "aws_iam_role" "jt_notes_get_all_role" {
  name = "jt_notes_get_all_role"
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
    role = "${aws_iam_role.jt_notes_get_all_role.name}"
    policy_arn = "arn:aws:iam::aws:policy/AmazonDynamoDBReadOnlyAccess"
}

resource "aws_lambda_function" "jt_notes_lambda_get_all" {
  # If the file is not in the current working directory you will need to include a
  # path.module in the filename.
  image_uri     = "306698408315.dkr.ecr.ap-northeast-1.amazonaws.com/jt-notes-lambda-get-all:latest"
  function_name = "jt-notes-lambda-get-all"
  role          = aws_iam_role.jt_notes_get_all_role.arn
  package_type  = "Image"
}