package entity

type Pelayanan struct {
	ID            int `json:"id"`
	NamaPelayanan string `json:"namaPelayanan"`
	Satuan        string `json:"satuan"`
	Harga         float64 `json:"harga"`
}