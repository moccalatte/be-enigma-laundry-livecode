package transaksi

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Transaksi struct {
	ID             int
	IDPelanggan    int
	IDPelayanan    int
	Jumlah         int
	TanggalMasuk   string
	TanggalSelesai string
	DiterimaOleh   string
	TotalHarga     float64
}

func ShowMenuTransaksi(db *sql.DB) {
	for {
		fmt.Println("\n==== Menu Transaksi ====")
		fmt.Println("1. Daftar Transaksi Laundry")
		fmt.Println("2. Tambah Transaksi Laundry")
		fmt.Println("3. Update Transaksi Laundry")
		fmt.Println("4. Delete Transaksi Laundry")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilih menu (0-4): ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			ViewTransaksi(db)
		case 2:
			InsertTransaksi(db)
		case 3:
			UpdateTransaksi(db)
		case 4:
			DeleteTransaksi(db)
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak ada.")
		}
	}
}

func ViewTransaksi(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM transaksi_laundry")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("\n==== Data Transaksi Laundry ====")
	fmt.Printf("%-5s %-15s %-15s %-15s %-15s %-15s %-20s %-15s\n", "ID", "ID Pelanggan", "ID Pelayanan", "Jumlah", "Tanggal Masuk", "Tanggal Selesai", "Diterima oleh", "Total Harga")
	fmt.Println("===============================================================================================")

	for rows.Next() {
		var t Transaksi
		if err := rows.Scan(&t.ID, &t.IDPelanggan, &t.IDPelayanan, &t.Jumlah, &t.TanggalMasuk, &t.TanggalSelesai, &t.DiterimaOleh, &t.TotalHarga); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%-5d %-15d %-15d %-15d %-15s %-15s %-20s %-15.2f\n", t.ID, t.IDPelanggan, t.IDPelayanan, t.Jumlah, t.TanggalMasuk, t.TanggalSelesai, t.DiterimaOleh, t.TotalHarga)
	}
	fmt.Println()
}

func InsertTransaksi(db *sql.DB) {
	var t Transaksi
	var tanggalMasuk, tanggalSelesai string

	fmt.Print("Masukkan ID Pelanggan: ")
	fmt.Scan(&t.IDPelanggan)

	// Validasi apakah ID Pelanggan valid
	validPelanggan, err := isPelangganValid(db, t.IDPelanggan)
	if err != nil {
		log.Fatal(err)
	}
	if !validPelanggan {
		fmt.Println("ID Pelanggan tidak valid.")
		return
	}

	fmt.Print("Masukkan ID Pelayanan: ")
	fmt.Scan(&t.IDPelayanan)

	// Validasi apakah ID Pelayanan valid
	validPelayanan, err := isPelayananValid(db, t.IDPelayanan)
	if err != nil {
		log.Fatal(err)
	}
	if !validPelayanan {
		fmt.Println("ID Pelayanan tidak valid.")
		return
	}

	fmt.Print("Masukkan Jumlah: ")
	fmt.Scan(&t.Jumlah)

	// Validasi apakah jumlah lebih dari 0
	if t.Jumlah <= 0 {
		fmt.Println("Jumlah harus lebih dari 0.")
		return
	}

	fmt.Print("Masukkan Tanggal Masuk (format: YYYY-MM-DD): ")
	fmt.Scan(&tanggalMasuk)

	// Validasi apakah format tanggal masuk valid
	if !isValidDate(tanggalMasuk) {
		fmt.Println("Salah. Gunakan format YYYY-MM-DD.")
		return
	}

	fmt.Print("Masukkan Tanggal Selesai (format: YYYY-MM-DD): ")
	fmt.Scan(&tanggalSelesai)

	// Validasi apakah format tanggal selesai valid
	if !isValidDate(tanggalSelesai) {
		fmt.Println("Salah. Gunakan format YYYY-MM-DD.")
		return
	}

	t.TanggalMasuk = tanggalMasuk
	t.TanggalSelesai = tanggalSelesai

	fmt.Print("Masukkan Diterima oleh: ")
	fmt.Scan(&t.DiterimaOleh)

	// Validasi apakah total harga sesuai dengan pelayanan
	hargaPelayanan, err := getHargaPelayanan(db, t.IDPelayanan)
	if err != nil {
		log.Fatal(err)
	}
	t.TotalHarga = float64(t.Jumlah) * hargaPelayanan

	fmt.Printf("Total Harga (otomatis): %.2f\n", t.TotalHarga)

	// Validasi apakah diterima oleh valid
	if t.DiterimaOleh == "" {
		fmt.Println("Nama penerima tidak boleh kosong.")
		return
	}

	_, err = db.Exec("INSERT INTO transaksi_laundry (id_pelanggan, id_pelayanan, jumlah, tanggal_masuk, tanggal_selesai, diterima_oleh, total_harga) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		t.IDPelanggan, t.IDPelayanan, t.Jumlah, t.TanggalMasuk, t.TanggalSelesai, t.DiterimaOleh, t.TotalHarga)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data Transaksi Laundry berhasil ditambahkan.")
}

