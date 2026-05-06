package carts

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgtype"
	repo "github.com/thisgleammm/mantis-backend/internal/adapters/postgresql/sqlc"
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

type cartItemResponse struct {
	ID                string   `json:"id"`
	CartID            string   `json:"cart_id"`
	ProductID         int64    `json:"product_id"`
	ProductVariantID  *int64   `json:"product_variant_id,omitempty"`
	Quantity          int32    `json:"quantity"`
	ProductName       string   `json:"product_name"`
	ProductSlug       string   `json:"product_slug"`
	ProductPrice      float64  `json:"product_price"`
	VariantName       *string  `json:"variant_name,omitempty"`
	VariantPriceExtra *float64 `json:"variant_price_extra,omitempty"`
	CreatedAt         string   `json:"created_at"`
	UpdatedAt         string   `json:"updated_at"`
}

type addItemResponse struct {
	ID               string `json:"id"`
	CartID           string `json:"cart_id"`
	ProductID        int64  `json:"product_id"`
	ProductVariantID *int64 `json:"product_variant_id,omitempty"`
	Quantity         int32  `json:"quantity"`
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

// ListCartItems godoc
// @Summary List items in a cart
// @Description Get all items within a specific cart
// @Tags carts
// @Accept  json
// @Produce  json
// @Param id path string true "Cart ID"
// @Success 200 {array} cartItemResponse
// @Security ApiKeyAuth
// @Router /carts/{id}/items [get]
func (h *handler) ListCartItems(w http.ResponseWriter, r *http.Request) {
	cartID := chi.URLParam(r, "id")
	items, err := h.service.ListCartItems(r.Context(), cartID)
	if err != nil {
		slog.Error("ListCartItems failed", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	var res []cartItemResponse
	for _, i := range items {
		price, _ := i.ProductPrice.Float64Value()
		var variantName *string
		if i.VariantName.Valid {
			variantName = &i.VariantName.String
		}
		var variantPriceExtra *float64
		if i.VariantPriceExtra.Valid {
			vpe, _ := i.VariantPriceExtra.Float64Value()
			variantPriceExtra = &vpe.Float64
		}
		var variantID *int64
		if i.ProductVariantID.Valid {
			variantID = &i.ProductVariantID.Int64
		}

		res = append(res, cartItemResponse{
			ID:                i.ID.String(),
			CartID:            i.CartID.String(),
			ProductID:         i.ProductID,
			ProductVariantID:  variantID,
			Quantity:          i.Quantity,
			ProductName:       i.ProductName,
			ProductSlug:       i.ProductSlug,
			ProductPrice:      price.Float64,
			VariantName:       variantName,
			VariantPriceExtra: variantPriceExtra,
			CreatedAt:         i.CreatedAt.Time.String(),
			UpdatedAt:         i.UpdatedAt.Time.String(),
		})
	}

	json.Write(w, http.StatusOK, res)
}

type addItemRequest struct {
	ProductID        int64  `json:"product_id"`
	ProductVariantID *int64 `json:"product_variant_id"`
	Quantity         int32  `json:"quantity"`
}

// AddItemToCart godoc
// @Summary Add item to cart
// @Description Add a product or variant to a cart. If exists, increments quantity.
// @Tags carts
// @Accept  json
// @Produce  json
// @Param id path string true "Cart ID"
// @Param body body addItemRequest true "Item Details"
// @Success 201 {object} addItemResponse
// @Security ApiKeyAuth
// @Router /carts/{id}/items [post]
func (h *handler) AddItemToCart(w http.ResponseWriter, r *http.Request) {
	cartIDStr := chi.URLParam(r, "id")
	var cartID pgtype.UUID
	if err := cartID.Scan(cartIDStr); err != nil {
		http.Error(w, "invalid cart id", http.StatusBadRequest)
		return
	}

	var req addItemRequest
	if err := json.Read(w, r, &req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	arg := repo.AddItemToCartParams{
		CartID:    cartID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}
	if req.ProductVariantID != nil {
		arg.ProductVariantID = pgtype.Int8{Int64: *req.ProductVariantID, Valid: true}
	}

	item, err := h.service.AddItemToCart(r.Context(), arg)
	if err != nil {
		slog.Error("AddItemToCart failed", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	var vID *int64
	if item.ProductVariantID.Valid {
		vID = &item.ProductVariantID.Int64
	}

	json.Write(w, http.StatusCreated, addItemResponse{
		ID:               item.ID.String(),
		CartID:           item.CartID.String(),
		ProductID:        item.ProductID,
		ProductVariantID: vID,
		Quantity:         item.Quantity,
	})
}

type updateQuantityRequest struct {
	Quantity int32 `json:"quantity"`
}

// UpdateItemQuantity godoc
// @Summary Update item quantity
// @Description Update the quantity of a specific item in the cart
// @Tags carts
// @Accept  json
// @Produce  json
// @Param id path string true "Item ID"
// @Param body body updateQuantityRequest true "New Quantity"
// @Success 200 {object} addItemResponse
// @Security ApiKeyAuth
// @Router /carts/items/{id} [patch]
func (h *handler) UpdateItemQuantity(w http.ResponseWriter, r *http.Request) {
	itemIDStr := chi.URLParam(r, "id")
	var itemID pgtype.UUID
	if err := itemID.Scan(itemIDStr); err != nil {
		http.Error(w, "invalid item id", http.StatusBadRequest)
		return
	}

	var req updateQuantityRequest
	if err := json.Read(w, r, &req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	item, err := h.service.UpdateItemQuantity(r.Context(), repo.UpdateItemQuantityParams{
		ID:       itemID,
		Quantity: req.Quantity,
	})
	if err != nil {
		slog.Error("UpdateItemQuantity failed", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	var vID *int64
	if item.ProductVariantID.Valid {
		vID = &item.ProductVariantID.Int64
	}

	json.Write(w, http.StatusOK, addItemResponse{
		ID:               item.ID.String(),
		CartID:           item.CartID.String(),
		ProductID:        item.ProductID,
		ProductVariantID: vID,
		Quantity:         item.Quantity,
	})
}

// RemoveItemFromCart godoc
// @Summary Remove item from cart
// @Description Delete a specific item from the cart
// @Tags carts
// @Param id path string true "Item ID"
// @Success 204 "No Content"
// @Security ApiKeyAuth
// @Router /carts/items/{id} [delete]
func (h *handler) RemoveItemFromCart(w http.ResponseWriter, r *http.Request) {
	itemID := chi.URLParam(r, "id")
	if err := h.service.RemoveItemFromCart(r.Context(), itemID); err != nil {
		slog.Error("RemoveItemFromCart failed", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
