package main

import (
	"coba/partial"
	"fmt"
	"strings"
)

func main() {

	var toko = partial.Toko{SemuaBarang: []partial.Barang{
		{Namabarang: "buku", Harga: 3000},
		{Namabarang: "pensil", Harga: 2000},
		{Namabarang: "gunting", Harga: 5000},
		{Namabarang: "pulpen", Harga: 4000},
	}}
	var pilihan, jumlah, pilihMenu int
	var step string
	transaksi := partial.Transaksi{Item: []partial.ItemTransaksi{}, Uang: 0}
	for {
		fmt.Println("==================================================")
		fmt.Println("1.Tambah Barang ")
		fmt.Println("2.Transaksi")
		fmt.Println("==================================================")

		fmt.Println("Masukkan Menu Yang akan Dipilih")
		fmt.Scan(&pilihMenu)

		if pilihMenu == 1 {
			fmt.Println("Masukkan Nama Barang")
			var nama string
			fmt.Scan(&nama)
			fmt.Println("Masukkan Harga Barang")
			var harga int
			fmt.Scan(&harga)

			toko = toko.TambahBarang(partial.Barang{Namabarang: nama, Harga: harga})
		} else if pilihMenu == 2 {
			for {
				toko.PrintSemuaBarang()
				fmt.Println("Masukkan No barang yang anda akan beli")
				fmt.Scan(&pilihan)
				pilihan = pilihan - 1
				barang := toko.SemuaBarang[pilihan]
				fmt.Println("Anda membeli", toko.SemuaBarang[pilihan].Namabarang, "dengan harga", toko.SemuaBarang[pilihan].Harga, "Masukkan jumlah  barang yang anda akan beli")
				fmt.Scan(&jumlah)
				transaksi = transaksi.TambahTransaksi(partial.ItemTransaksi{Barang: barang, Jumlah: jumlah})
				fmt.Println(transaksi.Item)

				itemtransaksi := partial.ItemTransaksiBaru(barang, jumlah)
				subtotal := itemtransaksi.GetSubTotal()

				// subtotal := toko.SemuaBarang[pilihan].Harga * jumlah
				fmt.Println("Subtotal belanjaan anda adalah ", subtotal)
				fmt.Println("Apakah Anda Ingin membeli barang yang lain (Y/T)")
				fmt.Scan(&step)
				if strings.EqualFold(step, "T") {

					break
				}

			}

			fmt.Println("Total Belanjan Anda Adalah ", transaksi.GetTotal())
			fmt.Println("Masukkan uang pembayaran")
			fmt.Scan(&transaksi.Uang)
			var bayar int
			bayar = transaksi.Uang - transaksi.GetTotal()
			if bayar > 0 {
				fmt.Println("Pembayaran Sukses ,Uang kembalian anda sebesar ", bayar)
			} else if bayar == 0 {
				fmt.Println("Pembayaran Sukses ,Uang Anda pas ")
			} else {
				fmt.Println("Pembayaran Gagal ,Uang yang Anda Masukkan Kurang ")
			}
		}
	}
}
