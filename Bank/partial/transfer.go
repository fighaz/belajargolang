package partial

type Transfer struct {
	Nasabah         Nasabah
	Money           int
	NasabahPenerima Nasabah
}

func NewTransfer(nasabah Nasabah, money int, nasabahpenerima Nasabah) Transfer {
	return Transfer{nasabah, money, nasabahpenerima}
}
