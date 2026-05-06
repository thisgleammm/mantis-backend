package users

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	repo "github.com/thisgleammm/mantis-backend/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListUsers(ctx context.Context) ([]repo.ListUsersRow, error)
	FindUserByID(ctx context.Context, id string) (repo.FindUserByIDRow, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{
		repo: repo,
	}
}

func (s *svc) ListUsers(ctx context.Context) ([]repo.ListUsersRow, error) {
	return s.repo.ListUsers(ctx)
}

func (s *svc) FindUserByID(ctx context.Context, id string) (repo.FindUserByIDRow, error) {
	var uuid pgtype.UUID
	if err := uuid.Scan(id); err != nil {
		return repo.FindUserByIDRow{}, err
	}
	return s.repo.FindUserByID(ctx, uuid)
}
