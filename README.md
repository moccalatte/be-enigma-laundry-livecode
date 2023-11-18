# Aplikasi Enigma Laundry

Aplikasi ini adalah aplikasi konsol sederhana untuk manajemen laundry. Kita bisa melihat, menambah, memperbarui, dan menghapus data pelanggan, pelayanan, dan transaksi laundry.

## Persyaratan

- Go (Golang) terinstal di sistem
- PostgreSQL terinstal dan dapat diakses
- Koneksi internet untuk mendownload dependensi Golang

## Instalasi

1. Clone repositori ini:

   ```bash
   git clone https://git.enigmacamp.com/enigma-20/andre-milano/challenge-godb.git
   cd challenge-godb
   ```

2. Atur konfigurasi database:

   Buka file `enigmalaundry.go` dan sesuaikan variabel konfigurasi database di bagian `const` sesuai dengan informasi database.

3. Unduh dependensi:

   ```bash
   go mod download
   ```

4. Kompilasi aplikasi:

   ```bash
   go build
   ```

## Penggunaan

Jalankan aplikasi dengan perintah:

```bash
go run enigmalaundry.go
```

## MENU UTAMA

Aplikasi akan menampilkan menu utama dengan opsi sebagai berikut:

1. Daftar Pelanggan
2. Tambah Pelayanan
3. Lihat, Tambah, Update, dan Delete Transaksi Laundry
4. Keluar

## CONTOH PENGGUNAAN

1. Daftar Pelanggan
   Pilih menu "Daftar Pelanggan."
   Kita akan melihat daftar pelanggan yang ada.
   Untuk menambah pelanggan, pilih menu "Tambah Pelanggan" dan ikuti petunjuk yang diberikan.

2. Tambah Pelayanan
   Pilih menu "Tambah Pelayanan."
   Kita akan melihat daftar layanan laundry yang ada.
   Untuk menambah layanan, pilih menu "Tambah Pelayanan" dan ikuti petunjuk yang diberikan.

3. Lihat, Tambah, Update, dan Delete Transaksi Laundry
   Pilih menu "Lihat, Tambah, Update, dan Delete Transaksi Laundry."
   Kita akan melihat daftar transaksi laundry yang ada.
   Untuk menambah transaksi, pilih menu "Tambah Transaksi Laundry" dan ikuti petunjuk yang diberikan.
   Untuk memperbarui atau menghapus transaksi, pilih menu yang sesuai dan ikuti petunjuk yang diberikan.

4. Keluar dari Aplikasi
   Pilih menu "Keluar" pada menu utama untuk keluar dari aplikasi.

## Kontribusi

Silahkan berkontribusi dengan membuat pull request untuk perbaikan bug atau peningkatan fungsionalitas.
