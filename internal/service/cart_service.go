package service

import (
	"context"
	"github.com/thisgleammm/mantis-backend/internal/domain"
)

type CartService struct {
	repo domain.CartRepository
}

func NewCartService(repo domain.CartRepository) *CartService {
	return &CartService{repo: repo}
}

func (s *CartService) ListCarts(ctx context.Context, userID string) ([]domain.Cart, error) {
	return s.repo.ListByUserID(ctx, userID)
}

func (s *CartService) ListCartItems(ctx context.Context, cartID string) ([]domain.CartItem, error) {
	return s.repo.ListItems(ctx, cartID)
}

func (s *CartService) AddItemToCart(ctx context.Context, item domain.CartItem) (domain.CartItem, error) {
	return s.repo.AddItem(ctx, item)
}

func (s *CartService) UpdateItemQuantity(ctx context.Context, itemID string, quantity int32) (domain.CartItem, error) {
	return s.repo.UpdateItemQuantity(ctx, itemID, quantity)
}

func (s *CartService) RemoveItemFromCart(ctx context.Context, itemID string) error {
	return s.repo.RemoveItem(ctx, itemID)
}
