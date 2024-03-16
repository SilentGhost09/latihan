package main

/*
Nama : Reza Hidayatulloh
Nim : 3420210019
Prodi : Teknik Informatika
*/
import (
	"fmt"
)

/*
sebuah fungsi adalah sebuah blok dari pernyataan pernyataan yang bisa digunakan secara
berulang berulang dalam sebuah program.
sebuah fungsi tidak akan di eksekusi secara otomatis ketika program di jalankan.
sebuah fungsi akan di eksekusi ketika fungsi tersebut dipanggil.
*/

func tampilSalam(nama string) string {
	return "Assalamualaikum" + nama
}

func jumlahin(num1, num2 int) int {
	return num1 + num2
}

func main() {
	const nama string = "reza"
	const a, b int = 2, 3

	fmt.Println(tampilSalam(nama) + "apa kabar")
	fmt.Println("hasil dari a ", a, " ditambah ", b, " adalah ", jumlahin(a, b))
}
