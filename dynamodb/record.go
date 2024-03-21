package dynamodb

type Record struct {
	ID          string `dynamodbav:"id`
	Owner       string `dynamodbav:"owner"`
	CreatedDate string `dynamodbav:"createdDate"`
	Name        string `dynamodbav:"name"`
	Title       string `dynamodbav:"title"`
	Description string `dynamodbav:"description"`
}
