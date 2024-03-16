package main

import (
	"fmt"

	"github.com/reza/praktikum/latihan10/fungsi"
)

/*
Nama : Reza Hidayatulloh
Nim : 3420210019
Prodi : Teknik Informatika
*/

func main() {
	fungsi.Garis()
	//Contoh Looping Dengan Conditions
	for i := 0; i < 5; i++ {
		if i == 4 {
			fmt.Print(i, "\n")
		} else {
			fmt.Print(i, ", ")
		}
	}
	fungsi.Garis()
	//Contoh Looping Dengan Continue
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print(i, " ")
		} else {
			continue
		}

	}
	fmt.Print("\n")
	fungsi.Garis()
	//Contoh Looping Dengan Break
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
		if i == 5 {
			break
		}
	}
	fmt.Print("\n")
	fungsi.Garis()
}
