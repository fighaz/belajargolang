package partial

type ItemTransaksi struct {
	Barang Barang
	Jumlah int
}

func ItemTransaksiBaru(barang Barang, jumlah int) ItemTransaksi {
	return ItemTransaksi{barang, jumlah}
}
func (item ItemTransaksi) GetSubTotal() int {
	subtotal := item.Barang.Harga * item.Jumlah
	return subtotal
}
