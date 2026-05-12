package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/resend/resend-go/v3"
	"github.com/thisgleammm/mantis-backend/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo domain.UserRepository
	authRepo domain.AuthRepository
	resend   *resend.Client
	secret   string
}

func NewAuthService(userRepo domain.UserRepository, authRepo domain.AuthRepository, resend *resend.Client, secret string) *AuthService {
	return &AuthService{userRepo: userRepo, authRepo: authRepo, resend: resend, secret: secret}
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

func (s *AuthService) ForgotPassword(ctx context.Context, email string) error {
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		// We return nil to avoid email enumeration.
		return nil
	}

	// Generate a random token
	token := domain.GenerateRandomToken(32)
	expiresAt := time.Now().Add(1 * time.Hour)

	err = s.authRepo.UpsertPasswordReset(ctx, domain.PasswordReset{
		Email:     user.Email,
		Token:     token,
		ExpiresAt: expiresAt,
	})
	if err != nil {
		return err
	}

	// Send email via Resend
	if s.resend != nil {
		resetLink := fmt.Sprintf("https://mantis-marketplace.vercel.app/reset-password?token=%s", token)
		params := &resend.SendEmailRequest{
			From:    "Mantis <onboarding@resend.dev>",
			To:      []string{user.Email},
			Subject: "Reset Your Password - Mantis",
			Html:    fmt.Sprintf("<p>Hello %s,</p><p>We received a request to reset your password. Click the link below to proceed:</p><p><a href='%s'>Reset Password</a></p><p>This link will expire in 1 hour.</p>", user.Name, resetLink),
		}

		_, err := s.resend.Emails.Send(params)
		if err != nil {
			slog.Error("Failed to send reset email", "error", err, "email", user.Email)
			return errors.New("failed to send reset email")
		}
	}

	slog.Info("Password reset token generated and sent", "email", email, "token", token)

	return nil
}

func (s *AuthService) ResetPassword(ctx context.Context, token, newPassword string) error {
	pr, err := s.authRepo.FindPasswordResetByToken(ctx, token)
	if err != nil {
		return errors.New("invalid or expired token")
	}

	if pr.ExpiresAt.Before(time.Now()) {
		s.authRepo.DeletePasswordReset(ctx, pr.Email)
		return errors.New("token expired")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = s.authRepo.UpdatePassword(ctx, pr.Email, string(hashed))
	if err != nil {
		return err
	}

	return s.authRepo.DeletePasswordReset(ctx, pr.Email)
}

func (s *AuthService) ConfirmPassword(ctx context.Context, userID string, password string) error {
	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		return errors.New("user not found")
	}

	userWithPass, err := s.userRepo.FindByEmail(ctx, user.Email)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userWithPass.Password), []byte(password)); err != nil {
		return errors.New("invalid password")
	}

	return nil
}
