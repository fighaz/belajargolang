package partial

import "fmt"

type Bank struct {
	SemuaNasabah    []Nasabah
	RiwayatTransfer []Transfer
}

func (data Bank) PrintDataNasabah() {
	for index, nasabah := range data.SemuaNasabah {
		fmt.Println((index + 1), ".", nasabah.Nama, "-", nasabah.NoRek, "-", nasabah.NamaBank, "-", nasabah.Saldo)
	}
}

func (bank *Bank) TopUp(nasabah int, Money int) {
	bank.SemuaNasabah[nasabah].Saldo = bank.SemuaNasabah[nasabah].Saldo + Money

}
func (bank *Bank) KirimUang(nasabahpengirim, money, nasabahpenerima int) {
	if money > bank.SemuaNasabah[nasabahpengirim].Saldo {
		fmt.Println("Saldo Anda Tidak Mencukupi")
	} else {
		bank.SemuaNasabah[nasabahpengirim].Saldo = bank.SemuaNasabah[nasabahpengirim].Saldo - money
		bank.SemuaNasabah[nasabahpenerima].Saldo = bank.SemuaNasabah[nasabahpenerima].Saldo + money
	}

}
