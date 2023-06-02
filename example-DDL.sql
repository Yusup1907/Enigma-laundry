CREATE DATABASE enigmaLaundry;

CREATE TABLE mst_customers (
    id_customer VARCHAR(100) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    no_telp VARCHAR(20),
    alamat VARCHAR(150)
);

CREATE TABLE trx_order(
    id_order VARCHAR(100) PRIMARY KEY,
    customer_id VARCHAR(100),
    tanggal_masuk DATE,
    tanggal_keluar DATE,
    FOREIGN KEY (customer_id) REFERENCES mst_customers(id_customer)
);

CREATE TABLE mst_layanan (
    id_layanan VARCHAR(100) PRIMARY KEY,
    nama_layanan VARCHAR(100),
    harga INT,
    satuan VARCHAR(20) 
);

CREATE TABLE trx_order_detail (
    id_order_detail VARCHAR(100) PRIMARY KEY,    
    order_id VARCHAR(100),
    layanan_id VARCHAR(100),
    quantity INT,
    FOREIGN KEY (order_id) REFERENCES trx_order(id_order),
    FOREIGN KEY (layanan_id) REFERENCES mst_layanan(id_layanan)
);
