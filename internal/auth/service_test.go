package auth

import (
	"context"
	"errors"
	"testing"

	repo "github.com/thisgleammm/mantis-backend/internal/adapters/postgresql/sqlc"
	"golang.org/x/crypto/bcrypt"
)

// --- Register validation tests ---

func TestRegisterRequest_validate(t *testing.T) {
	tests := []struct {
		name    string
		req     RegisterRequest
		wantErr string
	}{
		{
			name:    "valid request",
			req:     RegisterRequest{Username: "jdoe", Name: "John Doe", Email: "john@example.com", Password: "secret123"},
			wantErr: "",
		},
		{
			name:    "missing username",
			req:     RegisterRequest{Name: "John", Email: "john@example.com", Password: "secret123"},
			wantErr: "username is required",
		},
		{
			name:    "missing name",
			req:     RegisterRequest{Username: "jdoe", Email: "john@example.com", Password: "secret123"},
			wantErr: "name is required",
		},
		{
			name:    "invalid email",
			req:     RegisterRequest{Username: "jdoe", Name: "John", Email: "not-an-email", Password: "secret123"},
			wantErr: "valid email is required",
		},
		{
			name:    "password too short",
			req:     RegisterRequest{Username: "jdoe", Name: "John", Email: "john@example.com", Password: "short"},
			wantErr: "password must be at least 8 characters",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.req.validate()
			if tc.wantErr == "" {
				if err != nil {
					t.Errorf("expected no error, got: %v", err)
				}
				return
			}
			if err == nil {
				t.Errorf("expected error %q, got nil", tc.wantErr)
				return
			}
			if err.Error() != tc.wantErr {
				t.Errorf("expected %q, got %q", tc.wantErr, err.Error())
			}
		})
	}
}

// --- Login validation tests ---

func TestLoginRequest_validate(t *testing.T) {
	tests := []struct {
		name    string
		req     LoginRequest
		wantErr string
	}{
		{
			name:    "valid request",
			req:     LoginRequest{Email: "john@example.com", Password: "secret123"},
			wantErr: "",
		},
		{
			name:    "missing email",
			req:     LoginRequest{Password: "secret123"},
			wantErr: "email is required",
		},
		{
			name:    "missing password",
			req:     LoginRequest{Email: "john@example.com"},
			wantErr: "password is required",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.req.validate()
			if tc.wantErr == "" {
				if err != nil {
					t.Errorf("expected no error, got: %v", err)
				}
				return
			}
			if err == nil {
				t.Errorf("expected error %q, got nil", tc.wantErr)
				return
			}
			if err.Error() != tc.wantErr {
				t.Errorf("expected %q, got %q", tc.wantErr, err.Error())
			}
		})
	}
}

// --- Service Tests ---

func TestService_Register(t *testing.T) {
	ctx := context.Background()
	stub := &stubQuerier{
		createUserFn: func(ctx context.Context, arg repo.CreateUserParams) (repo.CreateUserRow, error) {
			return repo.CreateUserRow{ID: 1, Username: arg.Username}, nil
		},
	}
	s := NewService(stub)

	req := RegisterRequest{Username: "jdoe", Name: "John", Email: "john@example.com", Password: "password123"}
	user, err := s.Register(ctx, req)
	if err != nil {
		t.Fatalf("Register failed: %v", err)
	}
	if user.ID != 1 || user.Username != "jdoe" {
		t.Errorf("unexpected user: %+v", user)
	}
}

func TestService_Login(t *testing.T) {
	ctx := context.Background()
	password := "password123"
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	stub := &stubQuerier{
		findUserByEmailForLoginFn: func(ctx context.Context, email string) (repo.FindUserByEmailForLoginRow, error) {
			if email == "john@example.com" {
				return repo.FindUserByEmailForLoginRow{ID: 1, Password: string(hash)}, nil
			}
			return repo.FindUserByEmailForLoginRow{}, errors.New("not found")
		},
	}
	s := NewService(stub)

	t.Run("success", func(t *testing.T) {
		token, err := s.Login(ctx, LoginRequest{Email: "john@example.com", Password: password})
		if err != nil {
			t.Fatalf("Login failed: %v", err)
		}
		if token == "" {
			t.Error("expected token, got empty")
		}
	})

	t.Run("invalid password", func(t *testing.T) {
		_, err := s.Login(ctx, LoginRequest{Email: "john@example.com", Password: "wrong"})
		if err == nil || err.Error() != "invalid email or password" {
			t.Errorf("expected invalid password error, got: %v", err)
		}
	})
}

// --- Stub ---

type stubQuerier struct {
	repo.Querier
	createUserFn              func(ctx context.Context, arg repo.CreateUserParams) (repo.CreateUserRow, error)
	findUserByEmailForLoginFn func(ctx context.Context, email string) (repo.FindUserByEmailForLoginRow, error)
}

func (s *stubQuerier) CreateUser(ctx context.Context, arg repo.CreateUserParams) (repo.CreateUserRow, error) {
	return s.createUserFn(ctx, arg)
}

func (s *stubQuerier) FindUserByEmailForLogin(ctx context.Context, email string) (repo.FindUserByEmailForLoginRow, error) {
	return s.findUserByEmailForLoginFn(ctx, email)
}
