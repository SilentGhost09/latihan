package main

/*
Nama : Reza Hidayatulloh
Nim : 3420210019
Prodi : Teknik Informatika
*/
import (
	"fmt"
)

func main() {
	//arrays
	var namaBand = [3]string{}

	//pemberian nilai menggunakan index array
	namaBand[0] = "MUSE"
	namaBand[1] = "MY CHEMICAL ROMANCE"
	namaBand[2] = "PARAMORE"
	fmt.Printf("band-band favorit %s \n", namaBand)

	macamBuah := [...]string{"mangga", "durian", "kelengkeng", "pisang"}

	//menentukan panjang sebuah array
	fmt.Printf("jumlah elemen dalam array macamBuah adalah : %d\n", len(macamBuah))
	fmt.Println("pengambilan nilai dengan slice : ", macamBuah[1:3])

	//inisialisasi array langsung
	angka1 := [5]int{}              //belum di inisialisasi
	angka2 := [5]int{1, 2}          //di inisialisasi sebagian
	angka3 := [5]int{1, 2, 3, 4, 5} //di inisialisasi lengkap

	//cetak array
	fmt.Println(angka1)
	fmt.Println(angka2)
	fmt.Println(angka3)
}
