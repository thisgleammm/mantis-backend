package auth

import (
	"context"
	"log/slog"
	"net/http"
	"strings"
)

type contextKey string

const UserIDKey contextKey = "user_id"

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Phase 1 tracing: log entry point for each protected request
		slog.Debug("auth middleware", "method", r.Method, "path", r.URL.Path)

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			slog.Warn("auth middleware: missing Authorization header", "path", r.URL.Path)
			http.Error(w, "missing authorization header", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			slog.Warn("auth middleware: malformed Authorization header", "path", r.URL.Path)
			http.Error(w, "invalid authorization header", http.StatusUnauthorized)
			return
		}

		userID, err := VerifyToken(parts[1])
		if err != nil {
			// Log root cause — expired vs invalid vs tampered
			slog.Warn("auth middleware: token verification failed", "path", r.URL.Path, "error", err)
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		slog.Debug("auth middleware: token valid", "user_id", userID, "path", r.URL.Path)
		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
