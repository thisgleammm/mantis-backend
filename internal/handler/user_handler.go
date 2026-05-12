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

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.svc.ListUsers(r.Context())
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.Write(w, http.StatusOK, users)
}

func (h *UserHandler) FindUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	user, err := h.svc.FindUserByID(r.Context(), id)
	if err != nil {
		json.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	json.Write(w, http.StatusOK, user)
}

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
