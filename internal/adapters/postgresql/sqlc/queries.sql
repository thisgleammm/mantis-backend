-- name: ListProducts :many
SELECT 
    p.id, p.category_id, p.name, p.slug, p.base_price, p.discount_price, 
    p.rating_average, p.rating_count, p.created_at,
    COALESCE(img.image_url, '')::TEXT as main_image
FROM products p
LEFT JOIN LATERAL (
    SELECT image_url FROM product_images WHERE product_id = p.id ORDER BY sort_order ASC LIMIT 1
) img ON true
WHERE p.deleted_at IS NULL 
  AND (p.created_at < $1 OR $1 IS NULL) -- $1 adalah waktu dari item terakhir di halaman sebelumnya
ORDER BY p.created_at DESC
LIMIT $2;

-- name: CountProducts :one
SELECT COUNT(*) FROM products 
WHERE deleted_at IS NULL
  AND (name ILIKE '%' || sqlc.arg('search_query')::text || '%' OR sqlc.arg('search_query')::text = '');

-- name: ListProductsOffset :many
SELECT 
    p.id, p.category_id, p.name, p.slug, p.base_price, p.discount_price, 
    p.rating_average, p.rating_count, p.created_at,
    COALESCE(img.image_url, '')::TEXT as main_image
FROM products p
LEFT JOIN LATERAL (
    SELECT image_url FROM product_images WHERE product_id = p.id ORDER BY sort_order ASC LIMIT 1
) img ON true
WHERE p.deleted_at IS NULL 
  AND (p.name ILIKE '%' || sqlc.arg('search_query')::text || '%' OR sqlc.arg('search_query')::text = '')
ORDER BY p.created_at DESC
LIMIT $1 OFFSET $2;

-- name: FindProductByID :one
-- Untuk detail satu produk, baru kita tarik seluruh data beratnya.
SELECT 
    id, category_id, name, slug, description, base_price, 
    discount_price, weight, specifications, rating_average, 
    rating_count, created_at, updated_at
FROM products 
WHERE id = $1 AND deleted_at IS NULL;

-- name: FindProductDetailBySlug :one
SELECT 
    p.id, p.category_id, p.name, p.slug, p.description, p.base_price, 
    p.discount_price, p.weight, p.specifications, p.rating_average, 
    p.rating_count, p.created_at,
    COALESCE(
        (SELECT jsonb_agg(
            jsonb_build_object('id', pi.id, 'image_url', pi.image_url, 'sort_order', pi.sort_order)
        ) FROM (
            SELECT id, image_url, sort_order 
            FROM product_images 
            WHERE product_id = p.id 
            ORDER BY sort_order ASC
        ) pi), 
    '[]'::jsonb) AS images,
    COALESCE(
        (SELECT jsonb_agg(
            jsonb_build_object('id', pv.id, 'variant_name', pv.variant_name, 'price_extra', pv.price_extra, 'stock', pv.stock)
        ) FROM product_variants pv 
        WHERE pv.product_id = p.id AND pv.deleted_at IS NULL), 
    '[]'::jsonb) AS variants
FROM products p
WHERE p.slug = $1 AND p.deleted_at IS NULL;

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
WHERE user_id = $1
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

-- name: AddItemToCart :one
INSERT INTO cart_items (cart_id, product_id, product_variant_id, quantity)
VALUES ($1, $2, $3, $4)
ON CONFLICT (cart_id, product_id, (COALESCE(product_variant_id, 0)))
DO UPDATE SET 
    quantity = cart_items.quantity + EXCLUDED.quantity,
    updated_at = NOW()
RETURNING *;

-- name: UpdateItemQuantity :one
UPDATE cart_items
SET quantity = $2, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: RemoveItemFromCart :exec
DELETE FROM cart_items
WHERE id = $1;

-- name: ListCartItems :many
SELECT 
    ci.id, ci.cart_id, ci.product_id, ci.product_variant_id, ci.quantity, ci.created_at, ci.updated_at,
    p.name as product_name, p.slug as product_slug, p.base_price as product_price,
    COALESCE(img.image_url, '')::TEXT as product_image,
    pv.variant_name, pv.price_extra as variant_price_extra
FROM cart_items ci
JOIN products p ON p.id = ci.product_id
LEFT JOIN LATERAL (
    SELECT image_url 
    FROM product_images 
    WHERE product_id = p.id 
    ORDER BY sort_order ASC 
    LIMIT 1
) img ON true
LEFT JOIN product_variants pv ON pv.id = ci.product_variant_id
WHERE ci.cart_id = $1
ORDER BY ci.created_at ASC;

-- name: CreateOrder :one
INSERT INTO orders (
    user_id, invoice_number, status, total_amount,
    shipping_cost, grand_total, shipping_address
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: CreateOrderItem :one
INSERT INTO order_items (
    order_id, product_id, product_variant_id, product_name, quantity, price_at_purchase
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: ListOrders :many
SELECT id, user_id, invoice_number, status, total_amount, shipping_cost, grand_total, shipping_address, tracking_number, courier_name, created_at, updated_at
FROM orders
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: ListOrderItems :many
SELECT id, order_id, product_id, product_variant_id, product_name, quantity, price_at_purchase, created_at, updated_at
FROM order_items
WHERE order_id = $1
ORDER BY created_at ASC;

-- name: ClearCartItems :exec
DELETE FROM cart_items
WHERE cart_id IN (SELECT id FROM carts WHERE user_id = $1);

-- name: UpsertPasswordReset :exec
INSERT INTO password_resets (email, token, expires_at)
VALUES ($1, $2, $3)
ON CONFLICT (email) DO UPDATE SET
    token = EXCLUDED.token,
    expires_at = EXCLUDED.expires_at,
    created_at = CURRENT_TIMESTAMP;

-- name: FindPasswordResetByToken :one
SELECT email, token, expires_at
FROM password_resets
WHERE token = $1;

-- name: DeletePasswordReset :exec
DELETE FROM password_resets
WHERE email = $1;

-- name: UpdateUserPassword :exec
UPDATE users
SET password = $2
WHERE email = $1;