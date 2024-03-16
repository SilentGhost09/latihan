package main

import (
	"fmt"

	"github.com/reza/praktikum/go_project2/latihan10/fungsi"
)

/*
Nama : Reza Hidayatulloh
Nim : 3420210019
Prodi : Teknik Informatika
*/

func main() {
	fungsi.Garis()
	nilaiA, nilaiB := 20, 30
	fmt.Println("nilaiA adalah", nilaiA)
	fmt.Println("nilaiß adalah", nilaiB)
	//Jika nilaiA lebih besar dari nilaiB
	if nilaiA > nilaiB {
		fmt.Println("nilaiA lebih besar dari nilaiß.")
	} else {
		fmt.Println("nilai lebih besar dari nilaiA.")
	}
	fungsi.Garis()
	jam := 14
	fmt.Println("Sekarang adalah pukul", jam, ":00")
	if jam < 10 {
		fmt.Println("Selamat Pagi.")
	} else if jam < 15 {
		fmt.Println("Selamat Siang.")
	} else {
		fmt.Println("Selamat Malam.")
	}
	fungsi.Garis()
	if jam <= 8 { //Contoh Penggunaan Nested If
		fmt.Println("Waktunya sarapan pagi.")
	} else if jam < 15 {
		fmt.Println("Waktunya makan siang.")
		if jam == 14 {
			fmt.Println("Setelah itu bobo siang ya.")
		}
	} else {
		fmt.Println("Waktunya makan malam.")
	}
}
