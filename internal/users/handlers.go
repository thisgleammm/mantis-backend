package users

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/thisgleammm/mantis-backend/internal/auth"
	"github.com/thisgleammm/mantis-backend/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

type userResponse struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
}

// ListUsers godoc
// @Summary List all users
// @Description Get a list of all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} userResponse
// @Router /users [get]
func (h *handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.ListUsers(r.Context())
	if err != nil {
		slog.Error("ListUsers failed", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, users)
}

// FindUserByID godoc
// @Summary Find user by ID
// @Description Get a user by its ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} userResponse
// @Failure 404 {string} string "user not found"
// @Router /users/{id} [get]
func (h *handler) FindUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		slog.Warn("FindUserByID: empty id")
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	user, err := h.service.FindUserByID(r.Context(), id)
	if err != nil {
		slog.Warn("FindUserByID: not found", "id", id, "error", err)
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	json.Write(w, http.StatusOK, user)
}

// GetMe godoc
// @Summary Get current user
// @Description Get the currently authenticated user
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} userResponse
// @Failure 401 {string} string "unauthorized"
// @Failure 404 {string} string "user not found"
// @Security ApiKeyAuth
// @Router /users/me [get]
func (h *handler) GetMe(w http.ResponseWriter, r *http.Request) {
	userIDVal := r.Context().Value(auth.UserIDKey)
	if userIDVal == nil {
		slog.Warn("GetMe: missing user id in context")
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	userID, ok := userIDVal.(string)
	if !ok {
		slog.Error("GetMe: user id in context is not string", "value", userIDVal)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.service.FindUserByID(r.Context(), userID)
	if err != nil {
		slog.Error("GetMe: user not found", "user_id", userID, "error", err)
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	json.Write(w, http.StatusOK, user)
}
