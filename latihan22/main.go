package main

import (
	"fmt"

	"github.com/reza/praktikum/latihan22/tampil"
)

type Orang struct {
	nama         string
	umur         int
	pekerjaan    string
	ukuranSepatu float32
}

// Struct Sebagai Argumen Dari Function
func cetakOrang(org Orang) {
	fmt.Print("Nama orang pertama adalah ", org.nama, ". Saat ini dia berumur ", org.umur, " tahun")
	fmt.Print(". Sehari-hari pekerjaannya adalah seorang ", org.pekerjaan)
	fmt.Print(" dan ukuran sepatunya adalah ", org.ukuranSepatu)
	fmt.Print("\n")
}
func main() {
	tampil.Cetakjudul("Program Struct")
	//deklarasi variabel
	var orang1 Orang
	var orang2 Orang

	//memberi nilai pada struct
	orang1.nama = "Budi"
	orang1.umur = 22
	orang1.pekerjaan = "Nelayan"
	orang1.ukuranSepatu = 43.5

	//memberi nilai pada struct
	orang2.nama = "Eko"
	orang2.umur = 24
	orang2.pekerjaan = "Petani"
	orang2.ukuranSepatu = 41.5

	//mengakses nilai pada struct
	cetakOrang(orang1)
	cetakOrang(orang2)
}
