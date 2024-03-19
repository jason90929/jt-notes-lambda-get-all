package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	b, _ := json.Marshal(request)
	log.Println("raw request", string(b))
	params := request.QueryStringParameters
	name := params["name"]
	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       fmt.Sprintf("your query parameter name: %v", name),
	}
	return response, nil
}
