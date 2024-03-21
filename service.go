package getAll

import (
	"context"
)

type Service struct {
	r Repository
}

func NewService(r Repository) *Service {
	return &Service{
		r: r,
	}
}

func (s *Service) GetAll(ctx context.Context, req *Request) ([]Response, error) {
	return s.r.GetAll(ctx, req)
}
