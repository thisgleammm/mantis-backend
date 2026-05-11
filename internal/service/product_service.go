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

func (s *ProductService) ListProducts(ctx context.Context, limit, offset int32) ([]domain.Product, error) {
	products, err := s.repo.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	for i := range products {
		images, _ := s.repo.ListImages(ctx, products[i].ID)
		products[i].Images = images
	}

	return products, nil
}

func (s *ProductService) FindProductBySlug(ctx context.Context, slug string) (domain.Product, error) {
	product, err := s.repo.FindBySlug(ctx, slug)
	if err != nil {
		return domain.Product{}, err
	}

	images, _ := s.repo.ListImages(ctx, product.ID)
	variants, _ := s.repo.ListVariants(ctx, product.ID)

	product.Images = images
	product.Variants = variants

	return product, nil
}

func (s *ProductService) CreateProduct(ctx context.Context, p domain.Product) (domain.Product, error) {
	return s.repo.Create(ctx, p)
}
