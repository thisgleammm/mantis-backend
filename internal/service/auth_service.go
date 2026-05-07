package service

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/thisgleammm/mantis-backend/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo domain.UserRepository
	secret   string
}

func NewAuthService(userRepo domain.UserRepository, secret string) *AuthService {
	return &AuthService{userRepo: userRepo, secret: secret}
}

func (s *AuthService) Register(ctx context.Context, u domain.User) (domain.User, error) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hashed)
	return s.userRepo.Create(ctx, u)
}

func (s *AuthService) Login(ctx context.Context, email, password string) (domain.User, string, error) {
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return domain.User{}, "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.secret))
	return user, tokenString, err
}
