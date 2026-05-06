package products

import (
	"context"

	repo "github.com/thisgleammm/mantis-backend/internal/adapters/postgresql/sqlc"
)

type ProductDetail struct {
	Product  repo.Product                  `json:"product"`
	Images   []repo.ListProductImagesRow   `json:"images"`
	Variants []repo.ListProductVariantsRow `json:"variants"`
}

type Service interface {
	ListProducts(ctx context.Context, limit, offset int32) ([]repo.ListProductsRow, error)
	FindProductByID(ctx context.Context, id int64) (ProductDetail, error)
	FindProductBySlug(ctx context.Context, slug string) (ProductDetail, error)
	CreateProduct(ctx context.Context, params repo.CreateProductParams) (repo.CreateProductRow, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) ListProducts(ctx context.Context, limit, offset int32) ([]repo.ListProductsRow, error) {
	return s.repo.ListProducts(ctx, repo.ListProductsParams{
		Limit:  limit,
		Offset: offset,
	})
}

func (s *svc) FindProductByID(ctx context.Context, id int64) (ProductDetail, error) {
	product, err := s.repo.FindProductByID(ctx, id)
	if err != nil {
		return ProductDetail{}, err
	}

	// Map FindProductByIDRow to repo.Product
	p := repo.Product{
		ID:             product.ID,
		CategoryID:     product.CategoryID,
		Name:           product.Name,
		Slug:           product.Slug,
		Description:    product.Description,
		BasePrice:      product.BasePrice,
		DiscountPrice:  product.DiscountPrice,
		Weight:         product.Weight,
		Specifications: product.Specifications,
		RatingAverage:  product.RatingAverage,
		RatingCount:    product.RatingCount,
		CreatedAt:      product.CreatedAt,
		UpdatedAt:      product.UpdatedAt,
	}

	return s.fetchProductExtras(ctx, p)
}

func (s *svc) FindProductBySlug(ctx context.Context, slug string) (ProductDetail, error) {
	product, err := s.repo.FindProductBySlug(ctx, slug)
	if err != nil {
		return ProductDetail{}, err
	}

	p := repo.Product{
		ID:             product.ID,
		CategoryID:     product.CategoryID,
		Name:           product.Name,
		Slug:           product.Slug,
		Description:    product.Description,
		BasePrice:      product.BasePrice,
		DiscountPrice:  product.DiscountPrice,
		Weight:         product.Weight,
		Specifications: product.Specifications,
		RatingAverage:  product.RatingAverage,
		RatingCount:    product.RatingCount,
		CreatedAt:      product.CreatedAt,
		UpdatedAt:      product.UpdatedAt,
	}

	return s.fetchProductExtras(ctx, p)
}

func (s *svc) fetchProductExtras(ctx context.Context, p repo.Product) (ProductDetail, error) {
	images, err := s.repo.ListProductImages(ctx, p.ID)
	if err != nil {
		images = []repo.ListProductImagesRow{}
	}

	variants, err := s.repo.ListProductVariants(ctx, p.ID)
	if err != nil {
		variants = []repo.ListProductVariantsRow{}
	}

	return ProductDetail{
		Product:  p,
		Images:   images,
		Variants: variants,
	}, nil
}

func (s *svc) CreateProduct(ctx context.Context, params repo.CreateProductParams) (repo.CreateProductRow, error) {
	return s.repo.CreateProduct(ctx, params)
}
