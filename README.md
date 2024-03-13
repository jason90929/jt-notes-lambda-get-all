
upload image to ECR

`docker build -t 306698408315.dkr.ecr.ap-northeast-1.amazonaws.com/login-info-lambda-query:latest .`
`docker push 306698408315.dkr.ecr.ap-northeast-1.amazonaws.com/login-info-lambda-query:latest`

Go in to the container locally

`docker run --rm -it --entrypoint bash 306698408315.dkr.ecr.ap-northeast-1.amazonaws.com/login-info-lambda-query:latest`
