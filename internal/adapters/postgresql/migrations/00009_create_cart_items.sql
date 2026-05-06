-- +goose Up
CREATE TABLE IF NOT EXISTS cart_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    cart_id UUID NOT NULL REFERENCES carts(id) ON DELETE CASCADE,
    product_id BIGINT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    product_variant_id BIGINT REFERENCES product_variants(id) ON DELETE SET NULL,
    quantity INTEGER NOT NULL DEFAULT 1,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Index for performance
CREATE INDEX IF NOT EXISTS idx_cart_items_cart_id ON cart_items(cart_id);

-- Ensure unique product variant per cart
CREATE UNIQUE INDEX IF NOT EXISTS idx_cart_items_unique ON cart_items (cart_id, product_id, COALESCE(product_variant_id, 0));

-- +goose Down
DROP TABLE IF EXISTS cart_items;
