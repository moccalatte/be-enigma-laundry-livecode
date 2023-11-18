-- Sample Data Pelanggan
INSERT INTO pelanggan (nama, no_hp) VALUES
    ('Mirna', '0812345678'),
    ('Jessica', '0812654987');

-- Sample Data Pelayanan
INSERT INTO pelayanan (nama_pelayanan, satuan, harga) VALUES
    ('Cuci + Setrika', 'KG', 7000.00),
    ('Laundry Bedcover', 'Buah', 50000.00),
    ('Laundry Boneka', 'Buah', 25000.00);

-- Sample Data Transaksi
INSERT INTO transaksi_laundry (id_pelanggan, id_pelayanan, jumlah, tanggal_masuk, tanggal_selesai, diterima_oleh, total_harga) VALUES
    (1, 1, 5, '2022-08-18', '2022-08-20', 'Mirna', 35000.00),
    (2, 2, 1, '2022-08-18', '2022-08-20', 'Jessica', 50000.00),
    (1, 3, 2, '2022-08-18', '2022-08-20', 'Mirna', 50000.00);
