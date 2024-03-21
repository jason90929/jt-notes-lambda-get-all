package handler

import (
	"context"
	"encoding/json"
	"log"
	"log/slog"

	"github.com/aws/aws-lambda-go/events"
	getAll "github.com/jason90929/jt-notes-lambda-get-all"
)

type server struct {
	c *getAll.Service
}

func NewEvtHandler(c *getAll.Service) *server {
	return &server{
		c: c,
	}
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
