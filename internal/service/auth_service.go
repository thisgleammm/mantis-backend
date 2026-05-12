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
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, err
	}
	u.Password = string(hashed)
	return s.userRepo.Create(ctx, u)
}

func (s *AuthService) Login(ctx context.Context, email, password string) (domain.User, string, string, error) {
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, "", "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return domain.User{}, "", "", errors.New("invalid credentials")
	}

	now := time.Now()
	accessClaims := jwt.MapClaims{
		"user_id": user.ID,
		"type":    "access",
		"iss":     "mantis-api",
		"aud":     "mantis-client",
		"nbf":     now.Unix(),
		"exp":     now.Add(15 * time.Minute).Unix(),
	}
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString([]byte(s.secret))
	if err != nil {
		return domain.User{}, "", "", err
	}

	refreshClaims := jwt.MapClaims{
		"user_id": user.ID,
		"type":    "refresh",
		"iss":     "mantis-api",
		"aud":     "mantis-client",
		"nbf":     now.Unix(),
		"exp":     now.Add(7 * 24 * time.Hour).Unix(),
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(s.secret))

	return user, accessToken, refreshToken, err
}

func (s *AuthService) RefreshToken(ctx context.Context, refreshTokenString string) (string, string, error) {
	token, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.secret), nil
	})
	if err != nil {
		return "", "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims["type"] != "refresh" {
		return "", "", errors.New("invalid refresh token")
	}

	userID := claims["user_id"].(string)
	now := time.Now()

	accessClaims := jwt.MapClaims{
		"user_id": userID,
		"type":    "access",
		"iss":     "mantis-api",
		"aud":     "mantis-client",
		"nbf":     now.Unix(),
		"exp":     now.Add(15 * time.Minute).Unix(),
	}
	newAccessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString([]byte(s.secret))
	if err != nil {
		return "", "", err
	}

	refreshClaims := jwt.MapClaims{
		"user_id": userID,
		"type":    "refresh",
		"iss":     "mantis-api",
		"aud":     "mantis-client",
		"nbf":     now.Unix(),
		"exp":     now.Add(7 * 24 * time.Hour).Unix(),
	}
	newRefreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(s.secret))

	return newAccessToken, newRefreshToken, err
}
