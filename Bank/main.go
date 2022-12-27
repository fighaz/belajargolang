package main

import (
	"bank/partial"
	"fmt"
)

func main() {
	var scan int
	fmt.Println("Selamat Datang")
	bank := partial.Bank{SemuaNasabah: []partial.Nasabah{
		{Nama: "Adi", NoRek: 12345, NamaBank: "Mandiri", Saldo: 0},
		{Nama: "Iqbal", NoRek: 11111, NamaBank: "BCA", Saldo: 0},
		{Nama: "Rama", NoRek: 22222, NamaBank: "BNI", Saldo: 0},
		{Nama: "Yusri", NoRek: 33333, NamaBank: "Mandiri", Saldo: 0},
	}}
	for {
		fmt.Println("Pilih Menu Dibawah Ini")
		fmt.Println("1.Top Up Saldo")
		fmt.Println("2.Transfer")

		fmt.Scan(&scan)
		if scan == 1 {

			bank.PrintDataNasabah()
			var pilih int
			fmt.Println("Masukkan No Akun Yang anda Ingin Top Up")
			fmt.Scan(&pilih)

			nasabah := pilih - 1
			var money int
			fmt.Println("Masukkan Uang Yang anda Ingin Top Up")
			fmt.Scan(&money)
			// topup := partial.NewTopup(nasabah, money)
			// topup.TopUp(money)
			bank.TopUp(nasabah, money)
			bank.PrintDataNasabah()
		} else if scan == 2 {
			var pengirim, penerima int
			bank.PrintDataNasabah()
			fmt.Println("Pilih Akun Anda")
			fmt.Scan(&pengirim)
			nasabah := pengirim - 1
			var jumlah int
			fmt.Println("Masukkan Uang Yang anda Ingin kirim")
			fmt.Scan(&jumlah)
			fmt.Println("Pilih Akun yang anda akan Kirim")
			fmt.Scan(&penerima)
			nasabahpenerima := penerima - 1
			bank.KirimUang(nasabah, jumlah, nasabahpenerima)
			bank.PrintDataNasabah()
		} else {
			fmt.Println("Terima Kasih")
		}
	}
}
