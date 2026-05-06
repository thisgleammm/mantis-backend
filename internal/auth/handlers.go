package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"

	mantisJson "github.com/thisgleammm/mantis-backend/internal/json"
)

// isSecureCookie returns true in production (HTTPS required).
// In development over plain HTTP, Secure=true causes browsers to silently
// drop the cookie — so we disable it when APP_ENV != "production".
func isSecureCookie() bool {
	return os.Getenv("APP_ENV") == "production"
}

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{service: service}
}

type userResponse struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
}

// Register godoc
// @Summary Register a new user
// @Description Create a new user account
// @Tags auth
// @Accept  json
// @Produce  json
// @Param request body RegisterRequest true "Registration Info"
// @Success 201 {object} userResponse
// @Failure 400 {string} string "invalid request body or validation error"
// @Router /auth/register [post]
func (h *handler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.service.Register(r.Context(), req)
	if err != nil {
		if isValidationError(err) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	mantisJson.Write(w, http.StatusCreated, user)
}

// Login godoc
// @Summary Login user
// @Description Authenticate user and return a token
// @Tags auth
// @Accept  json
// @Produce  json
// @Param request body LoginRequest true "Login Credentials"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "invalid request body or validation error"
// @Failure 401 {string} string "unauthorized"
// @Router /auth/login [post]
func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	token, err := h.service.Login(r.Context(), req)
	if err != nil {
		if isValidationError(err) {
			mantisJson.Write(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
			return
		}
		mantisJson.Write(w, http.StatusUnauthorized, map[string]string{"message": err.Error()})
		return
	}

	secure := isSecureCookie()
	sameSite := http.SameSiteLaxMode
	if secure {
		sameSite = http.SameSiteNoneMode // SameSite=None requires Secure=true
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(72 * time.Hour),
		HttpOnly: true,
		Secure:   secure,
		Path:     "/",
		SameSite: sameSite,
	})

	mantisJson.Write(w, http.StatusOK, map[string]string{
		"message": "logged in successfully",
		"token":   token,
	})
}

// Logout godoc
// @Summary Logout user
// @Description Logout the current user (client-side should discard the token)
// @Tags auth
// @Produce  json
// @Success 200 {string} string "successfully logged out"
// @Router /auth/logout [post]
func (h *handler) Logout(w http.ResponseWriter, r *http.Request) {
	secure := isSecureCookie()
	sameSite := http.SameSiteLaxMode
	if secure {
		sameSite = http.SameSiteNoneMode
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   secure,
		Path:     "/",
		SameSite: sameSite,
	})
	mantisJson.Write(w, http.StatusOK, map[string]string{"message": "successfully logged out"})
}

// validationErrors are produced by request.validate() — safe to expose to client.
// DB errors and bcrypt errors are not in this set.
var validationMessages = map[string]struct{}{
	"username is required":                   {},
	"username must be at least 3 characters": {},
	"name is required":                       {},
	"valid email is required":                {},
	"password must be at least 8 characters": {},
	"email is required":                      {},
	"password is required":                   {},
}

func isValidationError(err error) bool {
	if err == nil {
		return false
	}
	_, ok := validationMessages[err.Error()]
	return ok
}

// sentinel so callers can type-check if needed
var ErrValidation = errors.New("validation error")
