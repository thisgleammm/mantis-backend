package handler

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/thisgleammm/mantis-backend/internal/domain"
	"github.com/thisgleammm/mantis-backend/internal/json"
	"github.com/thisgleammm/mantis-backend/internal/service"
)

type ProductHandler struct {
	svc *service.ProductService
}

func NewProductHandler(svc *service.ProductService) *ProductHandler {
	return &ProductHandler{svc: svc}
}

func (h *ProductHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 20
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	products, err := h.svc.ListProducts(r.Context(), int32(limit), int32(offset))
	if err != nil {
		slog.Error("ListProducts failed", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, products)
}

func (h *ProductHandler) FindProductBySlug(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	product, err := h.svc.FindProductBySlug(r.Context(), slug)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.Write(w, http.StatusOK, product)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p domain.Product
	if err := json.Read(w, r, &p); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	created, err := h.svc.CreateProduct(r.Context(), p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusCreated, created)
}
