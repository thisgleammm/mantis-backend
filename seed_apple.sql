-- =============================================================
-- MANTIS MARKETPLACE — APPLE SEED DATA
-- =============================================================

-- -------------------------------------------------------------
-- 1. CATEGORIES
-- -------------------------------------------------------------
INSERT INTO categories (name, slug, description, is_active)
VALUES
  ('Gadgets', 'gadgets', 'Premium smartphones, laptops, and tablets', true)
ON CONFLICT (slug) DO UPDATE SET name = EXCLUDED.name;

-- -------------------------------------------------------------
-- 2. PRODUCTS (APPLE)
-- -------------------------------------------------------------
INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES
  ((SELECT id FROM categories WHERE slug = 'gadgets'), 
   'iPhone 15 Pro', 'iphone-15-pro', 
   'Titanium design, A17 Pro chip, customizable Action button, and a more versatile Pro camera system.', 
   18999000, 18499000, 187, 
   '{"chip":"A17 Pro","display":"6.1 Super Retina XDR","camera":"48MP Main","material":"Titanium"}', 
   4.9, 1250),

  ((SELECT id FROM categories WHERE slug = 'gadgets'), 
   'MacBook Air M2', 'macbook-air-m2', 
   'Strikingly thin design, 13.6-inch Liquid Retina display, and the blazing-fast M2 chip.', 
   17499000, 16999000, 1240, 
   '{"chip":"Apple M2","display":"13.6 Liquid Retina","battery":"Up to 18 hours","weight":"1.24kg"}', 
   4.8, 850),

  ((SELECT id FROM categories WHERE slug = 'gadgets'), 
   'iPad Pro M2', 'ipad-pro-m2', 
   'The ultimate iPad experience. Now with next-level M2 performance and superfast wireless.', 
   15499000, NULL, 466, 
   '{"chip":"Apple M2","display":"11 Liquid Retina","connectivity":"Wi-Fi 6E","pencil":"Apple Pencil 2nd Gen"}', 
   4.7, 420)
ON CONFLICT (slug) DO UPDATE SET 
  category_id = EXCLUDED.category_id,
  name = EXCLUDED.name,
  base_price = EXCLUDED.base_price;

-- -------------------------------------------------------------
-- 3. PRODUCT VARIANTS
-- -------------------------------------------------------------

-- iPhone 15 Pro
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit)
SELECT id, '128GB / Natural Titanium', 0, 50, 'IP15P-128-NAT' FROM products WHERE slug = 'iphone-15-pro'
ON CONFLICT (stock_keeping_unit) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit)
SELECT id, '256GB / Natural Titanium', 2000000, 40, 'IP15P-256-NAT' FROM products WHERE slug = 'iphone-15-pro'
ON CONFLICT (stock_keeping_unit) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit)
SELECT id, '128GB / Blue Titanium', 0, 35, 'IP15P-128-BLU' FROM products WHERE slug = 'iphone-15-pro'
ON CONFLICT (stock_keeping_unit) DO NOTHING;

-- MacBook Air M2
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit)
SELECT id, '8GB RAM / 256GB SSD / Midnight', 0, 30, 'MBA-M2-8-256-MID' FROM products WHERE slug = 'macbook-air-m2'
ON CONFLICT (stock_keeping_unit) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit)
SELECT id, '16GB RAM / 512GB SSD / Midnight', 4500000, 20, 'MBA-M2-16-512-MID' FROM products WHERE slug = 'macbook-air-m2'
ON CONFLICT (stock_keeping_unit) DO NOTHING;

-- iPad Pro M2
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit)
SELECT id, '128GB / Space Gray', 0, 25, 'IPPRO-M2-128-SG' FROM products WHERE slug = 'ipad-pro-m2'
ON CONFLICT (stock_keeping_unit) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit)
SELECT id, '256GB / Space Gray', 1800000, 15, 'IPPRO-M2-256-SG' FROM products WHERE slug = 'ipad-pro-m2'
ON CONFLICT (stock_keeping_unit) DO NOTHING;

-- -------------------------------------------------------------
-- 4. PRODUCT IMAGES
-- -------------------------------------------------------------

-- iPhone 15 Pro
INSERT INTO product_images (product_id, image_url, sort_order)
SELECT id, 'https://images.unsplash.com/photo-1696446701796-da61225697cc?w=800', 0 FROM products WHERE slug = 'iphone-15-pro'
ON CONFLICT (product_id, sort_order) DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order)
SELECT id, 'https://images.unsplash.com/photo-1696446702183-cbd13d78e1e7?w=800', 1 FROM products WHERE slug = 'iphone-15-pro'
ON CONFLICT (product_id, sort_order) DO NOTHING;

-- MacBook Air M2
INSERT INTO product_images (product_id, image_url, sort_order)
SELECT id, 'https://images.unsplash.com/photo-1611186871348-b1ec696e523b?w=800', 0 FROM products WHERE slug = 'macbook-air-m2'
ON CONFLICT (product_id, sort_order) DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order)
SELECT id, 'https://images.unsplash.com/photo-1517336714731-489689fd1ca8?w=800', 1 FROM products WHERE slug = 'macbook-air-m2'
ON CONFLICT (product_id, sort_order) DO NOTHING;

-- iPad Pro M2
INSERT INTO product_images (product_id, image_url, sort_order)
SELECT id, 'https://images.unsplash.com/photo-1544244015-0df4b3ffc6b0?w=800', 0 FROM products WHERE slug = 'ipad-pro-m2'
ON CONFLICT (product_id, sort_order) DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order)
SELECT id, 'https://images.unsplash.com/photo-1585503418537-88331351ad99?w=800', 1 FROM products WHERE slug = 'ipad-pro-m2'
ON CONFLICT (product_id, sort_order) DO NOTHING;
