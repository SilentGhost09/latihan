package main

/*
Nama : Reza Hidayatulloh
Nim : 3420210019
Prodi : Teknik Informatika
*/
import (
	"fmt"

	"github.com/reza/praktikum/go_project2/latihan10/fungsi"
)

func main() {
	fungsi.Garis()

	//mengubah elemen dari slice
	harga := []int{10000, 20000, 30000}
	fmt.Println("isi dari harga adalah : ", harga)
	harga[2] = 50000
	fmt.Println("isi dari harga sekarang adalah : ", harga)
	fungsi.Garis()

	//menambahkan elemen dari slice
	//syntax --> slice_name = append(slice_name, elemen1, elemen2, ..)
	dataAngka := make([]int, 5, 10)
	dataAngka[0] = 7
	dataAngka[1] = 5
	dataAngka[2] = 9
	fmt.Printf("isi slice dataAngka adalah : %v\n", dataAngka)
	fmt.Printf("panjang slice dataAngka adalah : %d\n", len(dataAngka))
	fmt.Printf("kapasitas slice dataAngka adalah : %d\n", cap(dataAngka))
	fungsi.Garis()

	dataAngka = append(dataAngka, 2, 8)
	fmt.Printf("isi slice dataAngka adalah : %v\n", dataAngka)
	fmt.Printf("panjang slice dataAngka adalah : %d\n", len(dataAngka))
	fungsi.Garis()
}
