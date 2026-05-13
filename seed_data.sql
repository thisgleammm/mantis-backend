-- =============================================================
-- MANTIS MARKETPLACE — SEED DATA
-- Run after all migrations have been applied.
-- =============================================================

-- -------------------------------------------------------------
-- 1. CATEGORIES
-- -------------------------------------------------------------
INSERT INTO categories (name, slug, description, icon_image_url, banner_image_url, is_active)
VALUES
  ('Pakaian Pria',   'pakaian-pria',   'Koleksi pakaian pria terkini',       NULL, NULL, true),
  ('Pakaian Wanita', 'pakaian-wanita', 'Koleksi pakaian wanita terkini',     NULL, NULL, true),
  ('Elektronik',     'elektronik',     'Gadget & perangkat elektronik',       NULL, NULL, true),
  ('Olahraga',       'olahraga',       'Perlengkapan & pakaian olahraga',     NULL, NULL, true),
  ('Sepatu',         'sepatu',         'Sepatu pria, wanita, dan anak',       NULL, NULL, true),
  ('Tas & Dompet',   'tas-dompet',     'Tas fashion dan dompet premium',      NULL, NULL, true);

-- -------------------------------------------------------------
-- 2. PRODUCTS
-- -------------------------------------------------------------
INSERT INTO products (category_id, name, slug, description, base_price, discount_price, weight, specifications, rating_average, rating_count)
VALUES
  -- Pakaian Pria (id = 1)
  (1, 'Kemeja Linen Premium', 'kemeja-linen-premium',
   'Kemeja bahan linen berkualitas tinggi, nyaman dipakai sepanjang hari. Cocok untuk casual maupun semi-formal.',
   299000, 249000, 300,
   '{"material":"Linen 100%","fit":"Regular Fit","care":"Machine wash cold"}',
   4.7, 128),

  (1, 'Kaos Oversize Graphic', 'kaos-oversize-graphic',
   'Kaos oversize dengan desain grafis unik edisi terbatas. Bahan cotton combed 30s anti-kusut.',
   189000, NULL, 250,
   '{"material":"Cotton Combed 30s","fit":"Oversized","care":"Hand wash recommended"}',
   4.5, 89),

  -- Pakaian Wanita (id = 2)
  (2, 'Dress Midi Floral', 'dress-midi-floral',
   'Dress midi motif bunga elegan dengan bahan viscose yang jatuh sempurna. Cocok untuk pesta maupun hangout.',
   450000, 385000, 350,
   '{"material":"Viscose","length":"Midi (knee-length)","care":"Dry clean only"}',
   4.8, 214),

  (2, 'Blouse Satin Polos', 'blouse-satin-polos',
   'Blouse satin minimalis dengan pilihan warna netral. Mudah dipadukan dengan berbagai outfit.',
   275000, NULL, 200,
   '{"material":"Satin Silk","fit":"Relaxed","care":"Hand wash cold"}',
   4.6, 76),

  -- Elektronik (id = 3)
  (3, 'Earbuds TWS Pro Max', 'earbuds-tws-pro-max',
   'True wireless earbuds dengan active noise cancellation, bass yang dalam, dan baterai 36 jam. IPX5 water resistant.',
   899000, 749000, 180,
   '{"connectivity":"Bluetooth 5.3","battery":"36 hours total","driver":"11mm dynamic","waterproof":"IPX5"}',
   4.9, 502),

  (3, 'Smartwatch Fitness Ultra', 'smartwatch-fitness-ultra',
   'Smartwatch dengan monitoring kesehatan lengkap: heart rate, SpO2, stress level, dan GPS built-in.',
   1299000, NULL, 280,
   '{"display":"1.96 AMOLED","battery":"14 days","sensors":"Heart Rate, SpO2, GPS","waterproof":"5ATM"}',
   4.7, 337),

  -- Olahraga (id = 4)
  (4, 'Jersey Olahraga Dry-Fit', 'jersey-olahraga-dry-fit',
   'Jersey olahraga teknologi dry-fit yang menyerap keringat dengan cepat. Cocok untuk gym, lari, atau futsal.',
   220000, 185000, 200,
   '{"material":"Polyester Dry-Fit","technology":"Moisture wicking","fit":"Athletic Slim"}',
   4.5, 193),

  -- Sepatu (id = 5)
  (5, 'Sneakers Casual Low-Cut', 'sneakers-casual-low-cut',
   'Sneakers casual low-cut dengan sol EVA yang empuk dan upper canvas yang breathable. Ringan & nyaman sepanjang hari.',
   525000, 450000, 600,
   '{"upper":"Canvas","sole":"EVA","insole":"Memory foam","closure":"Lace-up"}',
   4.6, 418),

  -- Tas & Dompet (id = 6)
  (6, 'Tote Bag Canvas Premium', 'tote-bag-canvas-premium',
   'Tote bag canvas tebal dengan jahitan kuat. Mampu menampung laptop 14" dan perlengkapan harian.',
   350000, NULL, 450,
   '{"material":"Canvas 600D","capacity":"25L","laptop_compartment":"Up to 14 inch","strap":"Adjustable"}',
   4.7, 162);

