resource "aws_ecr_repository" "jt_notes_lambda_get_all" {
  name = "jt-notes-lambda-get-all"

  image_scanning_configuration {
    scan_on_push = true
  }
}

# Now it is only create a new ECR, I still need to deploy manually. Steps to deploy is in README.md
