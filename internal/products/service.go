package products

import (
	"context"

	repo "github.com/thisgleammm/mantis-backend/internal/adapters/postgresql/sqlc"
)

type ProductDetail struct {
	repo.FindProductBySlugRow
	Images   []repo.ListProductImagesRow   `json:"images"`
	Variants []repo.ListProductVariantsRow `json:"variants"`
}

type Service interface {
	ListProducts(ctx context.Context, limit, offset int32) ([]repo.ListProductsRow, error)
	FindProductByID(ctx context.Context, id int64) (repo.FindProductByIDRow, error)
	FindProductBySlug(ctx context.Context, slug string) (ProductDetail, error)
	CreateProduct(ctx context.Context, params repo.CreateProductParams) (repo.CreateProductRow, error)
}

type svc struct {
	//repository
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{
		repo: repo,
	}
}

func (s *svc) ListProducts(ctx context.Context, limit, offset int32) ([]repo.ListProductsRow, error) {
	return s.repo.ListProducts(ctx, repo.ListProductsParams{
		Limit:  limit,
		Offset: offset,
	})
}

func (s *svc) FindProductByID(ctx context.Context, id int64) (repo.FindProductByIDRow, error) {
	return s.repo.FindProductByID(ctx, id)
}

func (s *svc) FindProductBySlug(ctx context.Context, slug string) (ProductDetail, error) {
	product, err := s.repo.FindProductBySlug(ctx, slug)
	if err != nil {
		return ProductDetail{}, err
	}

	images, err := s.repo.ListProductImages(ctx, product.ID)
	if err != nil {
		images = []repo.ListProductImagesRow{}
	}

	variants, err := s.repo.ListProductVariants(ctx, product.ID)
	if err != nil {
		variants = []repo.ListProductVariantsRow{}
	}

	return ProductDetail{
		FindProductBySlugRow: product,
		Images:               images,
		Variants:             variants,
	}, nil
}

func (s *svc) CreateProduct(ctx context.Context, params repo.CreateProductParams) (repo.CreateProductRow, error) {
	return s.repo.CreateProduct(ctx, params)
}