func UpdateTransaksi(db *sql.DB) {
	var t Transaksi
	var tanggalMasuk, tanggalSelesai string

	fmt.Print("Masukkan ID Transaksi yang ingin diupdate: ")
	fmt.Scan(&t.ID)

	// Validasi apakah ID Transaksi valid
	validTransaksi, err := isTransaksiValid(db, t.ID)
	if err != nil {
		log.Fatal(err)
	}
	if !validTransaksi {
		fmt.Println("ID Transaksi tidak valid.")
		return
	}

	fmt.Print("Masukkan ID Pelanggan: ")
	fmt.Scan(&t.IDPelanggan)

	// Validasi apakah ID Pelanggan valid
	validPelanggan, err := isPelangganValid(db, t.IDPelanggan)

	if err != nil {
		log.Fatal(err)
	}
	if !validPelanggan {
		fmt.Println("ID Pelanggan tidak valid.")
		return
	}

	fmt.Print("Masukkan ID Pelayanan: ")
	fmt.Scan(&t.IDPelayanan)

	// Validasi apakah ID Pelayanan valid
	validPelayanan, err := isPelayananValid(db, t.IDPelayanan)
	if err != nil {
		log.Fatal(err)
	}
	if !validPelayanan {
		fmt.Println("ID Pelayanan tidak valid.")
		return
	}

	fmt.Print("Masukkan Jumlah: ")
	fmt.Scan(&t.Jumlah)

	// Validasi apakah jumlah lebih dari 0
	if t.Jumlah <= 0 {
		fmt.Println("Jumlah harus lebih dari 0.")
		return
	}

	fmt.Print("Masukkan Tanggal Masuk (format: YYYY-MM-DD): ")
	fmt.Scan(&tanggalMasuk)

	// Validasi apakah format tanggal masuk valid
	if !isValidDate(tanggalMasuk) {
		fmt.Println("Salah. Gunakan format YYYY-MM-DD.")
		return
	}

	fmt.Print("Masukkan Tanggal Selesai (format: YYYY-MM-DD): ")
	fmt.Scan(&tanggalSelesai)

	// Validasi apakah format tanggal selesai valid
	if !isValidDate(tanggalSelesai) {
		fmt.Println("Salah. Gunakan format YYYY-MM-DD.")
		return
	}

	t.TanggalMasuk = tanggalMasuk
	t.TanggalSelesai = tanggalSelesai

	fmt.Print("Masukkan Diterima oleh: ")
	fmt.Scan(&t.DiterimaOleh)

	// Validasi apakah total harga sesuai dengan pelayanan
	hargaPelayanan, err := getHargaPelayanan(db, t.IDPelayanan)
	if err != nil {
		log.Fatal(err)
	}
	t.TotalHarga = float64(t.Jumlah) * hargaPelayanan

	fmt.Printf("Total Harga (otomatis): %.2f\n", t.TotalHarga)

	// Validasi apakah diterima oleh valid
	if t.DiterimaOleh == "" {
		fmt.Println("Nama penerima tidak boleh kosong.")
		return
	}

	_, err = db.Exec("UPDATE transaksi_laundry SET id_pelanggan=$1, id_pelayanan=$2, jumlah=$3, tanggal_masuk=$4, tanggal_selesai=$5, diterima_oleh=$6, total_harga=$7 WHERE id_transaksi=$8",
		t.IDPelanggan, t.IDPelayanan, t.Jumlah, t.TanggalMasuk, t.TanggalSelesai, t.DiterimaOleh, t.TotalHarga, t.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data Transaksi Laundry berhasil diupdate.")
}

func DeleteTransaksi(db *sql.DB) {
	var idTransaksi int

	fmt.Print("Masukkan ID Transaksi yang mau dihapus: ")
	fmt.Scan(&idTransaksi)

	//Validasi apakah ID Transaksi valid
	validTransaksi, err := isTransaksiValid(db, idTransaksi)
	if err != nil {
		log.Fatal(err)
	}
	if !validTransaksi {
		fmt.Println("ID Transaksi tidak valid.")
		return
	}

	_, err = db.Exec("DELETE FROM transaksi_laundry WHERE id_transaksi=$1", idTransaksi)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data Transaksi Laundry berhasil dihapus.")
}

func isPelangganValid(db *sql.DB, idPelanggan int) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM pelanggan WHERE id_pelanggan = $1", idPelanggan).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	return count > 0, nil
}

func isPelayananValid(db *sql.DB, idPelayanan int) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM pelayanan WHERE id_pelayanan = $1", idPelayanan).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	return count > 0, nil
}

func isTransaksiValid(db *sql.DB, idTransaksi int) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM transaksi_laundry WHERE id_transaksi = $1", idTransaksi).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	return count > 0, nil
}

func isValidDate(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	return err == nil
}

func getHargaPelayanan(db *sql.DB, idPelayanan int) (float64, error) {
	var harga float64
	err := db.QueryRow("SELECT harga FROM pelayanan WHERE id_pelayanan = $1", idPelayanan).Scan(&harga)
	if err != nil {
		log.Fatal(err)
	}

	return harga, nil
}
