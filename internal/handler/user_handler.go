package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/thisgleammm/mantis-backend/internal/json"
	"github.com/thisgleammm/mantis-backend/internal/middleware"
	"github.com/thisgleammm/mantis-backend/internal/service"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

// ListUsers returns a list of all users.
// @Summary List Users
// @Description Get a list of all registered users.
// @Tags Users
// @Produce json
// @Success 200 {array} domain.User
// @Router /users [get]
func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.svc.ListUsers(r.Context())
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.Write(w, http.StatusOK, users)
}

// FindUserByID returns a single user by their ID.
// @Summary Get User by ID
// @Description Get information about a user using their unique ID.
// @Tags Users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} domain.User
// @Failure 404 {string} string "not found"
// @Router /users/{id} [get]
func (h *UserHandler) FindUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	user, err := h.svc.FindUserByID(r.Context(), id)
	if err != nil {
		json.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	json.Write(w, http.StatusOK, user)
}

// GetMe returns the currently authenticated user's profile.
// @Summary Get Current User
// @Description Get the profile of the currently logged-in user.
// @Tags Users
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} domain.User
// @Failure 401 {string} string "unauthorized"
// @Router /users/me [get]
func (h *UserHandler) GetMe(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		json.WriteError(w, http.StatusUnauthorized, "unauthorized")
		return
	}
	user, err := h.svc.FindUserByID(r.Context(), userID)
	if err != nil {
		json.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	json.Write(w, http.StatusOK, user)
}
