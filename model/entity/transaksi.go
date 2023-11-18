package entity

type Transaksi struct {
	ID             int `json:"id"`
	IDPelanggan    int `json:"idPelanggan"`
	IDPelayanan    int `json:"idPelayanan"`
	Jumlah         int `json:"jumlah"`
	TanggalMasuk   string `json:"tanggalMasuk"`
	TanggalSelesai string `json:"tanggalSelesai"`
	DiterimaOleh   string `json:"diterimaOleh"`
	TotalHarga     float64 `json:"totalHarga"`
}