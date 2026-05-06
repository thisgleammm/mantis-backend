package postgresql

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	repo "github.com/thisgleammm/mantis-backend/internal/adapters/postgresql/sqlc"
	"github.com/thisgleammm/mantis-backend/internal/domain"
)

type UserRepository struct {
	q *repo.Queries
}

func NewUserRepository(q *repo.Queries) *UserRepository {
	return &UserRepository{q: q}
}

func (r *UserRepository) Create(ctx context.Context, u domain.User) (domain.User, error) {
	row, err := r.q.CreateUser(ctx, repo.CreateUserParams{
		Username:    u.Username,
		Name:        u.Name,
		Email:       u.Email,
		Password:    u.Password,
		PhoneNumber: pgtype.Text{String: u.PhoneNumber, Valid: u.PhoneNumber != ""},
	})
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:          row.ID.String(),
		Username:    row.Username,
		Name:        row.Name,
		Email:       row.Email,
		PhoneNumber: row.PhoneNumber.String,
		CreatedAt:   row.CreatedAt.Time,
	}, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (domain.User, error) {
	var uuid pgtype.UUID
	if err := uuid.Scan(id); err != nil {
		return domain.User{}, err
	}

	row, err := r.q.FindUserByID(ctx, uuid)
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:          row.ID.String(),
		Name:        row.Name,
		Email:       row.Email,
		PhoneNumber: row.PhoneNumber.String,
		CreatedAt:   row.CreatedAt.Time,
	}, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	row, err := r.q.FindUserByEmailForLogin(ctx, email)
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:       row.ID.String(),
		Name:     row.Name,
		Email:    row.Email,
		Password: row.Password,
	}, nil
}

func (r *UserRepository) List(ctx context.Context) ([]domain.User, error) {
	rows, err := r.q.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	var users []domain.User
	for _, row := range rows {
		users = append(users, domain.User{
			ID:          row.ID.String(),
			Name:        row.Name,
			Email:       row.Email,
			PhoneNumber: row.PhoneNumber.String,
			CreatedAt:   row.CreatedAt.Time,
		})
	}
	return users, nil
}
