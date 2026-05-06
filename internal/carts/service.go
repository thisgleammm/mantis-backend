package carts

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	repo "github.com/thisgleammm/mantis-backend/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListCarts(ctx context.Context, userID string) ([]repo.Cart, error)
	ListCartItems(ctx context.Context, cartID string) ([]repo.ListCartItemsRow, error)
	AddItemToCart(ctx context.Context, arg repo.AddItemToCartParams) (repo.CartItem, error)
	UpdateItemQuantity(ctx context.Context, arg repo.UpdateItemQuantityParams) (repo.CartItem, error)
	RemoveItemFromCart(ctx context.Context, id string) error
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) ListCarts(ctx context.Context, userID string) ([]repo.Cart, error) {
	var userUUID pgtype.UUID
	if err := userUUID.Scan(userID); err != nil {
		return nil, err
	}
	return s.repo.ListCarts(ctx, userUUID)
}

func (s *svc) ListCartItems(ctx context.Context, cartID string) ([]repo.ListCartItemsRow, error) {
	var cartUUID pgtype.UUID
	if err := cartUUID.Scan(cartID); err != nil {
		return nil, err
	}
	return s.repo.ListCartItems(ctx, cartUUID)
}

func (s *svc) AddItemToCart(ctx context.Context, arg repo.AddItemToCartParams) (repo.CartItem, error) {
	return s.repo.AddItemToCart(ctx, arg)
}

func (s *svc) UpdateItemQuantity(ctx context.Context, arg repo.UpdateItemQuantityParams) (repo.CartItem, error) {
	return s.repo.UpdateItemQuantity(ctx, arg)
}

func (s *svc) RemoveItemFromCart(ctx context.Context, id string) error {
	var itemUUID pgtype.UUID
	if err := itemUUID.Scan(id); err != nil {
		return err
	}
	return s.repo.RemoveItemFromCart(ctx, itemUUID)
}
