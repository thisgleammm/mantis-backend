package auth

import (
	"context"
	"errors"
	"log/slog"
	"net/mail"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
	repo "github.com/thisgleammm/mantis-backend/internal/adapters/postgresql/sqlc"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(ctx context.Context, req RegisterRequest) (repo.CreateUserRow, error)
	Login(ctx context.Context, req LoginRequest) (string, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

type RegisterRequest struct {
	Username    string `json:"username"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

func (r RegisterRequest) validate() error {
	if len(strings.TrimSpace(r.Username)) < 3 {
		return errors.New("username must be at least 3 characters")
	}
	if strings.TrimSpace(r.Name) == "" {
		return errors.New("name is required")
	}
	if _, err := mail.ParseAddress(r.Email); err != nil {
		return errors.New("valid email is required")
	}
	if len(r.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	return nil
}

func (s *svc) Register(ctx context.Context, req RegisterRequest) (repo.CreateUserRow, error) {
	if err := req.validate(); err != nil {
		slog.Warn("Register: validation failed", "error", err)
		return repo.CreateUserRow{}, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("Register: bcrypt failed", "error", err)
		return repo.CreateUserRow{}, err
	}

	user, err := s.repo.CreateUser(ctx, repo.CreateUserParams{
		Username:    req.Username,
		Name:        req.Name,
		Email:       req.Email,
		Password:    string(hashedPassword),
		PhoneNumber: pgtype.Text{String: req.PhoneNumber, Valid: req.PhoneNumber != ""},
	})
	if err != nil {
		slog.Error("Register: database error", "error", err, "email", req.Email)
		return repo.CreateUserRow{}, err
	}

	slog.Info("Register: user created", "user_id", user.ID, "username", user.Username)
	return user, nil
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r LoginRequest) validate() error {
	if strings.TrimSpace(r.Email) == "" {
		return errors.New("email is required")
	}
	if strings.TrimSpace(r.Password) == "" {
		return errors.New("password is required")
	}
	return nil
}

func (s *svc) Login(ctx context.Context, req LoginRequest) (string, error) {
	if err := req.validate(); err != nil {
		slog.Warn("Login: validation failed", "error", err)
		return "", err
	}

	user, err := s.repo.FindUserByEmailForLogin(ctx, req.Email)
	if err != nil {
		slog.Warn("Login: user not found", "email", req.Email)
		return "", errors.New("email tidak terdaftar")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		slog.Warn("Login: invalid password", "user_id", user.ID)
		return "", errors.New("invalid email or password")
	}

	token, err := GenerateToken(user.ID.String())
	if err != nil {
		slog.Error("Login: token generation failed", "error", err, "user_id", user.ID)
		return "", err
	}

	slog.Info("Login: user logged in", "user_id", user.ID)
	return token, nil
}
