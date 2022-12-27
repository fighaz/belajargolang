package main

import (
	"fmt"
	"strings"
)

type barang struct {
	id         int
	namabarang string
	harga      int
}

func main() {

	var semuaBarang = []barang{
		{id: 1, namabarang: "buku", harga: 3000},
		{id: 2, namabarang: "pensil", harga: 2000},
		{id: 3, namabarang: "gunting", harga: 5000},
		{id: 4, namabarang: "pulpen", harga: 4000},
	}
	var pilihan, total, uang int
	var step string

	for {
		tampilBarang(semuaBarang)

		fmt.Println("Masukkan No barang yang anda akan beli")
		fmt.Scan(&pilihan)
		total += beliBarang(semuaBarang, pilihan)
		fmt.Println("Apakah Anda Ingin membeli barang yang lain (Y/T)")
		fmt.Scan(&step)
		if strings.EqualFold(step, "T") {
			break
		}
	}
	fmt.Println("Total Belanjan Anda Adalah ", total)
	fmt.Println("Masukkan uang pembayaran")
	fmt.Scan(&uang)
	pembayaran(uang, total)
}
func tampilBarang(b []barang) {
	for _, b := range b {
		fmt.Println(b.id, " . ", b.namabarang, "-", b.harga)
	}

}
func pembayaran(uang int, total int) {
	var bayar int
	bayar = uang - total
	if bayar > 0 {
		fmt.Println("Pembayaran Sukses ,Uang kembalian anda sebesar ", bayar)
	} else if bayar == 0 {
		fmt.Println("Pembayaran Sukses ,Uang Anda pas ")
	} else {
		fmt.Println("Pembayaran Gagal ,Uang yang Anda Masukkan Kurang ")
	}
}

func beliBarang(b []barang, pilih int) int {
	var subtotal, jumlah int
	for i := 0; i < len(b); i++ {
		if (pilih - 1) == i {
			fmt.Println("Anda membeli", b[i].namabarang, "dengan harga", b[i].harga, "Masukkan jumlah  barang yang anda akan beli")
			fmt.Scan(&jumlah)
			subtotal = b[i].harga * jumlah
			fmt.Println("Subtotal belanjaan anda adalah ", subtotal)
		}

	}
	return subtotal

}
