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

// Checkout creates a new order from the user's active cart.
// @Summary Checkout
// @Description Create an order from cart items and specify a shipping address.
// @Tags Orders
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param request body object true "Shipping details"
// @Success 201 {object} domain.Order
// @Failure 401 {string} string "unauthorized"
// @Router /orders/checkout [post]
func (h *OrderHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		json.WriteError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	var req struct {
		ShippingAddress string `json:"shipping_address"`
	}
	if err := json.Read(w, r, &req); err != nil {
		json.WriteError(w, http.StatusBadRequest, "invalid request")
		return
	}
	if req.ShippingAddress == "" {
		json.WriteError(w, http.StatusBadRequest, "shipping_address is required")
		return
	}

	order, err := h.svc.Checkout(r.Context(), userID, req.ShippingAddress)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.Write(w, http.StatusCreated, order)
}

// ListOrders returns a list of orders for the current user.
// @Summary List User Orders
// @Description Get all orders associated with the authenticated user.
// @Tags Orders
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {array} domain.Order
// @Router /orders [get]
func (h *OrderHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(string)

	orders, err := h.svc.ListOrders(r.Context(), userID)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.Write(w, http.StatusOK, orders)
}
