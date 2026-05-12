package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/thisgleammm/mantis-backend/internal/json"
	"github.com/thisgleammm/mantis-backend/internal/service"
)

type CategoryHandler struct {
	svc *service.CategoryService
}

func NewCategoryHandler(svc *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{svc: svc}
}

// ListCategories returns a list of all categories.
// @Summary List Categories
// @Description Get a list of all product categories.
// @Tags Categories
// @Produce json
// @Success 200 {array} domain.Category
// @Router /categories [get]
func (h *CategoryHandler) ListCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.svc.ListCategories(r.Context())
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.Write(w, http.StatusOK, categories)
}

// FindCategoryByID returns a single category by its ID.
// @Summary Get Category by ID
// @Description Get information about a category using its numeric ID.
// @Tags Categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} domain.Category
// @Failure 400 {string} string "invalid category id"
// @Failure 404 {string} string "not found"
// @Router /categories/{id} [get]
func (h *CategoryHandler) FindCategoryByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, "invalid category id")
		return
	}
	category, err := h.svc.FindCategoryByID(r.Context(), id)
	if err != nil {
		json.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	json.Write(w, http.StatusOK, category)
}
