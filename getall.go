package getAll

import "context"

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Repository interface {
	GetAll(ctx context.Context, req *Request) ([]Response, error)
}
