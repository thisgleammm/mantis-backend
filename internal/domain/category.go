package domain

import "time"

type Category struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	Slug           string    `json:"slug"`
	Description    string    `json:"description"`
	IconImageUrl   string    `json:"icon_image_url"`
	BannerImageUrl string    `json:"banner_image_url"`
	IsActive       bool      `json:"is_active"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
