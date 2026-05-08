package domain

import "time"

type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusShipped    OrderStatus = "shipped"
	OrderStatusDelivered  OrderStatus = "delivered"
	OrderStatusCancelled  OrderStatus = "cancelled"
)

type Order struct {
	ID              string      `json:"id"`
	UserID          string      `json:"user_id"`
	InvoiceNumber   string      `json:"invoice_number"`
	Status          OrderStatus `json:"status"`
	TotalAmount     float64     `json:"total_amount"`
	ShippingCost    float64     `json:"shipping_cost"`
	GrandTotal      float64     `json:"grand_total"`
	ShippingAddress string      `json:"shipping_address"`
	TrackingNumber  string      `json:"tracking_number"`
	CourierName     string      `json:"courier_name"`
	Items           []OrderItem `json:"items,omitempty"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}

type OrderItem struct {
	ID               string    `json:"id"`
	OrderID          string    `json:"order_id"`
	ProductID        int64     `json:"product_id"`
	ProductVariantID *int64    `json:"product_variant_id,omitempty"`
	ProductName      string    `json:"product_name"`
	Quantity         int32     `json:"quantity"`
	PriceAtPurchase  float64   `json:"price_at_purchase"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
