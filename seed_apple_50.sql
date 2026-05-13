-- Apple Seed Data (50 Products)
INSERT INTO categories (name, slug, description, is_active) VALUES ('Apple', 'apple', 'Complete Apple Ecosystem', true) ON CONFLICT (slug) DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'iPhone 15 Pro Max', 'iphone-15-pro-max', 'Titanium, A17 Pro, 5x Telephoto.', 17000000, 15300000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'IPHONE-15-PRO-MAX-0' FROM products WHERE slug = 'iphone-15-pro-max' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'IPHONE-15-PRO-MAX-1' FROM products WHERE slug = 'iphone-15-pro-max' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1510557880182-3d4d3cba35a5?w=800', 0 FROM products WHERE slug = 'iphone-15-pro-max' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1542744094-24638eff58bb?w=800', 1 FROM products WHERE slug = 'iphone-15-pro-max' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'iPhone 15 Pro', 'iphone-15-pro', 'Titanium, A17 Pro, Pro camera.', 37000000, 33300000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'IPHONE-15-PRO-0' FROM products WHERE slug = 'iphone-15-pro' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'IPHONE-15-PRO-1' FROM products WHERE slug = 'iphone-15-pro' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1542744094-24638eff58bb?w=800', 0 FROM products WHERE slug = 'iphone-15-pro' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1544244015-0df4b3ffc6b0?w=800', 1 FROM products WHERE slug = 'iphone-15-pro' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'iPhone 15 Plus', 'iphone-15-plus', 'A16 Bionic, 48MP camera, big battery.', 12000000, 10800000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'IPHONE-15-PLUS-0' FROM products WHERE slug = 'iphone-15-plus' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'IPHONE-15-PLUS-1' FROM products WHERE slug = 'iphone-15-plus' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1511707171634-5f897ff02aa9?w=800', 0 FROM products WHERE slug = 'iphone-15-plus' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1517336714731-489689fd1ca8?w=800', 1 FROM products WHERE slug = 'iphone-15-plus' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'iPhone 15', 'iphone-15', 'A16 Bionic, 48MP camera, Dynamic Island.', 13000000, 11700000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'IPHONE-15-0' FROM products WHERE slug = 'iphone-15' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'IPHONE-15-1' FROM products WHERE slug = 'iphone-15' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1611186871348-b1ec696e523b?w=800', 0 FROM products WHERE slug = 'iphone-15' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1517336714731-489689fd1ca8?w=800', 1 FROM products WHERE slug = 'iphone-15' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'iPhone 14', 'iphone-14', 'A15 Bionic, dual-camera system.', 35000000, 31500000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'IPHONE-14-0' FROM products WHERE slug = 'iphone-14' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'IPHONE-14-1' FROM products WHERE slug = 'iphone-14' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1542744094-24638eff58bb?w=800', 0 FROM products WHERE slug = 'iphone-14' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1544244015-0df4b3ffc6b0?w=800', 1 FROM products WHERE slug = 'iphone-14' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'iPhone SE', 'iphone-se-3', 'A15 Bionic, compact design.', 44000000, 39600000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'IPHONE-SE-3-0' FROM products WHERE slug = 'iphone-se-3' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'IPHONE-SE-3-1' FROM products WHERE slug = 'iphone-se-3' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1511707171634-5f897ff02aa9?w=800', 0 FROM products WHERE slug = 'iphone-se-3' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1542744094-24638eff58bb?w=800', 1 FROM products WHERE slug = 'iphone-se-3' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'MacBook Pro 14 M3', 'macbook-pro-14-m3', 'M3 chip, Liquid Retina XDR.', 26000000, 23400000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'MACBOOK-PRO-14-M3-0' FROM products WHERE slug = 'macbook-pro-14-m3' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'MACBOOK-PRO-14-M3-1' FROM products WHERE slug = 'macbook-pro-14-m3' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1591337676887-a217a6970a8a?w=800', 0 FROM products WHERE slug = 'macbook-pro-14-m3' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1556656793-062ff9878258?w=800', 1 FROM products WHERE slug = 'macbook-pro-14-m3' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'MacBook Pro 16 M3 Max', 'macbook-pro-16-m3-max', 'M3 Max chip, extreme performance.', 34000000, 30600000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'MACBOOK-PRO-16-M3-MAX-0' FROM products WHERE slug = 'macbook-pro-16-m3-max' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'MACBOOK-PRO-16-M3-MAX-1' FROM products WHERE slug = 'macbook-pro-16-m3-max' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1542744094-24638eff58bb?w=800', 0 FROM products WHERE slug = 'macbook-pro-16-m3-max' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1611186871348-b1ec696e523b?w=800', 1 FROM products WHERE slug = 'macbook-pro-16-m3-max' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'MacBook Air 13 M2', 'macbook-air-13-m2', 'M2 chip, thin and light.', 12000000, 10800000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'MACBOOK-AIR-13-M2-0' FROM products WHERE slug = 'macbook-air-13-m2' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'MACBOOK-AIR-13-M2-1' FROM products WHERE slug = 'macbook-air-13-m2' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1517336714731-489689fd1ca8?w=800', 0 FROM products WHERE slug = 'macbook-air-13-m2' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1510557880182-3d4d3cba35a5?w=800', 1 FROM products WHERE slug = 'macbook-air-13-m2' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'MacBook Air 15 M2', 'macbook-air-15-m2', 'M2 chip, big display, thin.', 28000000, 25200000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'MACBOOK-AIR-15-M2-0' FROM products WHERE slug = 'macbook-air-15-m2' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'MACBOOK-AIR-15-M2-1' FROM products WHERE slug = 'macbook-air-15-m2' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1591337676887-a217a6970a8a?w=800', 0 FROM products WHERE slug = 'macbook-air-15-m2' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1544244015-0df4b3ffc6b0?w=800', 1 FROM products WHERE slug = 'macbook-air-15-m2' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'iPad Pro 12.9 M2', 'ipad-pro-12-9-m2', 'M2, Liquid Retina XDR, ProMotion.', 24000000, 21600000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'IPAD-PRO-12-9-M2-0' FROM products WHERE slug = 'ipad-pro-12-9-m2' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'IPAD-PRO-12-9-M2-1' FROM products WHERE slug = 'ipad-pro-12-9-m2' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1585503418537-88331351ad99?w=800', 0 FROM products WHERE slug = 'ipad-pro-12-9-m2' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1591337676887-a217a6970a8a?w=800', 1 FROM products WHERE slug = 'ipad-pro-12-9-m2' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'iPad Pro 11 M2', 'ipad-pro-11-m2', 'M2, Liquid Retina, ProMotion.', 49000000, 44100000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'IPAD-PRO-11-M2-0' FROM products WHERE slug = 'ipad-pro-11-m2' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'IPAD-PRO-11-M2-1' FROM products WHERE slug = 'ipad-pro-11-m2' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1591337676887-a217a6970a8a?w=800', 0 FROM products WHERE slug = 'ipad-pro-11-m2' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1585503418537-88331351ad99?w=800', 1 FROM products WHERE slug = 'ipad-pro-11-m2' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'iPad Air M1', 'ipad-air-m1', 'M1 chip, powerful and colorful.', 21000000, 18900000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'IPAD-AIR-M1-0' FROM products WHERE slug = 'ipad-air-m1' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'IPAD-AIR-M1-1' FROM products WHERE slug = 'ipad-air-m1' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1611186871348-b1ec696e523b?w=800', 0 FROM products WHERE slug = 'ipad-air-m1' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1542744094-24638eff58bb?w=800', 1 FROM products WHERE slug = 'ipad-air-m1' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'iPad 10th Gen', 'ipad-10-gen', 'All-screen design, A14 Bionic.', 22000000, 19800000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'IPAD-10-GEN-0' FROM products WHERE slug = 'ipad-10-gen' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'IPAD-10-GEN-1' FROM products WHERE slug = 'ipad-10-gen' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1544244015-0df4b3ffc6b0?w=800', 0 FROM products WHERE slug = 'ipad-10-gen' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1591337676887-a217a6970a8a?w=800', 1 FROM products WHERE slug = 'ipad-10-gen' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'iPad Mini', 'ipad-mini-6', 'A15 Bionic, portable power.', 21000000, 18900000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'IPAD-MINI-6-0' FROM products WHERE slug = 'ipad-mini-6' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'IPAD-MINI-6-1' FROM products WHERE slug = 'ipad-mini-6' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1517336714731-489689fd1ca8?w=800', 0 FROM products WHERE slug = 'ipad-mini-6' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1591337676887-a217a6970a8a?w=800', 1 FROM products WHERE slug = 'ipad-mini-6' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'Apple Watch Ultra 2', 'apple-watch-ultra-2', 'Rugged, 3000 nits, S9 chip.', 37000000, 33300000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'APPLE-WATCH-ULTRA-2-0' FROM products WHERE slug = 'apple-watch-ultra-2' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'APPLE-WATCH-ULTRA-2-1' FROM products WHERE slug = 'apple-watch-ultra-2' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1585503418537-88331351ad99?w=800', 0 FROM products WHERE slug = 'apple-watch-ultra-2' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1556656793-062ff9878258?w=800', 1 FROM products WHERE slug = 'apple-watch-ultra-2' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'Apple Watch Series 9', 'apple-watch-series-9', 'S9 chip, Double Tap gesture.', 40000000, 36000000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'APPLE-WATCH-SERIES-9-0' FROM products WHERE slug = 'apple-watch-series-9' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'APPLE-WATCH-SERIES-9-1' FROM products WHERE slug = 'apple-watch-series-9' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1616348436168-de43ad0db179?w=800', 0 FROM products WHERE slug = 'apple-watch-series-9' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1591337676887-a217a6970a8a?w=800', 1 FROM products WHERE slug = 'apple-watch-series-9' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'Apple Watch SE', 'apple-watch-se-2', 'Essential features, affordable.', 32000000, 28800000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'APPLE-WATCH-SE-2-0' FROM products WHERE slug = 'apple-watch-se-2' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'APPLE-WATCH-SE-2-1' FROM products WHERE slug = 'apple-watch-se-2' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1616348436168-de43ad0db179?w=800', 0 FROM products WHERE slug = 'apple-watch-se-2' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1511707171634-5f897ff02aa9?w=800', 1 FROM products WHERE slug = 'apple-watch-se-2' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'AirPods Pro 2', 'airpods-pro-2-usb-c', 'ANC, Transparency, USB-C.', 47000000, 42300000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'AIRPODS-PRO-2-USB-C-0' FROM products WHERE slug = 'airpods-pro-2-usb-c' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'AIRPODS-PRO-2-USB-C-1' FROM products WHERE slug = 'airpods-pro-2-usb-c' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1585503418537-88331351ad99?w=800', 0 FROM products WHERE slug = 'airpods-pro-2-usb-c' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1544244015-0df4b3ffc6b0?w=800', 1 FROM products WHERE slug = 'airpods-pro-2-usb-c' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'AirPods 3', 'airpods-3', 'Spatial audio, sweat resistant.', 35000000, 31500000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'AIRPODS-3-0' FROM products WHERE slug = 'airpods-3' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'AIRPODS-3-1' FROM products WHERE slug = 'airpods-3' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1616348436168-de43ad0db179?w=800', 0 FROM products WHERE slug = 'airpods-3' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1542744094-24638eff58bb?w=800', 1 FROM products WHERE slug = 'airpods-3' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'AirPods Max', 'airpods-max', 'High-fidelity audio, ANC.', 46000000, 41400000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'AIRPODS-MAX-0' FROM products WHERE slug = 'airpods-max' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'AIRPODS-MAX-1' FROM products WHERE slug = 'airpods-max' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1556656793-062ff9878258?w=800', 0 FROM products WHERE slug = 'airpods-max' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1591337676887-a217a6970a8a?w=800', 1 FROM products WHERE slug = 'airpods-max' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'iMac 24 M3', 'imac-24-m3', 'Colorful, thin, M3 chip.', 38000000, 34200000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'IMAC-24-M3-0' FROM products WHERE slug = 'imac-24-m3' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'IMAC-24-M3-1' FROM products WHERE slug = 'imac-24-m3' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1591337676887-a217a6970a8a?w=800', 0 FROM products WHERE slug = 'imac-24-m3' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1510557880182-3d4d3cba35a5?w=800', 1 FROM products WHERE slug = 'imac-24-m3' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'Mac Mini M2', 'mac-mini-m2', 'M2 power in a compact box.', 44000000, 39600000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'MAC-MINI-M2-0' FROM products WHERE slug = 'mac-mini-m2' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'MAC-MINI-M2-1' FROM products WHERE slug = 'mac-mini-m2' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1517336714731-489689fd1ca8?w=800', 0 FROM products WHERE slug = 'mac-mini-m2' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1511707171634-5f897ff02aa9?w=800', 1 FROM products WHERE slug = 'mac-mini-m2' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'Mac Studio M2 Ultra', 'mac-studio-m2-ultra', 'Extreme performance for pros.', 45000000, 40500000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'MAC-STUDIO-M2-ULTRA-0' FROM products WHERE slug = 'mac-studio-m2-ultra' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'MAC-STUDIO-M2-ULTRA-1' FROM products WHERE slug = 'mac-studio-m2-ultra' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1611186871348-b1ec696e523b?w=800', 0 FROM products WHERE slug = 'mac-studio-m2-ultra' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1585503418537-88331351ad99?w=800', 1 FROM products WHERE slug = 'mac-studio-m2-ultra' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'Apple TV 4K', 'apple-tv-4k-3', 'A15 Bionic, HDR10+, Siri Remote.', 17000000, 15300000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'APPLE-TV-4K-3-0' FROM products WHERE slug = 'apple-tv-4k-3' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'APPLE-TV-4K-3-1' FROM products WHERE slug = 'apple-tv-4k-3' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1544244015-0df4b3ffc6b0?w=800', 0 FROM products WHERE slug = 'apple-tv-4k-3' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1616348436168-de43ad0db179?w=800', 1 FROM products WHERE slug = 'apple-tv-4k-3' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'Studio Display', 'studio-display', '5K Retina, 12MP Center Stage.', 44000000, 39600000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'STUDIO-DISPLAY-0' FROM products WHERE slug = 'studio-display' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'STUDIO-DISPLAY-1' FROM products WHERE slug = 'studio-display' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1585503418537-88331351ad99?w=800', 0 FROM products WHERE slug = 'studio-display' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1542744094-24638eff58bb?w=800', 1 FROM products WHERE slug = 'studio-display' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'Pro Display XDR', 'pro-display-xdr', '32-inch 6K, extreme brightness.', 32000000, 28800000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'PRO-DISPLAY-XDR-0' FROM products WHERE slug = 'pro-display-xdr' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'PRO-DISPLAY-XDR-1' FROM products WHERE slug = 'pro-display-xdr' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1510557880182-3d4d3cba35a5?w=800', 0 FROM products WHERE slug = 'pro-display-xdr' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1542744094-24638eff58bb?w=800', 1 FROM products WHERE slug = 'pro-display-xdr' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'HomePod 2', 'homepod-2', 'Immersive sound, smart home hub.', 37000000, 33300000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'HOMEPOD-2-0' FROM products WHERE slug = 'homepod-2' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'HOMEPOD-2-1' FROM products WHERE slug = 'homepod-2' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1616348436168-de43ad0db179?w=800', 0 FROM products WHERE slug = 'homepod-2' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1544244015-0df4b3ffc6b0?w=800', 1 FROM products WHERE slug = 'homepod-2' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'HomePod Mini', 'homepod-mini', 'Room-filling sound, Siri.', 15000000, 13500000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'HOMEPOD-MINI-0' FROM products WHERE slug = 'homepod-mini' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'HOMEPOD-MINI-1' FROM products WHERE slug = 'homepod-mini' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1591337676887-a217a6970a8a?w=800', 0 FROM products WHERE slug = 'homepod-mini' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1511707171634-5f897ff02aa9?w=800', 1 FROM products WHERE slug = 'homepod-mini' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'AirTag', 'airtag', 'Find your keys, wallet, everything.', 18000000, 16200000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'AIRTAG-0' FROM products WHERE slug = 'airtag' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'AIRTAG-1' FROM products WHERE slug = 'airtag' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1517336714731-489689fd1ca8?w=800', 0 FROM products WHERE slug = 'airtag' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1556656793-062ff9878258?w=800', 1 FROM products WHERE slug = 'airtag' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'Apple Watch Ultra 2 (Refurbished 2024)', 'apple-watch-ultra-2-refurb-6863', 'Certified Refurbished: Rugged, 3000 nits, S9 chip.', 20000000, 18000000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'APPLE-WATCH-ULTRA-2-REFURB-6863-0' FROM products WHERE slug = 'apple-watch-ultra-2-refurb-6863' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'APPLE-WATCH-ULTRA-2-REFURB-6863-1' FROM products WHERE slug = 'apple-watch-ultra-2-refurb-6863' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1542744094-24638eff58bb?w=800', 0 FROM products WHERE slug = 'apple-watch-ultra-2-refurb-6863' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1510557880182-3d4d3cba35a5?w=800', 1 FROM products WHERE slug = 'apple-watch-ultra-2-refurb-6863' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'iPad Pro 12.9 M2 (Refurbished 2024)', 'ipad-pro-12-9-m2-refurb-6273', 'Certified Refurbished: M2, Liquid Retina XDR, ProMotion.', 39000000, 35100000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'IPAD-PRO-12-9-M2-REFURB-6273-0' FROM products WHERE slug = 'ipad-pro-12-9-m2-refurb-6273' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'IPAD-PRO-12-9-M2-REFURB-6273-1' FROM products WHERE slug = 'ipad-pro-12-9-m2-refurb-6273' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1556656793-062ff9878258?w=800', 0 FROM products WHERE slug = 'ipad-pro-12-9-m2-refurb-6273' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1556656793-062ff9878258?w=800', 1 FROM products WHERE slug = 'ipad-pro-12-9-m2-refurb-6273' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'iPad Mini (Refurbished 2024)', 'ipad-mini-6-refurb-1084', 'Certified Refurbished: A15 Bionic, portable power.', 36000000, 32400000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'IPAD-MINI-6-REFURB-1084-0' FROM products WHERE slug = 'ipad-mini-6-refurb-1084' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'IPAD-MINI-6-REFURB-1084-1' FROM products WHERE slug = 'ipad-mini-6-refurb-1084' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1511707171634-5f897ff02aa9?w=800', 0 FROM products WHERE slug = 'ipad-mini-6-refurb-1084' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1544244015-0df4b3ffc6b0?w=800', 1 FROM products WHERE slug = 'ipad-mini-6-refurb-1084' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'Apple Watch SE (Refurbished 2024)', 'apple-watch-se-2-refurb-2982', 'Certified Refurbished: Essential features, affordable.', 15000000, 13500000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'APPLE-WATCH-SE-2-REFURB-2982-0' FROM products WHERE slug = 'apple-watch-se-2-refurb-2982' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'APPLE-WATCH-SE-2-REFURB-2982-1' FROM products WHERE slug = 'apple-watch-se-2-refurb-2982' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1616348436168-de43ad0db179?w=800', 0 FROM products WHERE slug = 'apple-watch-se-2-refurb-2982' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1585503418537-88331351ad99?w=800', 1 FROM products WHERE slug = 'apple-watch-se-2-refurb-2982' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'AirPods Pro 2 (Refurbished 2024)', 'airpods-pro-2-usb-c-refurb-9695', 'Certified Refurbished: ANC, Transparency, USB-C.', 48000000, 43200000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'AIRPODS-PRO-2-USB-C-REFURB-9695-0' FROM products WHERE slug = 'airpods-pro-2-usb-c-refurb-9695' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'AIRPODS-PRO-2-USB-C-REFURB-9695-1' FROM products WHERE slug = 'airpods-pro-2-usb-c-refurb-9695' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1591337676887-a217a6970a8a?w=800', 0 FROM products WHERE slug = 'airpods-pro-2-usb-c-refurb-9695' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1556656793-062ff9878258?w=800', 1 FROM products WHERE slug = 'airpods-pro-2-usb-c-refurb-9695' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'iPhone 14 (Refurbished 2024)', 'iphone-14-refurb-3440', 'Certified Refurbished: A15 Bionic, dual-camera system.', 48000000, 43200000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'IPHONE-14-REFURB-3440-0' FROM products WHERE slug = 'iphone-14-refurb-3440' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'IPHONE-14-REFURB-3440-1' FROM products WHERE slug = 'iphone-14-refurb-3440' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1511707171634-5f897ff02aa9?w=800', 0 FROM products WHERE slug = 'iphone-14-refurb-3440' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1611186871348-b1ec696e523b?w=800', 1 FROM products WHERE slug = 'iphone-14-refurb-3440' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'iPad Air M1 (Refurbished 2023)', 'ipad-air-m1-refurb-1394', 'Certified Refurbished: M1 chip, powerful and colorful.', 48000000, 43200000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'IPAD-AIR-M1-REFURB-1394-0' FROM products WHERE slug = 'ipad-air-m1-refurb-1394' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'IPAD-AIR-M1-REFURB-1394-1' FROM products WHERE slug = 'ipad-air-m1-refurb-1394' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1616348436168-de43ad0db179?w=800', 0 FROM products WHERE slug = 'ipad-air-m1-refurb-1394' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1517336714731-489689fd1ca8?w=800', 1 FROM products WHERE slug = 'ipad-air-m1-refurb-1394' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'AirPods 3 (Refurbished 2024)', 'airpods-3-refurb-7489', 'Certified Refurbished: Spatial audio, sweat resistant.', 46000000, 41400000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'AIRPODS-3-REFURB-7489-0' FROM products WHERE slug = 'airpods-3-refurb-7489' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'AIRPODS-3-REFURB-7489-1' FROM products WHERE slug = 'airpods-3-refurb-7489' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1542744094-24638eff58bb?w=800', 0 FROM products WHERE slug = 'airpods-3-refurb-7489' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1544244015-0df4b3ffc6b0?w=800', 1 FROM products WHERE slug = 'airpods-3-refurb-7489' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'MacBook Pro 14 M3 (Refurbished 2024)', 'macbook-pro-14-m3-refurb-6420', 'Certified Refurbished: M3 chip, Liquid Retina XDR.', 33000000, 29700000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'MACBOOK-PRO-14-M3-REFURB-6420-0' FROM products WHERE slug = 'macbook-pro-14-m3-refurb-6420' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'MACBOOK-PRO-14-M3-REFURB-6420-1' FROM products WHERE slug = 'macbook-pro-14-m3-refurb-6420' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1556656793-062ff9878258?w=800', 0 FROM products WHERE slug = 'macbook-pro-14-m3-refurb-6420' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1517336714731-489689fd1ca8?w=800', 1 FROM products WHERE slug = 'macbook-pro-14-m3-refurb-6420' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'MacBook Air 13 M2 (Refurbished 2023)', 'macbook-air-13-m2-refurb-8454', 'Certified Refurbished: M2 chip, thin and light.', 38000000, 34200000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'MACBOOK-AIR-13-M2-REFURB-8454-0' FROM products WHERE slug = 'macbook-air-13-m2-refurb-8454' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'MACBOOK-AIR-13-M2-REFURB-8454-1' FROM products WHERE slug = 'macbook-air-13-m2-refurb-8454' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1616348436168-de43ad0db179?w=800', 0 FROM products WHERE slug = 'macbook-air-13-m2-refurb-8454' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1616348436168-de43ad0db179?w=800', 1 FROM products WHERE slug = 'macbook-air-13-m2-refurb-8454' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'Apple Watch SE (Refurbished 2022)', 'apple-watch-se-2-refurb-6457', 'Certified Refurbished: Essential features, affordable.', 17000000, 15300000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'APPLE-WATCH-SE-2-REFURB-6457-0' FROM products WHERE slug = 'apple-watch-se-2-refurb-6457' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'APPLE-WATCH-SE-2-REFURB-6457-1' FROM products WHERE slug = 'apple-watch-se-2-refurb-6457' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1517336714731-489689fd1ca8?w=800', 0 FROM products WHERE slug = 'apple-watch-se-2-refurb-6457' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1585503418537-88331351ad99?w=800', 1 FROM products WHERE slug = 'apple-watch-se-2-refurb-6457' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'AirPods 3 (Refurbished 2022)', 'airpods-3-refurb-7891', 'Certified Refurbished: Spatial audio, sweat resistant.', 20000000, 18000000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'AIRPODS-3-REFURB-7891-0' FROM products WHERE slug = 'airpods-3-refurb-7891' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'AIRPODS-3-REFURB-7891-1' FROM products WHERE slug = 'airpods-3-refurb-7891' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1542744094-24638eff58bb?w=800', 0 FROM products WHERE slug = 'airpods-3-refurb-7891' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1585503418537-88331351ad99?w=800', 1 FROM products WHERE slug = 'airpods-3-refurb-7891' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'iPhone 15 Plus (Refurbished 2022)', 'iphone-15-plus-refurb-9326', 'Certified Refurbished: A16 Bionic, 48MP camera, big battery.', 24000000, 21600000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'IPHONE-15-PLUS-REFURB-9326-0' FROM products WHERE slug = 'iphone-15-plus-refurb-9326' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'IPHONE-15-PLUS-REFURB-9326-1' FROM products WHERE slug = 'iphone-15-plus-refurb-9326' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1611186871348-b1ec696e523b?w=800', 0 FROM products WHERE slug = 'iphone-15-plus-refurb-9326' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1511707171634-5f897ff02aa9?w=800', 1 FROM products WHERE slug = 'iphone-15-plus-refurb-9326' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'iPhone 15 Pro Max (Refurbished 2022)', 'iphone-15-pro-max-refurb-3483', 'Certified Refurbished: Titanium, A17 Pro, 5x Telephoto.', 36000000, 32400000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'IPHONE-15-PRO-MAX-REFURB-3483-0' FROM products WHERE slug = 'iphone-15-pro-max-refurb-3483' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'IPHONE-15-PRO-MAX-REFURB-3483-1' FROM products WHERE slug = 'iphone-15-pro-max-refurb-3483' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1510557880182-3d4d3cba35a5?w=800', 0 FROM products WHERE slug = 'iphone-15-pro-max-refurb-3483' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1544244015-0df4b3ffc6b0?w=800', 1 FROM products WHERE slug = 'iphone-15-pro-max-refurb-3483' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'iPhone 15 (Refurbished 2024)', 'iphone-15-refurb-6434', 'Certified Refurbished: A16 Bionic, 48MP camera, Dynamic Island.', 40000000, 36000000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'IPHONE-15-REFURB-6434-0' FROM products WHERE slug = 'iphone-15-refurb-6434' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'IPHONE-15-REFURB-6434-1' FROM products WHERE slug = 'iphone-15-refurb-6434' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1585503418537-88331351ad99?w=800', 0 FROM products WHERE slug = 'iphone-15-refurb-6434' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1591337676887-a217a6970a8a?w=800', 1 FROM products WHERE slug = 'iphone-15-refurb-6434' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'iPhone 15 Pro Max (Refurbished 2023)', 'iphone-15-pro-max-refurb-2928', 'Certified Refurbished: Titanium, A17 Pro, 5x Telephoto.', 45000000, 40500000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'IPHONE-15-PRO-MAX-REFURB-2928-0' FROM products WHERE slug = 'iphone-15-pro-max-refurb-2928' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'IPHONE-15-PRO-MAX-REFURB-2928-1' FROM products WHERE slug = 'iphone-15-pro-max-refurb-2928' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1544244015-0df4b3ffc6b0?w=800', 0 FROM products WHERE slug = 'iphone-15-pro-max-refurb-2928' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1511707171634-5f897ff02aa9?w=800', 1 FROM products WHERE slug = 'iphone-15-pro-max-refurb-2928' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'Apple Watch Series 9 (Refurbished 2023)', 'apple-watch-series-9-refurb-6795', 'Certified Refurbished: S9 chip, Double Tap gesture.', 10000000, 9000000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'APPLE-WATCH-SERIES-9-REFURB-6795-0' FROM products WHERE slug = 'apple-watch-series-9-refurb-6795' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'APPLE-WATCH-SERIES-9-REFURB-6795-1' FROM products WHERE slug = 'apple-watch-series-9-refurb-6795' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1591337676887-a217a6970a8a?w=800', 0 FROM products WHERE slug = 'apple-watch-series-9-refurb-6795' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1542744094-24638eff58bb?w=800', 1 FROM products WHERE slug = 'apple-watch-series-9-refurb-6795' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'iPhone 15 (Refurbished 2023)', 'iphone-15-refurb-7313', 'Certified Refurbished: A16 Bionic, 48MP camera, Dynamic Island.', 29000000, 26100000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'IPHONE-15-REFURB-7313-0' FROM products WHERE slug = 'iphone-15-refurb-7313' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'IPHONE-15-REFURB-7313-1' FROM products WHERE slug = 'iphone-15-refurb-7313' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1585503418537-88331351ad99?w=800', 0 FROM products WHERE slug = 'iphone-15-refurb-7313' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1517336714731-489689fd1ca8?w=800', 1 FROM products WHERE slug = 'iphone-15-refurb-7313' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'iPad Air M1 (Refurbished 2024)', 'ipad-air-m1-refurb-3125', 'Certified Refurbished: M1 chip, powerful and colorful.', 29000000, 26100000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'IPAD-AIR-M1-REFURB-3125-0' FROM products WHERE slug = 'ipad-air-m1-refurb-3125' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'IPAD-AIR-M1-REFURB-3125-1' FROM products WHERE slug = 'ipad-air-m1-refurb-3125' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1585503418537-88331351ad99?w=800', 0 FROM products WHERE slug = 'ipad-air-m1-refurb-3125' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1556656793-062ff9878258?w=800', 1 FROM products WHERE slug = 'ipad-air-m1-refurb-3125' ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES ((SELECT id FROM categories WHERE slug = 'apple'), 'MacBook Air 15 M2 (Refurbished 2024)', 'macbook-air-15-m2-refurb-6166', 'Certified Refurbished: M2 chip, big display, thin.', 10000000, 9000000, 500, '{}', 4.8, 100)
ON CONFLICT (slug) DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 1', 0, 50, 'MACBOOK-AIR-15-M2-REFURB-6166-0' FROM products WHERE slug = 'macbook-air-15-m2-refurb-6166' ON CONFLICT DO NOTHING;
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit) SELECT id, 'Variant 2', 100000, 50, 'MACBOOK-AIR-15-M2-REFURB-6166-1' FROM products WHERE slug = 'macbook-air-15-m2-refurb-6166' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1544244015-0df4b3ffc6b0?w=800', 0 FROM products WHERE slug = 'macbook-air-15-m2-refurb-6166' ON CONFLICT DO NOTHING;
INSERT INTO product_images (product_id, image_url, sort_order) SELECT id, 'https://images.unsplash.com/photo-1517336714731-489689fd1ca8?w=800', 1 FROM products WHERE slug = 'macbook-air-15-m2-refurb-6166' ON CONFLICT DO NOTHING;