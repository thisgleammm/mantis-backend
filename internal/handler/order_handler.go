package handler

import (
	"net/http"

	"github.com/thisgleammm/mantis-backend/internal/json"
	"github.com/thisgleammm/mantis-backend/internal/middleware"
	"github.com/thisgleammm/mantis-backend/internal/service"
)

type OrderHandler struct {
	svc *service.OrderService
}

func NewOrderHandler(svc *service.OrderService) *OrderHandler {
	return &OrderHandler{svc: svc}
}

func (h *OrderHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(string)

	var req struct {
		ShippingAddress string `json:"shipping_address"`
	}
	if err := json.Read(w, r, &req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	if req.ShippingAddress == "" {
		http.Error(w, "shipping_address is required", http.StatusBadRequest)
		return
	}

	order, err := h.svc.Checkout(r.Context(), userID, req.ShippingAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusCreated, order)
}

func (h *OrderHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(string)

	orders, err := h.svc.ListOrders(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, orders)
}
