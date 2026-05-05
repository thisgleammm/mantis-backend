package products

import (
	stdjson "encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgtype"
	repo "github.com/thisgleammm/mantis-backend/internal/adapters/postgresql/sqlc"
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

var _ = productResponse{}

// ListProducts godoc
// @Summary List all products
// @Description Get a list of all products
// @Tags products
// @Accept  json
// @Produce  json
// @Param limit query int false "Limit" default(20)
// @Param offset query int false "Offset" default(0)
// @Success 200 {array} productResponse
// @Router /products [get]
func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
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

	products, err := h.service.ListProducts(r.Context(), int32(limit), int32(offset))

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

type createProductRequest struct {
	CategoryID     int64   `json:"category_id"`
	Name           string  `json:"name"`
	Slug           string  `json:"slug"`
	Description    string  `json:"description"`
	BasePrice      float64 `json:"base_price"`
	DiscountPrice  float64 `json:"discount_price"`
	Weight         float64 `json:"weight"`
	Specifications any     `json:"specifications"` // Allow client to send JSON object
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product with the provided details
// @Tags products
// @Accept  json
// @Produce  json
// @Param product body createProductRequest true "Product details"
// @Success 201 {object} productResponse
// @Failure 400 {string} string "invalid request payload"
// @Failure 500 {string} string "internal server error"
// @Router /products [post]
func (h *handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req createProductRequest

	if err := json.Read(w, r, &req); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	// Map request to db params
	params := repo.CreateProductParams{
		CategoryID: pgtype.Int8{Int64: req.CategoryID, Valid: req.CategoryID != 0},
		Name:       req.Name,
		Slug:       req.Slug,
		Description: pgtype.Text{
			String: req.Description,
			Valid:  req.Description != "",
		},
	}

	// Convert float64 to pgtype.Numeric (simple approach: using float64 to numeric logic is tricky,
	// pgtype.Numeric requires setting Int/Exp, or we can scan from string)
	// Actually, pgtype.Numeric has a Scan method or we can use Float8 instead of Numeric if possible.
	// But let's use pgtype.Numeric.Scan(fmt.Sprintf("%f", req.BasePrice))
	// Wait, let's just use json mapping? No, we can set pgtype.Numeric by scanning string.
	_ = params.BasePrice.Scan(strconv.FormatFloat(req.BasePrice, 'f', -1, 64))
	_ = params.DiscountPrice.Scan(strconv.FormatFloat(req.DiscountPrice, 'f', -1, 64))
	params.Weight = int32(req.Weight)

	specs, _ := stdjson.Marshal(req.Specifications)
	params.Specifications = specs

	product, err := h.service.CreateProduct(r.Context(), params)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusCreated, product)
}
