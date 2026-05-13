package domain

import "context"

type AddressRepository interface {
	ListByUserID(ctx context.Context, userID string) ([]Address, error)
	Create(ctx context.Context, address Address) (Address, error)
	Delete(ctx context.Context, id string, userID string) error
	SetPrimary(ctx context.Context, id string, userID string) error
}
