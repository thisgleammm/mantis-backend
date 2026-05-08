package postgresql

import (
	"context"
	"math/big"

	"github.com/jackc/pgx/v5/pgtype"
	repo "github.com/thisgleammm/mantis-backend/internal/adapters/postgresql/sqlc"
	"github.com/thisgleammm/mantis-backend/internal/domain"
)

type OrderRepository struct {
	q *repo.Queries
}

func NewOrderRepository(q *repo.Queries) *OrderRepository {
	return &OrderRepository{q: q}
}

func (r *OrderRepository) Create(ctx context.Context, order domain.Order) (domain.Order, error) {
	var userUUID pgtype.UUID
	_ = userUUID.Scan(order.UserID)

	totalAmount := pgtype.Numeric{Int: big.NewInt(int64(order.TotalAmount * 100)), Exp: -2, Valid: true}
	shippingCost := pgtype.Numeric{Int: big.NewInt(int64(order.ShippingCost * 100)), Exp: -2, Valid: true}
	grandTotal := pgtype.Numeric{Int: big.NewInt(int64(order.GrandTotal * 100)), Exp: -2, Valid: true}

	row, err := r.q.CreateOrder(ctx, repo.CreateOrderParams{
		UserID:          userUUID,
		InvoiceNumber:   order.InvoiceNumber,
		Status:          repo.OrderStatus(order.Status),
		TotalAmount:     totalAmount,
		ShippingCost:    shippingCost,
		GrandTotal:      grandTotal,
		ShippingAddress: order.ShippingAddress,
	})
	if err != nil {
		return domain.Order{}, err
	}

	ta, _ := row.TotalAmount.Float64Value()
	sc, _ := row.ShippingCost.Float64Value()
	gt, _ := row.GrandTotal.Float64Value()

	return domain.Order{
		ID:              row.ID.String(),
		UserID:          row.UserID.String(),
		InvoiceNumber:   row.InvoiceNumber,
		Status:          domain.OrderStatus(row.Status),
		TotalAmount:     ta.Float64,
		ShippingCost:    sc.Float64,
		GrandTotal:      gt.Float64,
		ShippingAddress: row.ShippingAddress,
		TrackingNumber:  row.TrackingNumber.String,
		CourierName:     row.CourierName.String,
		CreatedAt:       row.CreatedAt.Time,
		UpdatedAt:       row.UpdatedAt.Time,
	}, nil
}

func (r *OrderRepository) CreateItem(ctx context.Context, item domain.OrderItem) error {
	var orderUUID pgtype.UUID
	_ = orderUUID.Scan(item.OrderID)

	var variantID pgtype.Int8
	if item.ProductVariantID != nil {
		variantID = pgtype.Int8{Int64: *item.ProductVariantID, Valid: true}
	}

	price := pgtype.Numeric{Int: big.NewInt(int64(item.PriceAtPurchase * 100)), Exp: -2, Valid: true}

	_, err := r.q.CreateOrderItem(ctx, repo.CreateOrderItemParams{
		OrderID:          orderUUID,
		ProductID:        item.ProductID,
		ProductVariantID: variantID,
		ProductName:      item.ProductName,
		Quantity:         item.Quantity,
		PriceAtPurchase:  price,
	})
	return err
}

func (r *OrderRepository) ListByUserID(ctx context.Context, userID string) ([]domain.Order, error) {
	var userUUID pgtype.UUID
	_ = userUUID.Scan(userID)

	rows, err := r.q.ListOrders(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	var orders []domain.Order
	for _, row := range rows {
		ta, _ := row.TotalAmount.Float64Value()
		sc, _ := row.ShippingCost.Float64Value()
		gt, _ := row.GrandTotal.Float64Value()

		orders = append(orders, domain.Order{
			ID:              row.ID.String(),
			UserID:          row.UserID.String(),
			InvoiceNumber:   row.InvoiceNumber,
			Status:          domain.OrderStatus(row.Status),
			TotalAmount:     ta.Float64,
			ShippingCost:    sc.Float64,
			GrandTotal:      gt.Float64,
			ShippingAddress: row.ShippingAddress,
			TrackingNumber:  row.TrackingNumber.String,
			CourierName:     row.CourierName.String,
			CreatedAt:       row.CreatedAt.Time,
			UpdatedAt:       row.UpdatedAt.Time,
		})
	}
	return orders, nil
}

func (r *OrderRepository) ListItems(ctx context.Context, orderID string) ([]domain.OrderItem, error) {
	var orderUUID pgtype.UUID
	_ = orderUUID.Scan(orderID)

	rows, err := r.q.ListOrderItems(ctx, orderUUID)
	if err != nil {
		return nil, err
	}

	var items []domain.OrderItem
	for _, row := range rows {
		price, _ := row.PriceAtPurchase.Float64Value()
		var variantID *int64
		if row.ProductVariantID.Valid {
			variantID = &row.ProductVariantID.Int64
		}

		items = append(items, domain.OrderItem{
			ID:               row.ID.String(),
			OrderID:          row.OrderID.String(),
			ProductID:        row.ProductID,
			ProductVariantID: variantID,
			ProductName:      row.ProductName,
			Quantity:         row.Quantity,
			PriceAtPurchase:  price.Float64,
			CreatedAt:        row.CreatedAt.Time,
			UpdatedAt:        row.UpdatedAt.Time,
		})
	}
	return items, nil
}

func (r *OrderRepository) ClearCart(ctx context.Context, userID string) error {
	var userUUID pgtype.UUID
	_ = userUUID.Scan(userID)
	return r.q.ClearCartItems(ctx, userUUID)
}
