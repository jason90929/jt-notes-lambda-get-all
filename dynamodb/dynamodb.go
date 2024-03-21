package dynamodb

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	getAll "github.com/jason90929/jt-notes-lambda-get-all"
)

type Repository struct {
	c         *dynamodb.Client
	tableName string
}

// NewMysqlRepository ...
func NewRepository() *Repository {
	// Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-northeast-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Using the Config value, create the DynamoDB client
	client := dynamodb.NewFromConfig(cfg)

	return &Repository{
		c:         client,
		tableName: "jt-notes",
	}
}

func (r *Repository) GetAll(ctx context.Context, req *getAll.Request) ([]getAll.Response, error) {
	var err error
	condition := expression.Name("name").Contains(req.Name)
	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		log.Fatalf("Got error building expression: %s", err)
		return nil, fmt.Errorf("failed to query db")
	}
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Condition(),
		TableName:                 aws.String(r.tableName),
	}
	resp, err := r.c.Scan(ctx, params)
	if err != nil {
		log.Fatalf("Got error scan ddb: %s", err)
		return nil, fmt.Errorf("failed to query db")
	}

	records := make([]Record, resp.Count)
	response := make([]getAll.Response, resp.Count)
	for i, item := range resp.Items {
		attributevalue.UnmarshalMap(item, &records[i])
		response[i] = toCaninocal(&records[i])
	}

	return response, nil
}

func toCaninocal(rec *Record) getAll.Response {
	return getAll.Response{
		Name:        rec.Name,
		Title:       rec.Title,
		Description: rec.Description,
	}
}
