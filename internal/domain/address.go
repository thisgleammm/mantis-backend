package domain

import "time"

type Address struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	RecipientName string    `json:"recipient_name"`
	PhoneNumber   string    `json:"phone_number"`
	Province      string    `json:"province"`
	City          string    `json:"city"`
	District      string    `json:"district"`
	PostalCode    string    `json:"postal_code"`
	FullAddress   string    `json:"full_address"`
	Label         string    `json:"label"`
	Coordinates   string    `json:"coordinates"`
	IsPrimary     bool      `json:"is_primary"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
