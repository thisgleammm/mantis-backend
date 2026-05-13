package service

import (
	"context"
	"github.com/thisgleammm/mantis-backend/internal/domain"
)

type ProductService struct {
	repo domain.ProductRepository
}

func NewProductService(repo domain.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) ListProducts(ctx context.Context, limit int32, cursor any) ([]domain.Product, error) {
	products, err := s.repo.List(ctx, limit, cursor)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *ProductService) FindProductBySlug(ctx context.Context, slug string) (domain.Product, error) {
	return s.repo.FindBySlug(ctx, slug)
}

func (s *ProductService) CreateProduct(ctx context.Context, p domain.Product) (domain.Product, error) {
	return s.repo.Create(ctx, p)
}
