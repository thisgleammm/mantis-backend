package domain

import "context"

type ProductRepository interface {
	List(ctx context.Context, limit, offset int32) ([]Product, error)
	FindByID(ctx context.Context, id int64) (Product, error)
	FindBySlug(ctx context.Context, slug string) (Product, error)
	Create(ctx context.Context, p Product) (Product, error)
	ListImages(ctx context.Context, productID int64) ([]ProductImage, error)
	ListVariants(ctx context.Context, productID int64) ([]ProductVariant, error)
}

type UserRepository interface {
	Create(ctx context.Context, u User) (User, error)
	FindByID(ctx context.Context, id string) (User, error)
	FindByEmail(ctx context.Context, email string) (User, error)
	List(ctx context.Context) ([]User, error)
}

type CategoryRepository interface {
	List(ctx context.Context) ([]Category, error)
	FindByID(ctx context.Context, id int64) (Category, error)
}

type CartRepository interface {
	ListByUserID(ctx context.Context, userID string) ([]Cart, error)
	ListItems(ctx context.Context, cartID string) ([]CartItem, error)
	AddItem(ctx context.Context, item CartItem) (CartItem, error)
	UpdateItemQuantity(ctx context.Context, itemID string, quantity int32) (CartItem, error)
	RemoveItem(ctx context.Context, itemID string) error
}

type AuthRepository interface {
	// Auth specific DB ops if any
}
