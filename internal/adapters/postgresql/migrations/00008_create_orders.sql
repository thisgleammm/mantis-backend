-- +goose Up
CREATE TYPE order_status AS ENUM ('pending', 'processing', 'shipped', 'delivered', 'cancelled');

CREATE TABLE IF NOT EXISTS orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    invoice_number VARCHAR(255) NOT NULL UNIQUE,
    status order_status NOT NULL DEFAULT 'pending',
    total_amount NUMERIC(15, 2) NOT NULL DEFAULT 0,
    shipping_cost NUMERIC(15, 2) NOT NULL DEFAULT 0,
    grand_total NUMERIC(15, 2) NOT NULL DEFAULT 0,
    shipping_address TEXT NOT NULL,
    tracking_number VARCHAR(255),
    courier_name VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_orders_user_id ON orders(user_id);
CREATE INDEX IF NOT EXISTS idx_orders_invoice_number ON orders(invoice_number);
CREATE INDEX IF NOT EXISTS idx_orders_status ON orders(status);

-- +goose Down
DROP TABLE IF EXISTS orders;
DROP TYPE IF EXISTS order_status;
