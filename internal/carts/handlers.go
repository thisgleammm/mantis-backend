package carts

import (
	"log/slog"
	"net/http"

	"github.com/thisgleammm/mantis-backend/internal/auth"
	"github.com/thisgleammm/mantis-backend/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{service: service}
}

type cartResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// ListCarts godoc
// @Summary List user's carts
// @Description Get a list of carts belonging to the authenticated user
// @Tags carts
// @Accept  json
// @Produce  json
// @Success 200 {array} cartResponse
// @Failure 401 {string} string "unauthorized"
// @Security ApiKeyAuth
// @Router /carts [get]
func (h *handler) ListCarts(w http.ResponseWriter, r *http.Request) {
	userIDVal := r.Context().Value(auth.UserIDKey)
	if userIDVal == nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	userID, ok := userIDVal.(string)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	carts, err := h.service.ListCarts(r.Context(), userID)
	if err != nil {
		slog.Error("ListCarts failed", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, carts)
}
