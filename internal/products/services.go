package products

import (
	"context"
)

type Service interface {
	ListProducts(ctx context.Context) error
}

type svc struct {
}

func (svc *svc) ListProducts(ctx context.Context) error {
	return nil
}

func NewService() Service {
	return &svc{}
}
