package service

import (
	"context"
	"github.com/thisgleammm/mantis-backend/internal/domain"
)

type CategoryService struct {
	repo domain.CategoryRepository
}

func NewCategoryService(repo domain.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) ListCategories(ctx context.Context) ([]domain.Category, error) {
	return s.repo.List(ctx)
}

func (s *CategoryService) FindCategoryByID(ctx context.Context, id int64) (domain.Category, error) {
	return s.repo.FindByID(ctx, id)
}
