package main

/*
Nama : Reza Hidayatulloh
Nim : 3420210019
Prodi : Teknik Informatika
*/
import (
	"fmt"

	"github.com/reza/praktikum/latihan4/aritmatik"
)

func main() {
	/* var
	-bisa digunakan diluar ataupun didalam fungsi
	-deklarasi variable dan pemberian nilai bisa dilakukan secara terpisah

	:=
	-hanya bisa digunakan didalam sebuah fungsi
	-deklarasi variable dan pemberian nilai tidak bisa dilakukan secara terpisah
	(harus dilakukan dalam baris yang sama)
	*/

	//go multiple variable declaration
	NilaiA, NilaiB := 2, 3
	var testJumlah int = aritmatik.Penjumlahan(NilaiA, NilaiB)

	fmt.Println("nilai A adalah:", NilaiA)
	fmt.Println("nilai B adalah:", NilaiB)

	fmt.Println("nilai A ditambah nilai B adalah", testJumlah)
}
