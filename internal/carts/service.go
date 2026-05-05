package carts

import (
	"context"

	repo "github.com/thisgleammm/mantis-backend/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListCarts(ctx context.Context) ([]repo.Cart, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) ListCarts(ctx context.Context) ([]repo.Cart, error) {
	return s.repo.ListCarts(ctx)
}
