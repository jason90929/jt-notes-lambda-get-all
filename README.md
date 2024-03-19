# Section 1: Update content to ECR

Login into AWS via aws-cli:

`aws ecr get-login-password --region ap-northeast-1 | docker login \
--username AWS --password-stdin 306698408315.dkr.ecr.ap-northeast-1.amazonaws.com`


Upload image to ECR:

`docker build -t 306698408315.dkr.ecr.ap-northeast-1.amazonaws.com/jt-notes-lambda-get-all:latest .`
`docker push 306698408315.dkr.ecr.ap-northeast-1.amazonaws.com/jt-notes-lambda-get-all:latest`

Trigger Lambda to retrieve latest image:

`aws lambda update-function-code \
           --function-name jt-notes-lambda-get-all \
           --image-uri 306698408315.dkr.ecr.ap-northeast-1.amazonaws.com/jt-notes-lambda-get-all:latest`

Go in to the container locally:

`docker run --rm -it --entrypoint bash 306698408315.dkr.ecr.ap-northeast-1.amazonaws.com/jt-notes-lambda-get-all:latest`


# Section 2: Deploy new terraform setup to AWS

Usually this just need to do one time until some terraform changes

`cd ./terraform`
`terraform plan`
`terraform apply`

If the cost is too high, command to destroy correlated services

`terraform destroy`
