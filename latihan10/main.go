package main

/*
Nama : Reza Hidayatulloh
Nim : 3420210019
Prodi : Teknik Informatika
*/
import (
	"fmt"

	"github.com/reza/praktikum/latihan10/fungsi"
)

func main() {
	//fungsi garis yang di buat dalam package
	fungsi.Garis()

	fmt.Sprintln("contoh membuat slice dengan make")
	dataNilai1 := make([]int, 5, 10)
	fmt.Printf("isi dari dataNilai : %v\n", dataNilai1)
	fmt.Printf("panjang dari dataNilai1 : %d\n", len(dataNilai1))
	fmt.Printf("kapasitas dari dataNilai1 : %d\n", cap(dataNilai1))

	fungsi.Garis()

	//jika parameter kapasitas tidak di definisikan maka kapasitas
	//akan sama dengan panjang dari slice
	dataNilai2 := make([]int, 5)
	fmt.Printf("isi dari dataNilai2 : %v\n", dataNilai2)
	fmt.Printf("panjang dari dataNilai2 : %d\n", len(dataNilai2))
	fmt.Printf("kapasitas dari dataNilai2 : %d\n", cap(dataNilai2))

	fungsi.Garis()
}
