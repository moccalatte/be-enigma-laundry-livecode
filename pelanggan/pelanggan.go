package pelanggan

import (
	"database/sql"
	"fmt"
	"log"
)

type Pelanggan struct {
	ID   int
	Nama string
	NoHP string
}

func ShowMenuPelanggan(db *sql.DB) {
	for {
		fmt.Println("\n==== Menu Pelanggan ====")
		fmt.Println("1. Daftar Pelanggan")
		fmt.Println("2. Tambah Pelanggan")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilih menu (0-2): ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			ViewPelanggan(db)
		case 2:
			InsertPelanggan(db)
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak ada.")
		}
	}
}

func ViewPelanggan(db *sql.DB) {
	rows, err := db.Query("SELECT id_pelanggan, nama, no_hp FROM pelanggan")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("\n==== Data Pelanggan ====")
	fmt.Printf("%-5s %-20s %-15s\n", "ID", "Nama", "No HP")
	fmt.Println("-------------------------")

	for rows.Next() {
		var idPelanggan int
		var nama, noHP string
		if err := rows.Scan(&idPelanggan, &nama, &noHP); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%-5d %-20s %-15s\n", idPelanggan, nama, noHP)
	}
	fmt.Println()
}

func InsertPelanggan(db *sql.DB) {
	var nama, noHP string

	fmt.Print("Masukkan Nama Pelanggan: ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan No HP Pelanggan: ")
	fmt.Scan(&noHP)

	//Validasi apakah no HP sudah ada
	exists, err := IsNoHPExists(db, noHP)
	if err != nil {
		log.Fatal(err)
	}

	if exists {
		fmt.Println("No HP sudah ada. Silahkan masukkan No HP lain.")
		return
	}

	_, err = db.Exec("INSERT INTO pelanggan (nama, no_hp) VALUES ($1, $2)", nama, noHP)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data Pelanggan berhasil ditambahkan.")
}

func IsNoHPExists(db *sql.DB, noHP string) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM pelanggan WHERE no_hp = $1", noHP).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}