package main

import (
	"be-enigma-laundry-livecode/database"
	"be-enigma-laundry-livecode/pelanggan"
	"be-enigma-laundry-livecode/pelayanan"
	"be-enigma-laundry-livecode/transaksi"
	"fmt"
	"log"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for {
		fmt.Println("\n==== Menu Utama ====")
		fmt.Println("1. Daftar Pelanggan")
		fmt.Println("2. Tambah Pelayanan")
		fmt.Println("3. Lihat, Tambah, Update, dan Delete Transaksi Laundry")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu (0-3): ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			pelanggan.ShowMenuPelanggan(db)
		case 2:
			pelayanan.ShowMenuPelayanan(db)
		case 3:
			transaksi.ShowMenuTransaksi(db)
		case 0:
			fmt.Println("Keluar dari aplikasi.")
			return
		default:
			fmt.Println("Pilihan tidak ada.")
		}
	}
}
