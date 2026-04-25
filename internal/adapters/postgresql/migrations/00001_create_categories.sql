-- +goose Up
CREATE TABLE IF NOT EXISTS categories (
    id          BIGSERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    slug        VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    icon_image_url TEXT,
    banner_image_url TEXT,
    is_active   BOOLEAN NOT NULL DEFAULT true,
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_categories_created_at ON categories(created_at);
CREATE INDEX IF NOT EXISTS idx_categories_slug ON categories(slug);

-- +goose Down
DROP INDEX IF EXISTS idx_categories_created_at;
DROP INDEX IF EXISTS idx_categories_slug;
DROP TABLE IF EXISTS categories;