-- -------------------------------------------------------------
-- 3. PRODUCT VARIANTS
-- Catatan: product IDs diambil dari urutan insert di atas (1-9).
-- Sesuaikan jika tabel sudah ada data sebelumnya.
-- -------------------------------------------------------------

-- Kemeja Linen Premium (product_id = 1)
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit)
VALUES
  (1, 'S / Putih',  0,     50, 'KLP-S-WHT'),
  (1, 'M / Putih',  0,     80, 'KLP-M-WHT'),
  (1, 'L / Putih',  0,     60, 'KLP-L-WHT'),
  (1, 'XL / Putih', 20000, 40, 'KLP-XL-WHT'),
  (1, 'S / Biru',   0,     45, 'KLP-S-BLU'),
  (1, 'M / Biru',   0,     70, 'KLP-M-BLU'),
  (1, 'L / Biru',   0,     55, 'KLP-L-BLU'),
  (1, 'XL / Biru',  20000, 30, 'KLP-XL-BLU');

-- Kaos Oversize Graphic (product_id = 2)
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit)
VALUES
  (2, 'S / Hitam',  0, 100, 'KOG-S-BLK'),
  (2, 'M / Hitam',  0, 120, 'KOG-M-BLK'),
  (2, 'L / Hitam',  0,  90, 'KOG-L-BLK'),
  (2, 'XL / Hitam', 0,  60, 'KOG-XL-BLK'),
  (2, 'S / Krem',   0,  80, 'KOG-S-CRM'),
  (2, 'M / Krem',   0, 100, 'KOG-M-CRM');

-- Dress Midi Floral (product_id = 3)
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit)
VALUES
  (3, 'XS / Merah Muda', 0,     30, 'DMF-XS-PNK'),
  (3, 'S / Merah Muda',  0,     50, 'DMF-S-PNK'),
  (3, 'M / Merah Muda',  0,     45, 'DMF-M-PNK'),
  (3, 'L / Merah Muda',  0,     25, 'DMF-L-PNK'),
  (3, 'S / Biru Muda',   10000, 40, 'DMF-S-LBL'),
  (3, 'M / Biru Muda',   10000, 35, 'DMF-M-LBL');

-- Blouse Satin Polos (product_id = 4)
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit)
VALUES
  (4, 'S / Krem', 0, 60, 'BSP-S-CRM'),
  (4, 'M / Krem', 0, 80, 'BSP-M-CRM'),
  (4, 'L / Krem', 0, 55, 'BSP-L-CRM'),
  (4, 'S / Abu',  0, 50, 'BSP-S-GRY'),
  (4, 'M / Abu',  0, 65, 'BSP-M-GRY');

-- Earbuds TWS Pro Max (product_id = 5)
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit)
VALUES
  (5, 'Hitam',  0,      200, 'ETPM-BLK'),
  (5, 'Putih',  0,      180, 'ETPM-WHT'),
  (5, 'Navy',   50000,   80, 'ETPM-NVY');

-- Smartwatch Fitness Ultra (product_id = 6)
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit)
VALUES
  (6, 'Hitam / Strap Silikon',    0,       120, 'SFU-BLK-SIL'),
  (6, 'Silver / Strap Silikon',   0,        90, 'SFU-SLV-SIL'),
  (6, 'Gold / Strap Milanese', 150000,     50, 'SFU-GLD-MIL');

-- Jersey Olahraga Dry-Fit (product_id = 7)
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit)
VALUES
  (7, 'S / Hitam',  0, 150, 'JOD-S-BLK'),
  (7, 'M / Hitam',  0, 180, 'JOD-M-BLK'),
  (7, 'L / Hitam',  0, 130, 'JOD-L-BLK'),
  (7, 'XL / Hitam', 0,  90, 'JOD-XL-BLK'),
  (7, 'S / Merah',  0, 100, 'JOD-S-RED'),
  (7, 'M / Merah',  0, 120, 'JOD-M-RED');

