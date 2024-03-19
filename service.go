package getAll

import "context"

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (s *Client) GetAll(ctx context.Context, req *Request) (*Response, error) {
	// TODO: retrieve from DDB

	return &Response{
		Text: req.Name,
	}, nil
}
