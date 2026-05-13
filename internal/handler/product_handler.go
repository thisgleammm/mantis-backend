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

// ListProducts returns a list of products with pagination.
// @Summary List Products
// @Description Get a paginated list of products.
// @Tags Products
// @Produce json
// @Param limit query int false "Limit (default 20)"
// @Param offset query int false "Offset (default 0)"
// @Success 200 {array} domain.Product
// @Router /products [get]
func (h *ProductHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	cursorStr := r.URL.Query().Get("cursor")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 20
	}

	products, err := h.svc.ListProducts(r.Context(), int32(limit), cursorStr)
	if err != nil {
		slog.Error("ListProducts failed", "error", err)
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.Write(w, http.StatusOK, products)
}

// FindProductBySlug returns a single product by its slug.
// @Summary Get Product by Slug
// @Description Get detailed information about a product using its slug.
// @Tags Products
// @Produce json
// @Param slug path string true "Product slug"
// @Success 200 {object} domain.Product
// @Failure 404 {string} string "not found"
// @Router /products/{slug} [get]
func (h *ProductHandler) FindProductBySlug(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	product, err := h.svc.FindProductBySlug(r.Context(), slug)
	if err != nil {
		json.WriteError(w, http.StatusNotFound, err.Error())
		return
	}

	json.Write(w, http.StatusOK, product)
}

// CreateProduct creates a new product.
// @Summary Create Product
// @Description Create a new product entry.
// @Tags Products
// @Accept json
// @Produce json
// @Param product body domain.Product true "Product object"
// @Success 201 {object} domain.Product
// @Router /products [post]
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p domain.Product
	if err := json.Read(w, r, &p); err != nil {
		json.WriteError(w, http.StatusBadRequest, "invalid request")
		return
	}

	created, err := h.svc.CreateProduct(r.Context(), p)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.Write(w, http.StatusCreated, created)
}
