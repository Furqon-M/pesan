package entities

type Pelanggan struct {
	Id        int
	NoPesanan string `validate:"required" label:"Noomor Pesanan"`
	Pesanan   string `validate:"required" label:"Pesanan"`
	Harga     string `validate:"required" label:"Harga"`
	Jumlah    string `validate:"required" label:"Jumlah"`
	Total     string `validate:"required" label:"Total"`
	Meja      string `validate:"required" label:"Meja"`
	Status    string
}
