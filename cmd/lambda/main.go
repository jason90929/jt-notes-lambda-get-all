package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	getAll "github.com/jason90929/lambda-ddb-query"
	handler "github.com/jason90929/lambda-ddb-query/apigateway"
)

func main() {
	client := getAll.NewClient()
	server, err := handler.Must(handler.NewServer(client))
	if err != nil {
		panic("init server failed")
	}

	lambda.Start(server.HandleEvent)
}
