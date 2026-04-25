package products

import (
	"context"

	repo "github.com/thisgleammm/mantis-backend/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context) ([]repo.ListProductsRow, error)
	FindProductByID(ctx context.Context, id int64) (repo.FindProductByIDRow, error)
	FindProductBySlug(ctx context.Context, slug string) (repo.FindProductBySlugRow, error)
}

type svc struct {
	//repository
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{
		repo: repo,
	}
}

func (s *svc) ListProducts(ctx context.Context) ([]repo.ListProductsRow, error) {
	return s.repo.ListProducts(ctx)
}

func (s *svc) FindProductByID(ctx context.Context, id int64) (repo.FindProductByIDRow, error) {
	return s.repo.FindProductByID(ctx, id)
}

func (s *svc) FindProductBySlug(ctx context.Context, slug string) (repo.FindProductBySlugRow, error) {
	return s.repo.FindProductBySlug(ctx, slug)
}
