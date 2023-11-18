package pelayanan

import (
	"database/sql"
	"fmt"
	"log"
)

type Pelayanan struct {
	ID            int
	NamaPelayanan string
	Satuan        string
	Harga         float64
}

func showMenuPelayanan(db *sql.DB) {
	for {
		fmt.Println("\n==== Menu Pelayanan ====")
		fmt.Println("1. Daftar Pelayanan")
		fmt.Println("2. Tambah Pelayanan")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilih menu (0-2): ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			viewPelayanan(db)
		case 2:
			insertPelayanan(db)
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak ada.")
		}
	}
}

func viewPelayanan(db *sql.DB) {
	rows, err := db.Query("SELECT id_pelayanan, nama_pelayanan, satuan, harga FROM pelayanan")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("\n==== Data Pelayanan ====")
	fmt.Printf("%-5s %-30s %-10s %-15s\n", "ID", "Nama Pelayanan", "Satuan", "Harga")
	fmt.Println("=================================")

	for rows.Next() {
		var idPelayanan int
		var namaPelayanan, satuan string
		var harga float64
		if err := rows.Scan(&idPelayanan, &namaPelayanan, &satuan, &harga); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%-5d %-30s %-10s %-15.2f\n", idPelayanan, namaPelayanan, satuan, harga)
	}
	fmt.Println()
}

func insertPelayanan(db *sql.DB) {
	var namaPelayanan, satuan string
	var harga float64

	fmt.Print("Masukkan Nama Pelayanan: ")
	fmt.Scan(&namaPelayanan)
	fmt.Print("Masukkan Satuan: ")
	fmt.Scan(&satuan)
	fmt.Print("Masukkan Harga: ")
	fmt.Scan(&harga)

	//Validasi apakah harga lebih dari 0
	if harga <= 0 {
		fmt.Println("Harga harus lebih dari 0.")
		return
	}

	_, err := db.Exec("INSERT INTO pelayanan (nama_pelayanan, satuan, harga) VALUES ($1, $2, $3)", namaPelayanan, satuan, harga)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data Pelayanan berhasil ditambahkan.")
}
