package partial

import "fmt"

type Transaksi struct {
	Item []ItemTransaksi
	Uang int
}

func (transaksi Transaksi) TambahTransaksi(item ItemTransaksi) Transaksi {
	transaksi.Item = append(transaksi.Item, item)
	return transaksi
}
func (transaksi Transaksi) GetTotal() int {
	total := 0
	fmt.Println(transaksi.Item)
	for _, i := range transaksi.Item {
		total += i.GetSubTotal()
	}
	return total
}
