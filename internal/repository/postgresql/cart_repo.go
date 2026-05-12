package postgresql

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	repo "github.com/thisgleammm/mantis-backend/internal/adapters/postgresql/sqlc"
	"github.com/thisgleammm/mantis-backend/internal/domain"
)

type CartRepository struct {
	q *repo.Queries
}

func NewCartRepository(q *repo.Queries) *CartRepository {
	return &CartRepository{q: q}
}

func (r *CartRepository) ListByUserID(ctx context.Context, userID string) ([]domain.Cart, error) {
	var userUUID pgtype.UUID
	if err := userUUID.Scan(userID); err != nil {
		return nil, err
	}

	rows, err := r.q.ListCarts(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	var carts []domain.Cart
	for _, row := range rows {
		carts = append(carts, domain.Cart{
			ID:        row.ID.String(),
			UserID:    row.UserID.String(),
			CreatedAt: row.CreatedAt.Time,
			UpdatedAt: row.UpdatedAt.Time,
		})
	}
	return carts, nil
}

func (r *CartRepository) ListItems(ctx context.Context, cartID string) ([]domain.CartItem, error) {
	var cartUUID pgtype.UUID
	if err := cartUUID.Scan(cartID); err != nil {
		return nil, err
	}

	rows, err := r.q.ListCartItems(ctx, cartUUID)
	if err != nil {
		return nil, err
	}

	var items []domain.CartItem
	for _, row := range rows {
		price, _ := row.ProductPrice.Float64Value()
		var vName *string
		if row.VariantName.Valid {
			vName = &row.VariantName.String
		}
		var vPriceExtra *float64
		if row.VariantPriceExtra.Valid {
			vpe, _ := row.VariantPriceExtra.Float64Value()
			vPriceExtra = &vpe.Float64
		}
		var vID *int64
		if row.ProductVariantID.Valid {
			vID = &row.ProductVariantID.Int64
		}

		var pImage *string
		if row.ProductImage != "" {
			img := row.ProductImage
			pImage = &img
		}

		items = append(items, domain.CartItem{
			ID:                row.ID.String(),
			CartID:            row.CartID.String(),
			ProductID:         row.ProductID,
			ProductVariantID:  vID,
			Quantity:          row.Quantity,
			ProductName:       row.ProductName,
			ProductSlug:       row.ProductSlug,
			ProductPrice:      price.Float64,
			ProductImage:      pImage,
			VariantName:       vName,
			VariantPriceExtra: vPriceExtra,
			CreatedAt:         row.CreatedAt.Time,
			UpdatedAt:         row.UpdatedAt.Time,
		})
	}
	return items, nil
}

func (r *CartRepository) AddItem(ctx context.Context, item domain.CartItem) (domain.CartItem, error) {
	var cartUUID pgtype.UUID
	if err := cartUUID.Scan(item.CartID); err != nil {
		return domain.CartItem{}, err
	}

	arg := repo.AddItemToCartParams{
		CartID:    cartUUID,
		ProductID: item.ProductID,
		Quantity:  item.Quantity,
	}
	if item.ProductVariantID != nil {
		arg.ProductVariantID = pgtype.Int8{Int64: *item.ProductVariantID, Valid: true}
	}

	row, err := r.q.AddItemToCart(ctx, arg)
	if err != nil {
		return domain.CartItem{}, err
	}

	var vID *int64
	if row.ProductVariantID.Valid {
		vID = &row.ProductVariantID.Int64
	}

	return domain.CartItem{
		ID:               row.ID.String(),
		CartID:           row.CartID.String(),
		ProductID:        row.ProductID,
		ProductVariantID: vID,
		Quantity:         row.Quantity,
		CreatedAt:        row.CreatedAt.Time,
		UpdatedAt:        row.UpdatedAt.Time,
	}, nil
}

func (r *CartRepository) UpdateItemQuantity(ctx context.Context, itemID string, quantity int32) (domain.CartItem, error) {
	var uuid pgtype.UUID
	_ = uuid.Scan(itemID)

	row, err := r.q.UpdateItemQuantity(ctx, repo.UpdateItemQuantityParams{
		ID:       uuid,
		Quantity: quantity,
	})
	if err != nil {
		return domain.CartItem{}, err
	}

	var vID *int64
	if row.ProductVariantID.Valid {
		vID = &row.ProductVariantID.Int64
	}

	return domain.CartItem{
		ID:               row.ID.String(),
		CartID:           row.CartID.String(),
		ProductID:        row.ProductID,
		ProductVariantID: vID,
		Quantity:         row.Quantity,
		CreatedAt:        row.CreatedAt.Time,
		UpdatedAt:        row.UpdatedAt.Time,
	}, nil
}

func (r *CartRepository) RemoveItem(ctx context.Context, itemID string) error {
	var uuid pgtype.UUID
	_ = uuid.Scan(itemID)
	return r.q.RemoveItemFromCart(ctx, uuid)
}
