package categories

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

type categoryResponse struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	Slug           string `json:"slug"`
	Description    string `json:"description"`
	IconImageUrl   string `json:"icon_image_url"`
	BannerImageUrl string `json:"banner_image_url"`
	IsActive       bool   `json:"is_active"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

// ListCategories godoc
// @Summary List all categories
// @Description Get a list of all categories
// @Tags categories
// @Accept  json
// @Produce  json
// @Success 200 {array} categoryResponse
// @Router /categories [get]
func (h *handler) ListCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.ListCategories(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, categories)
}

// FindCategoryByID godoc
// @Summary Find category by ID
// @Description Get a category by its ID
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Success 200 {object} categoryResponse
// @Failure 400 {string} string "invalid category id"
// @Failure 404 {string} string "not found"
// @Router /categories/{id} [get]
func (h *handler) FindCategoryByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid category id", http.StatusBadRequest)
		return
	}

	category, err := h.service.FindCategoryByID(r.Context(), id)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.Write(w, http.StatusOK, category)
}
