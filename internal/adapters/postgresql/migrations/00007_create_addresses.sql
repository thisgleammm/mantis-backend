-- +goose Up
-- +goose StatementBegin
CREATE TABLE addresses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    recipient_name TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    province TEXT NOT NULL,
    city TEXT NOT NULL,
    district TEXT NOT NULL,
    postal_code TEXT NOT NULL,
    full_address TEXT NOT NULL,
    label TEXT, -- e.g., 'Rumah', 'Kantor'
    coordinates TEXT, -- Bisa menyimpan string lat,long atau JSON
    is_primary BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_addresses_user_id ON addresses(user_id);
-- Menjamin hanya ada satu alamat utama per user
CREATE UNIQUE INDEX idx_addresses_user_primary ON addresses(user_id) WHERE is_primary = true;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS addresses;
-- +goose StatementEnd
