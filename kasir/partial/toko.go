package partial

import "fmt"

type Toko struct {
	SemuaBarang []Barang
	Transaksi   []ItemTransaksi
}

func TokoBaru(semuabarang []Barang, transaksi []ItemTransaksi) Toko {
	return Toko{SemuaBarang: semuabarang, Transaksi: transaksi}
}
func (toko Toko) TambahBarang(barang Barang) Toko {
	toko.SemuaBarang = append(toko.SemuaBarang, barang)
	return toko
}

func (toko Toko) PrintSemuaBarang() {
	for i, b := range toko.SemuaBarang {
		fmt.Println((i + 1), " . ", b.Namabarang, "-", b.Harga)
	}
}
