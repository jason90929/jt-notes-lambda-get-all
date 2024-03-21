package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	getAll "github.com/jason90929/jt-notes-lambda-get-all"
	handler "github.com/jason90929/jt-notes-lambda-get-all/apigateway"
	"github.com/jason90929/jt-notes-lambda-get-all/dynamodb"
)

//go:embed example.json
var example []byte

func main() {
	repo := dynamodb.NewRepository()
	client := getAll.NewService(repo)
	h := handler.NewEvtHandler(client)

	var req events.APIGatewayProxyRequest
	err := json.Unmarshal(example, &req)
	if err != nil {
		panic(err)
	}
	resp, err := h.HandleEvent(context.Background(), req)
	if err != nil {
		panic(err)
	}
	fmt.Println("resp", resp)
}
