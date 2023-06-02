----------------------- Customer -----------------------------------------------
-- Select Customer
SELECT * FROM mst_customers;

-- Select Customer Berdasarkan Id
SELECT id_customer, name, no_telp, alamat FROM mst_customers WHERE id_customer = $1;

-- Insert Customer
INSERT INTO mst_customers (id_customer, name, no_telp, alamat) VALUES ($1, $2, $3, $4)

-- Update Customer
UPDATE mst_customers SET name = $2, no_telp = $3, alamat = $4 WHERE id_customer = $1;

-- Delete Customer
DELETE FROM mst_customers WHERE id_customer = $1;

--------------------------------- Layanan ---------------------------------------

-- Select Semua Layanan
SELECT * FROM mst_layanan;

-- Select Layanan Berdasarkan Id
SELECT id_layanan, nama_layanan, harga, satuan FROM mst_layanan WHERE id_layanan = $1;

-- Insert Layanan
INSERT INTO mst_layanan (id_layanan, nama_layanan, harga, satuan) VALUES ($1, $2, $3, $4);

-- Update Layanan
UPDATE mst_layanan SET nama_layanan = $2, harga = $3, satuan = $4 WHERE id_layanan = $1;

-- Delete Layanan
DELETE FROM mst_layanan WHERE id_layanan = $1;


--------------------------- Order ---------------------------------------------------

-- Select Semua Order
SELECT id_order, customer_id, tanggal_masuk, tanggal_keluar FROM trx_order;

-- Select Order Berdasarkan Id
SELECT id_order, customer_id, tanggal_masuk, tanggal_keluar FROM trx_order WHERE id_order = $1;

-- Insert Order
INSERT INTO trx_order (id_order, customer_id, tanggal_masuk, tanggal_keluar) VALUES ($1, $2, $3, $4);


--------------------------- Order Detail ----------------------------------------------

-- Select Semua Order Detail
SELECT id_order_detail, order_id, layanan_id, quantity FROM trx_order_detail;

-- Select Order Detail Berdasarkan Id
SELECT id_order_detail, order_id, layanan_id, quantity FROM trx_order_detail WHERE id_order_detail = $1;

-- Insert Order Detail
INSERT INTO trx_order_detail (id_order_detail, order_id, layanan_id, quantity) VALUES ($1, $2, $3, $4);



