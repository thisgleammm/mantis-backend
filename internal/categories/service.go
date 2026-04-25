package categories

import (
	"context"

	repo "github.com/thisgleammm/mantis-backend/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListCategories(ctx context.Context) ([]repo.Category, error)
	FindCategoryByID(ctx context.Context, id int64) (repo.Category, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{
		repo: repo,
	}
}

func (s *svc) ListCategories(ctx context.Context) ([]repo.Category, error) {
	return s.repo.ListCategories(ctx)
}

func (s *svc) FindCategoryByID(ctx context.Context, id int64) (repo.Category, error) {
	return s.repo.FindCategoryByID(ctx, id)
}