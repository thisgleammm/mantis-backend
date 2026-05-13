-- =============================================================
-- MANTIS MARKETPLACE — SEED DATA
-- =============================================================

-- Truncate existing data to ensure clean state and reset IDs (optional but safer for seeding)
-- TRUNCATE categories, products, product_variants, product_images RESTART IDENTITY CASCADE;

-- -------------------------------------------------------------
-- 1. CATEGORIES
-- -------------------------------------------------------------
INSERT INTO categories (name, slug, description, is_active)
VALUES
  ('Pakaian Pria',   'pakaian-pria',   'Koleksi pakaian pria terkini', true),
  ('Pakaian Wanita', 'pakaian-wanita', 'Koleksi pakaian wanita terkini', true),
  ('Elektronik',     'elektronik',     'Gadget & perangkat elektronik', true),
  ('Olahraga',       'olahraga',       'Perlengkapan & pakaian olahraga', true),
  ('Sepatu',         'sepatu',         'Sepatu pria, wanita, dan anak', true),
  ('Tas & Dompet',   'tas-dompet',     'Tas fashion dan dompet premium', true)
ON CONFLICT (slug) DO UPDATE SET name = EXCLUDED.name;

-- -------------------------------------------------------------
-- 2. PRODUCTS
-- -------------------------------------------------------------
INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES
  ((SELECT id FROM categories WHERE slug = 'pakaian-pria'), 'Kemeja Linen Premium', 'kemeja-linen-premium', 'Kemeja bahan linen berkualitas tinggi.', 299000, 249000, 300, '{"material":"Linen"}', 4.7, 128),
  ((SELECT id FROM categories WHERE slug = 'pakaian-pria'), 'Kaos Oversize Graphic', 'kaos-oversize-graphic', 'Kaos oversize desain grafis.', 189000, NULL, 250, '{"material":"Cotton"}', 4.5, 89),
  ((SELECT id FROM categories WHERE slug = 'pakaian-wanita'), 'Dress Midi Floral', 'dress-midi-floral', 'Dress midi motif bunga.', 450000, 385000, 350, '{"material":"Viscose"}', 4.8, 214),
  ((SELECT id FROM categories WHERE slug = 'pakaian-wanita'), 'Blouse Satin Polos', 'blouse-satin-polos', 'Blouse satin minimalis.', 275000, NULL, 200, '{"material":"Satin"}', 4.6, 76),
  ((SELECT id FROM categories WHERE slug = 'elektronik'), 'Earbuds TWS Pro Max', 'earbuds-tws-pro-max', 'True wireless earbuds ANC.', 899000, 749000, 180, '{"connectivity":"BT 5.3"}', 4.9, 502),
  ((SELECT id FROM categories WHERE slug = 'elektronik'), 'Smartwatch Fitness Ultra', 'smartwatch-fitness-ultra', 'Smartwatch fitness monitoring.', 1299000, NULL, 280, '{"display":"AMOLED"}', 4.7, 337),
  ((SELECT id FROM categories WHERE slug = 'olahraga'), 'Jersey Olahraga Dry-Fit', 'jersey-olahraga-dry-fit', 'Jersey olahraga dry-fit.', 220000, 185000, 200, '{"material":"Polyester"}', 4.5, 193),
  ((SELECT id FROM categories WHERE slug = 'sepatu'), 'Sneakers Casual Low-Cut', 'sneakers-casual-low-cut', 'Sneakers casual sol EVA.', 525000, 450000, 600, '{"upper":"Canvas"}', 4.6, 418),
  ((SELECT id FROM categories WHERE slug = 'tas-dompet'), 'Tote Bag Canvas Premium', 'tote-bag-canvas-premium', 'Tote bag canvas tebal.', 350000, NULL, 450, '{"material":"Canvas"}', 4.7, 162)
ON CONFLICT (slug) DO UPDATE SET 
  category_id = EXCLUDED.category_id,
  name = EXCLUDED.name,
  base_price = EXCLUDED.base_price;

-- -------------------------------------------------------------
-- 3. PRODUCT VARIANTS
-- -------------------------------------------------------------

-- Kemeja Linen Premium
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit)
SELECT id, 'S / Putih', 0, 50, 'KLP-S-WHT' FROM products WHERE slug = 'kemeja-linen-premium'
ON CONFLICT (stock_keeping_unit) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit)
SELECT id, 'M / Putih', 0, 80, 'KLP-M-WHT' FROM products WHERE slug = 'kemeja-linen-premium'
ON CONFLICT (stock_keeping_unit) DO NOTHING;

-- Kaos Oversize Graphic
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit)
SELECT id, 'S / Hitam', 0, 100, 'KOG-S-BLK' FROM products WHERE slug = 'kaos-oversize-graphic'
ON CONFLICT (stock_keeping_unit) DO NOTHING;

-- Earbuds TWS Pro Max
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit)
SELECT id, 'Hitam', 0, 200, 'ETPM-BLK' FROM products WHERE slug = 'earbuds-tws-pro-max'
ON CONFLICT (stock_keeping_unit) DO NOTHING;

-- -------------------------------------------------------------
-- 4. PRODUCT IMAGES
-- -------------------------------------------------------------

INSERT INTO product_images (product_id, image_url, sort_order)
SELECT id, 'https://images.unsplash.com/photo-1596755094514-f87e34085b2c?w=800', 0 FROM products WHERE slug = 'kemeja-linen-premium'
ON CONFLICT (product_id, sort_order) DO NOTHING;

INSERT INTO product_images (product_id, image_url, sort_order)
SELECT id, 'https://images.unsplash.com/photo-1521572163474-6864f9cf17ab?w=800', 0 FROM products WHERE slug = 'kaos-oversize-graphic'
ON CONFLICT (product_id, sort_order) DO NOTHING;

INSERT INTO product_images (product_id, image_url, sort_order)
SELECT id, 'https://images.unsplash.com/photo-1590658268037-6bf12165a8df?w=800', 0 FROM products WHERE slug = 'earbuds-tws-pro-max'
ON CONFLICT (product_id, sort_order) DO NOTHING;
