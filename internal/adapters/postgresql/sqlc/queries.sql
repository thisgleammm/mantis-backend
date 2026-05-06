-- name: ListProducts :many
-- Untuk daftar produk, kita tidak menarik 'description' dan 'specifications' agar payload ringan.
SELECT 
    id, category_id, name, slug, base_price, discount_price, 
    rating_average, rating_count, created_at
FROM products
WHERE deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: FindProductByID :one
-- Untuk detail satu produk, baru kita tarik seluruh data beratnya.
SELECT 
    id, category_id, name, slug, description, base_price, 
    discount_price, weight, specifications, rating_average, 
    rating_count, created_at, updated_at
FROM products 
WHERE id = $1 AND deleted_at IS NULL;

-- name: FindProductBySlug :one
SELECT 
    id, category_id, name, slug, description, base_price, 
    discount_price, weight, specifications, rating_average, 
    rating_count, created_at, updated_at
FROM products 
WHERE slug = $1 AND deleted_at IS NULL;

-- name: ListUsers :many
-- Mengecualikan 'password' untuk keamanan.
SELECT 
    id, name, email, phone_number, created_at
FROM users
WHERE deleted_at IS NULL;

-- name: FindUserByID :one
-- Untuk detail user (misal untuk update profile), kita tidak menarik 'password' juga.
SELECT 
    id, name, email, phone_number, created_at
FROM users 
WHERE id = $1 AND deleted_at IS NULL;

-- name: CreateUser :one
INSERT INTO users (username, name, email, password, phone_number)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, username, name, email, phone_number, created_at;

-- name: FindUserByEmailForLogin :one
SELECT id, name, email, password
FROM users
WHERE email = $1 AND deleted_at IS NULL;

-- name: ListCategories :many
SELECT id, name, slug, description, icon_image_url, banner_image_url, is_active, created_at, updated_at
FROM categories
WHERE is_active = true
ORDER BY created_at DESC;

-- name: FindCategoryByID :one
SELECT id, name, slug, description, icon_image_url, banner_image_url, is_active, created_at, updated_at
FROM categories
WHERE id = $1 AND is_active = true;

-- name: CreateProduct :one
INSERT INTO products (
    category_id, name, slug, description, base_price, 
    discount_price, weight, specifications
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING id, category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count, created_at, updated_at;-- name: ListCarts :many
SELECT id, user_id, created_at, updated_at
FROM carts
ORDER BY created_at DESC;

-- name: ListProductImages :many
SELECT id, image_url, sort_order
FROM product_images
WHERE product_id = $1
ORDER BY sort_order ASC;

-- name: ListProductVariants :many
SELECT id, variant_name, price_extra, stock, stock_keeping_unit
FROM product_variants
WHERE product_id = $1 AND deleted_at IS NULL;