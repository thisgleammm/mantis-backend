package postgresql

import (
	"context"
	"encoding/json"
	"math/big"

	"github.com/jackc/pgx/v5/pgtype"
	repo "github.com/thisgleammm/mantis-backend/internal/adapters/postgresql/sqlc"
	"github.com/thisgleammm/mantis-backend/internal/domain"
)

type ProductRepository struct {
	q *repo.Queries
}

func NewProductRepository(q *repo.Queries) *ProductRepository {
	return &ProductRepository{q: q}
}

func (r *ProductRepository) List(ctx context.Context, limit int32, cursor any) ([]domain.Product, error) {
	var createdAt pgtype.Timestamptz
	if cursor != nil {
		if s, ok := cursor.(string); ok && s != "" {
			createdAt.UnmarshalJSON([]byte(`"` + s + `"`))
		}
	}

	rows, err := r.q.ListProducts(ctx, repo.ListProductsParams{
		CreatedAt: createdAt,
		Limit:     limit,
	})
	if err != nil {
		return nil, err
	}

	var products []domain.Product
	for _, row := range rows {
		basePrice, _ := row.BasePrice.Float64Value()
		discountPrice, _ := row.DiscountPrice.Float64Value()
		ratingAverage, _ := row.RatingAverage.Float64Value()

		products = append(products, domain.Product{
			ID:            row.ID,
			CategoryID:    row.CategoryID.Int64,
			Name:          row.Name,
			Slug:          row.Slug,
			BasePrice:     basePrice.Float64,
			DiscountPrice: discountPrice.Float64,
			RatingAverage: ratingAverage.Float64,
			RatingCount:   row.RatingCount,
			CreatedAt:     row.CreatedAt.Time,
			Images: []domain.ProductImage{
				{ImageUrl: row.MainImage},
			},
		})
	}
	return products, nil
}

func (r *ProductRepository) FindByID(ctx context.Context, id int64) (domain.Product, error) {
	row, err := r.q.FindProductByID(ctx, id)
	if err != nil {
		return domain.Product{}, err
	}

	basePrice, _ := row.BasePrice.Float64Value()
	discountPrice, _ := row.DiscountPrice.Float64Value()
	ratingAverage, _ := row.RatingAverage.Float64Value()

	return domain.Product{
		ID:             row.ID,
		CategoryID:     row.CategoryID.Int64,
		Name:           row.Name,
		Slug:           row.Slug,
		Description:    row.Description.String,
		BasePrice:      basePrice.Float64,
		DiscountPrice:  discountPrice.Float64,
		Weight:         row.Weight,
		Specifications: row.Specifications,
		RatingAverage:  ratingAverage.Float64,
		RatingCount:    row.RatingCount,
		CreatedAt:      row.CreatedAt.Time,
		UpdatedAt:      row.UpdatedAt.Time,
	}, nil
}

func (r *ProductRepository) FindBySlug(ctx context.Context, slug string) (domain.Product, error) {
	row, err := r.q.FindProductDetailBySlug(ctx, slug)
	if err != nil {
		return domain.Product{}, err
	}

	basePrice, _ := row.BasePrice.Float64Value()
	discountPrice, _ := row.DiscountPrice.Float64Value()
	ratingAverage, _ := row.RatingAverage.Float64Value()

	var images []domain.ProductImage
	if row.Images != nil {
		if b, ok := row.Images.([]byte); ok {
			json.Unmarshal(b, &images)
		}
	}

	var variants []domain.ProductVariant
	if row.Variants != nil {
		if b, ok := row.Variants.([]byte); ok {
			json.Unmarshal(b, &variants)
		}
	}

	return domain.Product{
		ID:             row.ID,
		CategoryID:     row.CategoryID.Int64,
		Name:           row.Name,
		Slug:           row.Slug,
		Description:    row.Description.String,
		BasePrice:      basePrice.Float64,
		DiscountPrice:  discountPrice.Float64,
		Weight:         row.Weight,
		Specifications: row.Specifications,
		RatingAverage:  ratingAverage.Float64,
		RatingCount:    row.RatingCount,
		CreatedAt:      row.CreatedAt.Time,
		Images:         images,
		Variants:       variants,
	}, nil
}

func (r *ProductRepository) Create(ctx context.Context, p domain.Product) (domain.Product, error) {
	var specs []byte
	if p.Specifications != nil {
		specs, _ = json.Marshal(p.Specifications)
	}

	row, err := r.q.CreateProduct(ctx, repo.CreateProductParams{
		CategoryID:     pgtype.Int8{Int64: p.CategoryID, Valid: p.CategoryID != 0},
		Name:           p.Name,
		Slug:           p.Slug,
		Description:    pgtype.Text{String: p.Description, Valid: p.Description != ""},
		BasePrice:      pgtype.Numeric{Int: big.NewInt(int64(p.BasePrice * 100)), Exp: -2, Valid: true},
		DiscountPrice:  pgtype.Numeric{Int: big.NewInt(int64(p.DiscountPrice * 100)), Exp: -2, Valid: p.DiscountPrice != 0},
		Weight:         p.Weight,
		Specifications: specs,
	})
	if err != nil {
		return domain.Product{}, err
	}

	basePrice, _ := row.BasePrice.Float64Value()
	discountPrice, _ := row.DiscountPrice.Float64Value()

	return domain.Product{
		ID:            row.ID,
		CategoryID:    row.CategoryID.Int64,
		Name:          row.Name,
		Slug:          row.Slug,
		Description:   row.Description.String,
		BasePrice:     basePrice.Float64,
		DiscountPrice: discountPrice.Float64,
		Weight:        row.Weight,
		CreatedAt:     row.CreatedAt.Time,
		UpdatedAt:     row.UpdatedAt.Time,
	}, nil
}

func (r *ProductRepository) ListImages(ctx context.Context, productID int64) ([]domain.ProductImage, error) {
	rows, err := r.q.ListProductImages(ctx, productID)
	if err != nil {
		return nil, err
	}

	var images []domain.ProductImage
	for _, row := range rows {
		images = append(images, domain.ProductImage{
			ID:       row.ID,
			ImageUrl: row.ImageUrl,
		})
	}
	return images, nil
}

func (r *ProductRepository) ListVariants(ctx context.Context, productID int64) ([]domain.ProductVariant, error) {
	rows, err := r.q.ListProductVariants(ctx, productID)
	if err != nil {
		return nil, err
	}

	var variants []domain.ProductVariant
	for _, row := range rows {
		priceExtra, _ := row.PriceExtra.Float64Value()
		variants = append(variants, domain.ProductVariant{
			ID:               row.ID,
			VariantName:      row.VariantName,
			PriceExtra:       priceExtra.Float64,
			Stock:            row.Stock,
			StockKeepingUnit: row.StockKeepingUnit,
		})
	}
	return variants, nil
}
