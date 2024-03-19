package handler

import (
	"context"
	"encoding/json"
	"log"
	"log/slog"

	"github.com/aws/aws-lambda-go/events"
	getAll "github.com/jason90929/lambda-ddb-query"
)

type server struct {
	c *getAll.Client
}

func NewServer(c *getAll.Client) *server {
	return &server{
		c: c,
	}
}

func Must(s *server) (*server, error) {
	if s == nil {
		panic("server must be implemented")
	}
	return s, nil
}

func (s *server) HandleEvent(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	params := request.QueryStringParameters
	name := params["name"]

	req := getAll.Request{
		Name: name,
	}

	resp, err := s.c.GetAll(ctx, &req)
	if err != nil {
		log.Println("failed to retrieve ddb", slog.Any("error", err))
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "failed query ddb",
		}, nil
	}

	b, err := json.Marshal(resp)
	if err != nil {
		log.Println("failed to marshal record", slog.Any("error", err))
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "failed query ddb",
		}, nil
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(b),
	}
	return response, nil
}
