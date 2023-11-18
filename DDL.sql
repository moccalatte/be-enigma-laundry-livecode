-- Tabel Master Pelanggan
CREATE TABLE pelanggan (
    id_pelanggan SERIAL PRIMARY KEY,
    nama VARCHAR(100),
    no_hp VARCHAR(15) UNIQUE
);

-- Tabel Master Pelayanan
CREATE TABLE pelayanan (
    id_pelayanan SERIAL PRIMARY KEY,
    nama_pelayanan VARCHAR(100),
    satuan VARCHAR(20),
    harga DECIMAL(10, 2)
);

-- Tabel Transaksi Laundry
CREATE TABLE transaksi_laundry (
    id_transaksi SERIAL PRIMARY KEY,
    id_pelanggan INT REFERENCES pelanggan(id_pelanggan),
    id_pelayanan INT REFERENCES pelayanan(id_pelayanan),
    jumlah INT,
    tanggal_masuk DATE,
    tanggal_selesai DATE,
    diterima_oleh VARCHAR(100),
    total_harga DECIMAL(10, 2),
    CONSTRAINT validasi1 CHECK (jumlah >= 0),
    CONSTRAINT validasi2 CHECK (total_harga >= 0)
);
