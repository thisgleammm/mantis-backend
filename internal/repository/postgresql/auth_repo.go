package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	repo "github.com/thisgleammm/mantis-backend/internal/adapters/postgresql/sqlc"
	"github.com/thisgleammm/mantis-backend/internal/domain"
)

type AuthRepository struct {
	q *repo.Queries
}

func NewAuthRepository(q *repo.Queries) *AuthRepository {
	return &AuthRepository{q: q}
}

func (r *AuthRepository) UpsertPasswordReset(ctx context.Context, pr domain.PasswordReset) error {
	return r.q.UpsertPasswordReset(ctx, repo.UpsertPasswordResetParams{
		Email:     pr.Email,
		Token:     pr.Token,
		ExpiresAt: pgtype.Timestamptz{Time: pr.ExpiresAt, Valid: true},
	})
}

func (r *AuthRepository) FindPasswordResetByToken(ctx context.Context, token string) (domain.PasswordReset, error) {
	row, err := r.q.FindPasswordResetByToken(ctx, token)
	if err != nil {
		return domain.PasswordReset{}, err
	}

	return domain.PasswordReset{
		Email:     row.Email,
		Token:     row.Token,
		ExpiresAt: row.ExpiresAt.Time,
	}, nil
}

func (r *AuthRepository) DeletePasswordReset(ctx context.Context, email string) error {
	return r.q.DeletePasswordReset(ctx, email)
}

func (r *AuthRepository) UpdatePassword(ctx context.Context, email, hashedPassword string) error {
	return r.q.UpdateUserPassword(ctx, repo.UpdateUserPasswordParams{
		Email:    email,
		Password: hashedPassword,
	})
}