-- Sneakers Casual Low-Cut (product_id = 8)
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit)
VALUES
  (8, '38 / Putih', 0,     60, 'SCL-38-WHT'),
  (8, '39 / Putih', 0,     80, 'SCL-39-WHT'),
  (8, '40 / Putih', 0,     90, 'SCL-40-WHT'),
  (8, '41 / Putih', 0,     75, 'SCL-41-WHT'),
  (8, '42 / Putih', 0,     50, 'SCL-42-WHT'),
  (8, '40 / Hitam', 20000, 70, 'SCL-40-BLK'),
  (8, '41 / Hitam', 20000, 65, 'SCL-41-BLK'),
  (8, '42 / Hitam', 20000, 40, 'SCL-42-BLK');

-- Tote Bag Canvas Premium (product_id = 9)
INSERT INTO product_variants (product_id, variant_name, price_extra, stock, stock_keeping_unit)
VALUES
  (9, 'Natural / Coklat Tua',  0,     120, 'TBCP-NAT-BRN'),
  (9, 'Hitam / Hitam',         0,     100, 'TBCP-BLK-BLK'),
  (9, 'Krem / Tan',            30000,  70, 'TBCP-CRM-TAN');

-- -------------------------------------------------------------
-- 4. PRODUCT IMAGES
-- Menggunakan Unsplash source yang stabil sebagai placeholder.
-- Ganti URL dengan CDN/storage aktual saat production.
-- -------------------------------------------------------------

-- Kemeja Linen Premium (product_id = 1)
INSERT INTO product_images (product_id, image_url, sort_order)
VALUES
  (1, 'https://images.unsplash.com/photo-1596755094514-f87e34085b2c?w=800', 0),
  (1, 'https://images.unsplash.com/photo-1602810318383-e386cc2a3ccf?w=800', 1),
  (1, 'https://images.unsplash.com/photo-1598033129183-c4f50c736f10?w=800', 2);

-- Kaos Oversize Graphic (product_id = 2)
INSERT INTO product_images (product_id, image_url, sort_order)
VALUES
  (2, 'https://images.unsplash.com/photo-1521572163474-6864f9cf17ab?w=800', 0),
  (2, 'https://images.unsplash.com/photo-1503341504253-dff4815485f1?w=800', 1);

-- Dress Midi Floral (product_id = 3)
INSERT INTO product_images (product_id, image_url, sort_order)
VALUES
  (3, 'https://images.unsplash.com/photo-1572804013309-59a88b7e92f1?w=800', 0),
  (3, 'https://images.unsplash.com/photo-1585487000160-6ebcfceb0d03?w=800', 1),
  (3, 'https://images.unsplash.com/photo-1496747611176-843222e1e57c?w=800', 2);

-- Blouse Satin Polos (product_id = 4)
INSERT INTO product_images (product_id, image_url, sort_order)
VALUES
  (4, 'https://images.unsplash.com/photo-1564257631407-4deb1f99d992?w=800', 0),
  (4, 'https://images.unsplash.com/photo-1509631179647-0177331693ae?w=800', 1);

-- Earbuds TWS Pro Max (product_id = 5)
INSERT INTO product_images (product_id, image_url, sort_order)
VALUES
  (5, 'https://images.unsplash.com/photo-1590658268037-6bf12165a8df?w=800', 0),
  (5, 'https://images.unsplash.com/photo-1608156639585-b3a776aee0e1?w=800', 1),
  (5, 'https://images.unsplash.com/photo-1606220945770-b5b6c2c55bf1?w=800', 2);

-- Smartwatch Fitness Ultra (product_id = 6)
INSERT INTO product_images (product_id, image_url, sort_order)
VALUES
  (6, 'https://images.unsplash.com/photo-1523275335684-37898b6baf30?w=800', 0),
  (6, 'https://images.unsplash.com/photo-1434493789847-2f02dc6ca35d?w=800', 1);

-- Jersey Olahraga Dry-Fit (product_id = 7)
INSERT INTO product_images (product_id, image_url, sort_order)
VALUES
  (7, 'https://images.unsplash.com/photo-1562183241-b937e9102303?w=800', 0),
  (7, 'https://images.unsplash.com/photo-1517466787929-bc90951d0974?w=800', 1);

-- Sneakers Casual Low-Cut (product_id = 8)
INSERT INTO product_images (product_id, image_url, sort_order)
VALUES
  (8, 'https://images.unsplash.com/photo-1542291026-7eec264c27ff?w=800', 0),
  (8, 'https://images.unsplash.com/photo-1515955656352-a1fa3ffcd111?w=800', 1),
  (8, 'https://images.unsplash.com/photo-1491553895911-0055eca6402d?w=800', 2);

-- Tote Bag Canvas Premium (product_id = 9)
INSERT INTO product_images (product_id, image_url, sort_order)
VALUES
  (9, 'https://images.unsplash.com/photo-1548036328-c9fa89d128fa?w=800', 0),
  (9, 'https://images.unsplash.com/photo-1547949003-9792a18a2601?w=800', 1);
