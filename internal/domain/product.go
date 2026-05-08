package domain

import "time"

type Product struct {
	ID             int64            `json:"id"`
	CategoryID     int64            `json:"category_id"`
	Name           string           `json:"name"`
	Slug           string           `json:"slug"`
	Description    string           `json:"description"`
	BasePrice      float64          `json:"base_price"`
	DiscountPrice  float64          `json:"discount_price"`
	Weight         int32            `json:"weight"`
	Specifications any              `json:"specifications"`
	RatingAverage  float64          `json:"rating_average"`
	RatingCount    int32            `json:"rating_count"`
	Images         []ProductImage   `json:"images,omitempty"`
	Variants       []ProductVariant `json:"variants,omitempty"`
	CreatedAt      time.Time        `json:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at"`
}

type ProductImage struct {
	ID        int64  `json:"id"`
	ImageUrl  string `json:"image_url"`
	SortOrder int32  `json:"sort_order"`
}

type ProductVariant struct {
	ID               int64   `json:"id"`
	VariantName      string  `json:"variant_name"`
	PriceExtra       float64 `json:"price_extra"`
	Stock            int32   `json:"stock"`
	StockKeepingUnit string  `json:"stock_keeping_unit"`
}
