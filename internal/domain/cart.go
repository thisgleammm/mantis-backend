package domain

import "time"

type Cart struct {
	ID        string     `json:"id"`
	UserID    string     `json:"user_id"`
	Items     []CartItem `json:"items,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type CartItem struct {
	ID                string    `json:"id"`
	CartID            string    `json:"cart_id"`
	ProductID         int64     `json:"product_id"`
	ProductVariantID  *int64    `json:"product_variant_id,omitempty"`
	Quantity          int32     `json:"quantity"`
	ProductName       string    `json:"product_name"`
	ProductSlug       string    `json:"product_slug"`
	ProductPrice      float64   `json:"product_price"`
	ProductImage      *string   `json:"product_image,omitempty"`
	VariantName       *string   `json:"variant_name,omitempty"`
	VariantPriceExtra *float64  `json:"variant_price_extra,omitempty"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
