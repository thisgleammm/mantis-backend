package carts

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	repo "github.com/thisgleammm/mantis-backend/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListCarts(ctx context.Context, userID string) ([]repo.Cart, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) ListCarts(ctx context.Context, userID string) ([]repo.Cart, error) {
	var userUUID pgtype.UUID
	if err := userUUID.Scan(userID); err != nil {
		return nil, err
	}
	return s.repo.ListCarts(ctx, userUUID)
}
