package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	getAll "github.com/jason90929/jt-notes-lambda-get-all"
	handler "github.com/jason90929/jt-notes-lambda-get-all/apigateway"
	"github.com/jason90929/jt-notes-lambda-get-all/dynamodb"
)

func main() {
	repo := dynamodb.NewRepository()
	client := getAll.NewService(repo)
	h := handler.NewEvtHandler(client)

	lambda.Start(h.HandleEvent)
}
