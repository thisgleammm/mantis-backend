package products

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/thisgleammm/mantis-backend/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

type productResponse struct {
	ID            int64   `json:"id"`
	CategoryID    int64   `json:"category_id"`
	Name          string  `json:"name"`
	Slug          string  `json:"slug"`
	Description   string  `json:"description"`
	BasePrice     float64 `json:"base_price"`
	DiscountPrice float64 `json:"discount_price"`
	Weight        float64 `json:"weight"`
	RatingAverage float64 `json:"rating_average"`
	RatingCount   int32   `json:"rating_count"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

// ListProducts godoc
// @Summary List all products
// @Description Get a list of all products
// @Tags products
// @Accept  json
// @Produce  json
// @Success 200 {array} productResponse
// @Router /products [get]
func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {

	products, err := h.service.ListProducts(r.Context())

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, products)
}

// FindProductByID godoc
// @Summary Find product by ID
// @Description Get a product by its ID
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} productResponse
// @Failure 400 {string} string "invalid product id"
// @Failure 404 {string} string "not found"
// @Router /products/{id} [get]
func (h *handler) FindProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}

	product, err := h.service.FindProductByID(r.Context(), id)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.Write(w, http.StatusOK, product)
}

// FindProductBySlug godoc
// @Summary Find product by slug
// @Description Get a product by its slug
// @Tags products
// @Accept  json
// @Produce  json
// @Param slug path string true "Product Slug"
// @Success 200 {object} productResponse
// @Failure 400 {string} string "invalid product slug"
// @Failure 404 {string} string "not found"
// @Router /products/{slug} [get]
func (h *handler) FindProductBySlug(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	if slug == "" {
		http.Error(w, "invalid product slug", http.StatusBadRequest)
		return
	}

	product, err := h.service.FindProductBySlug(r.Context(), slug)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.Write(w, http.StatusOK, product)
}
