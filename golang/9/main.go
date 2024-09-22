package main

import (
	"fmt"
)

type Produk struct {
	Nama    string
	Berat   int // dalam gram
	Panjang int // dalam mm
	Lebar   int // dalam mm
	Tinggi  int // dalam mm
}

func hitungBiayaKirim(produk []Produk) int {
	const tarifPerKg = 10000
	totalBiaya := 0

	for _, prod := range produk {
		beratAktual := float64(prod.Berat) / 1000

		// Hitung berat volumetrik: (panjang * lebar * tinggi) / 6000 (dalam cm³)
		beratVolumetrik := float64(prod.Panjang/10*prod.Lebar/10*prod.Tinggi/10) / 6000

		beratYangDikenakan := beratAktual
		if beratVolumetrik > beratAktual {
			beratYangDikenakan = beratVolumetrik
		}

		biayaProduk := int(beratYangDikenakan * float64(tarifPerKg))
		totalBiaya += biayaProduk
	}

	return totalBiaya
}

func main() {
	produk := []Produk{
		{"Produk A", 30, 115, 85, 25},
		{"Produk B", 28000, 1290, 300, 625},
	}

	totalBiaya := hitungBiayaKirim(produk)
	fmt.Printf("Total biaya kirim: Rp.%d,-\n", totalBiaya)
}
