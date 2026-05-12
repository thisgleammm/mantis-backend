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

func (h *CartHandler) ListCarts(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	carts, err := h.svc.ListCarts(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.Write(w, http.StatusOK, carts)
}

func (h *CartHandler) ListCartItems(w http.ResponseWriter, r *http.Request) {
	cartID := chi.URLParam(r, "id")
	items, err := h.svc.ListCartItems(r.Context(), cartID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.Write(w, http.StatusOK, items)
}

func (h *CartHandler) AddItemToCart(w http.ResponseWriter, r *http.Request) {
	cartID := chi.URLParam(r, "id")
	var item domain.CartItem
	if err := json.Read(w, r, &item); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	item.CartID = cartID
	created, err := h.svc.AddItemToCart(r.Context(), item)
	if err != nil {
		slog.Error("AddItemToCart failed", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.Write(w, http.StatusCreated, created)
}

func (h *CartHandler) UpdateItemQuantity(w http.ResponseWriter, r *http.Request) {
	itemID := chi.URLParam(r, "id")
	var req struct {
		Quantity int32 `json:"quantity"`
	}
	if err := json.Read(w, r, &req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	updated, err := h.svc.UpdateItemQuantity(r.Context(), itemID, req.Quantity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.Write(w, http.StatusOK, updated)
}

func (h *CartHandler) RemoveItemFromCart(w http.ResponseWriter, r *http.Request) {
	itemID := chi.URLParam(r, "id")
	if err := h.svc.RemoveItemFromCart(r.Context(), itemID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
