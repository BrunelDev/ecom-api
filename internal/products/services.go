package products

import (
	"context"

	repo "github.com/bruneldev/ecom-api/internal/adapter/postgresql"
)

type Service interface {
	GetProduct(ctx context.Context, id int64) (repo.Product, error)
	GetProductsList(ctx context.Context) ([]repo.Product, error)
}

type svc struct {
	repo repo.Querier
}

func (svc *svc) GetProductsList(ctx context.Context) ([]repo.Product, error) {
	return svc.repo.GetProductsList(ctx)
}

func (svc *svc) GetProduct(ctx context.Context, id int64) (repo.Product, error) {
	return svc.repo.GetProduct(ctx, id)
}

func NewService(repo repo.Querier) Service {
	return &svc{repo : repo }
}
