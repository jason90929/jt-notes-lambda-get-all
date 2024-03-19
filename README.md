login into AWS via aws-cli:

`aws ecr get-login-password --region ap-northeast-1 | docker login \
--username AWS --password-stdin 306698408315.dkr.ecr.ap-northeast-1.amazonaws.com`


upload image to ECR:

`docker build -t 306698408315.dkr.ecr.ap-northeast-1.amazonaws.com/jt-notes-lambda-get-all:latest .`
`docker push 306698408315.dkr.ecr.ap-northeast-1.amazonaws.com/jt-notes-lambda-get-all:latest`

Go in to the container locally:

`docker run --rm -it --entrypoint bash 306698408315.dkr.ecr.ap-northeast-1.amazonaws.com/jt-notes-lambda-get-all:latest`
