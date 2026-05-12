package handler

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/thisgleammm/mantis-backend/internal/domain"
	"github.com/thisgleammm/mantis-backend/internal/json"
	"github.com/thisgleammm/mantis-backend/internal/middleware"
	"github.com/thisgleammm/mantis-backend/internal/service"
)

type CartHandler struct {
	svc *service.CartService
}

func NewCartHandler(svc *service.CartService) *CartHandler {
	return &CartHandler{svc: svc}
}

// ListCarts returns all carts associated with the current user.
// @Summary List User Carts
// @Description Get a list of all shopping carts belonging to the authenticated user.
// @Tags Carts
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {array} domain.Cart
// @Failure 401 {string} string "unauthorized"
// @Router /carts [get]
func (h *CartHandler) ListCarts(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		json.WriteError(w, http.StatusUnauthorized, "unauthorized")
		return
	}
	carts, err := h.svc.ListCarts(r.Context(), userID)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.Write(w, http.StatusOK, carts)
}

// ListCartItems returns all items within a specific cart.
// @Summary List Cart Items
// @Description Get all products and their quantities in a specific cart.
// @Tags Carts
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Cart ID"
// @Success 200 {array} domain.CartItem
// @Router /carts/{id}/items [get]
func (h *CartHandler) ListCartItems(w http.ResponseWriter, r *http.Request) {
	cartID := chi.URLParam(r, "id")
	items, err := h.svc.ListCartItems(r.Context(), cartID)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.Write(w, http.StatusOK, items)
}

// AddItemToCart adds a product to a specific cart.
// @Summary Add Item to Cart
// @Description Add a new product or increase quantity of an existing product in a cart.
// @Tags Carts
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "Cart ID"
// @Param item body domain.CartItem true "Item details"
// @Success 201 {object} domain.CartItem
// @Router /carts/{id}/items [post]
func (h *CartHandler) AddItemToCart(w http.ResponseWriter, r *http.Request) {
	cartID := chi.URLParam(r, "id")
	var item domain.CartItem
	if err := json.Read(w, r, &item); err != nil {
		json.WriteError(w, http.StatusBadRequest, "invalid request")
		return
	}
	item.CartID = cartID
	created, err := h.svc.AddItemToCart(r.Context(), item)
	if err != nil {
		slog.Error("AddItemToCart failed", "error", err)
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.Write(w, http.StatusCreated, created)
}

// UpdateItemQuantity updates the quantity of a specific item in any cart.
// @Summary Update Item Quantity
// @Description Change the quantity of a cart item by its unique item ID.
// @Tags Carts
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "Item ID"
// @Param quantity body object true "Quantity object"
// @Success 200 {object} domain.CartItem
// @Router /carts/items/{id} [patch]
func (h *CartHandler) UpdateItemQuantity(w http.ResponseWriter, r *http.Request) {
	itemID := chi.URLParam(r, "id")
	var req struct {
		Quantity int32 `json:"quantity"`
	}
	if err := json.Read(w, r, &req); err != nil {
		json.WriteError(w, http.StatusBadRequest, "invalid request")
		return
	}
	updated, err := h.svc.UpdateItemQuantity(r.Context(), itemID, req.Quantity)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.Write(w, http.StatusOK, updated)
}

// RemoveItemFromCart removes an item from a cart.
// @Summary Remove Item from Cart
// @Description Delete a specific item from a shopping cart by its unique item ID.
// @Tags Carts
// @Security ApiKeyAuth
// @Param id path string true "Item ID"
// @Success 204
// @Router /carts/items/{id} [delete]
func (h *CartHandler) RemoveItemFromCart(w http.ResponseWriter, r *http.Request) {
	itemID := chi.URLParam(r, "id")
	if err := h.svc.RemoveItemFromCart(r.Context(), itemID); err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
