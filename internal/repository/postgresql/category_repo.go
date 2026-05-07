package postgresql

import (
	"context"
	repo "github.com/thisgleammm/mantis-backend/internal/adapters/postgresql/sqlc"
	"github.com/thisgleammm/mantis-backend/internal/domain"
)

type CategoryRepository struct {
	q *repo.Queries
}

func NewCategoryRepository(q *repo.Queries) *CategoryRepository {
	return &CategoryRepository{q: q}
}

func (r *CategoryRepository) List(ctx context.Context) ([]domain.Category, error) {
	rows, err := r.q.ListCategories(ctx)
	if err != nil {
		return nil, err
	}

	var categories []domain.Category
	for _, row := range rows {
		categories = append(categories, domain.Category{
			ID:             row.ID,
			Name:           row.Name,
			Slug:           row.Slug,
			Description:    row.Description.String,
			IconImageUrl:   row.IconImageUrl.String,
			BannerImageUrl: row.BannerImageUrl.String,
			IsActive:       row.IsActive,
			CreatedAt:      row.CreatedAt.Time,
			UpdatedAt:      row.UpdatedAt.Time,
		})
	}
	return categories, nil
}

func (r *CategoryRepository) FindByID(ctx context.Context, id int64) (domain.Category, error) {
	row, err := r.q.FindCategoryByID(ctx, id)
	if err != nil {
		return domain.Category{}, err
	}

	return domain.Category{
		ID:             row.ID,
		Name:           row.Name,
		Slug:           row.Slug,
		Description:    row.Description.String,
		IconImageUrl:   row.IconImageUrl.String,
		BannerImageUrl: row.BannerImageUrl.String,
		IsActive:       row.IsActive,
		CreatedAt:      row.CreatedAt.Time,
		UpdatedAt:      row.UpdatedAt.Time,
	}, nil
}
