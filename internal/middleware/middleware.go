package middleware

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
		var tokenString string

		if authHeader != "" {
			parts := strings.Split(authHeader, " ")
			if len(parts) == 2 && parts[0] == "Bearer" {
				tokenString = parts[1]
			}
		}

		// Fallback to cookie if header is missing or invalid
		if tokenString == "" {
			cookie, err := r.Cookie("token")
			if err == nil {
				tokenString = cookie.Value
			}
		}

		if tokenString == "" {
			slog.Warn("auth middleware: missing token", "path", r.URL.Path)
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		userID, err := VerifyToken(tokenString)
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
