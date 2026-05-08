package service

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/thisgleammm/mantis-backend/internal/domain"
)

type OrderService struct {
	orderRepo domain.OrderRepository
	cartRepo  domain.CartRepository
}

func NewOrderService(orderRepo domain.OrderRepository, cartRepo domain.CartRepository) *OrderService {
	return &OrderService{orderRepo: orderRepo, cartRepo: cartRepo}
}

func (s *OrderService) Checkout(ctx context.Context, userID, shippingAddress string) (domain.Order, error) {
	carts, err := s.cartRepo.ListByUserID(ctx, userID)
	if err != nil {
		return domain.Order{}, err
	}
	if len(carts) == 0 {
		return domain.Order{}, fmt.Errorf("cart is empty")
	}

	var allItems []domain.CartItem
	for _, cart := range carts {
		items, err := s.cartRepo.ListItems(ctx, cart.ID)
		if err != nil {
			return domain.Order{}, err
		}
		allItems = append(allItems, items...)
	}

	if len(allItems) == 0 {
		return domain.Order{}, fmt.Errorf("cart is empty")
	}

	var totalAmount float64
	for _, item := range allItems {
		unitPrice := item.ProductPrice
		if item.VariantPriceExtra != nil {
			unitPrice += *item.VariantPriceExtra
		}
		totalAmount += unitPrice * float64(item.Quantity)
	}

	invoiceNumber, err := generateInvoiceNumber()
	if err != nil {
		return domain.Order{}, err
	}

	order, err := s.orderRepo.Create(ctx, domain.Order{
		UserID:          userID,
		InvoiceNumber:   invoiceNumber,
		Status:          domain.OrderStatusPending,
		TotalAmount:     totalAmount,
		ShippingCost:    0,
		GrandTotal:      totalAmount,
		ShippingAddress: shippingAddress,
	})
	if err != nil {
		return domain.Order{}, err
	}

	for _, item := range allItems {
		unitPrice := item.ProductPrice
		if item.VariantPriceExtra != nil {
			unitPrice += *item.VariantPriceExtra
		}

		if err := s.orderRepo.CreateItem(ctx, domain.OrderItem{
			OrderID:          order.ID,
			ProductID:        item.ProductID,
			ProductVariantID: item.ProductVariantID,
			ProductName:      item.ProductName,
			Quantity:         item.Quantity,
			PriceAtPurchase:  unitPrice,
		}); err != nil {
			return domain.Order{}, err
		}
	}

	if err := s.orderRepo.ClearCart(ctx, userID); err != nil {
		return domain.Order{}, err
	}

	order.Items, _ = s.orderRepo.ListItems(ctx, order.ID)

	return order, nil
}

func (s *OrderService) ListOrders(ctx context.Context, userID string) ([]domain.Order, error) {
	orders, err := s.orderRepo.ListByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	for i := range orders {
		items, _ := s.orderRepo.ListItems(ctx, orders[i].ID)
		orders[i].Items = items
	}

	return orders, nil
}

func generateInvoiceNumber() (string, error) {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 6
	b := make([]byte, length)
	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		b[i] = charset[n.Int64()]
	}
	return fmt.Sprintf("INV-%s-%s", time.Now().Format("20060102"), string(b)), nil
}
