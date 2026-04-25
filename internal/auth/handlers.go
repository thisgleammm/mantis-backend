package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	mantisJson "github.com/thisgleammm/mantis-backend/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{service: service}
}

type registerResponse struct {
	ID          int64  `json:"id"`
	Username    string `json:"username"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
}

type loginResponse struct {
	Token string `json:"token"`
}

// Register godoc
// @Summary Register a new user
// @Description Create a new user account
// @Tags auth
// @Accept  json
// @Produce  json
// @Param request body RegisterRequest true "Registration Info"
// @Success 201 {object} registerResponse
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
// @Success 200 {object} loginResponse
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
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	mantisJson.Write(w, http.StatusOK, map[string]string{"token": token})
}

// validationErrors are produced by request.validate() — safe to expose to client.
// DB errors and bcrypt errors are not in this set.
var validationMessages = map[string]struct{}{
	"username is required":                   {},
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
