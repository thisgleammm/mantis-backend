package domain

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"time"
)

type PasswordReset struct {
	Email     string
	Token     string
	ExpiresAt time.Time
}

type AuthRepository interface {
	UpsertPasswordReset(ctx context.Context, pr PasswordReset) error
	FindPasswordResetByToken(ctx context.Context, token string) (PasswordReset, error)
	DeletePasswordReset(ctx context.Context, email string) error
	UpdatePassword(ctx context.Context, email, hashedPassword string) error
}

func GenerateRandomToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
