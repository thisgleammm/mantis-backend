package carts

import (
	"log/slog"
	"net/http"

	"github.com/thisgleammm/mantis-backend/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{service: service}
}

// ListCarts godoc
// @Summary List all carts
// @Description Get a list of all carts
// @Tags carts
// @Accept  json
// @Produce  json
// @Success 200 {array} repo.Cart
// @Router /carts [get]
func (h *handler) ListCarts(w http.ResponseWriter, r *http.Request) {
	carts, err := h.service.ListCarts(r.Context())
	if err != nil {
		slog.Error("ListCarts failed", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, carts)
}
