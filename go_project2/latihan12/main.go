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

	//menjumlahkan 2 buah slice
	//syntax --> slice3 = append(slice1, slice2, ...)
	dataNama1 := []string{"hafids", "abdhol", "rendy"}
	fmt.Printf("dataNama1 adalah %v\n", dataNama1)
	dataNama2 := []string{"bella", "ratu", "dita"}
	fmt.Println("dataNama2 adalah ", dataNama2)
	dataNama3 := append(dataNama1, dataNama2...)
	fmt.Print("dataNama3 adalah ", dataNama3, "\n")
	fmt.Printf("panjang slice dataNama3 adalah %d\n", len(dataNama3))
	fmt.Printf("kapasitas slice dataNama3 adalah %d\n", cap(dataNama3))
	fungsi.Garis()

	//merubah panjang dari slice
	arrayAngka := [6]int{9, 10, 11, 12, 13, 14} //sebuah array angka
	dataAngka1 := arrayAngka[1:5]               //membuat slice dari array
	fmt.Printf("isi dataAngka1 adalah %v\n", dataAngka1)
	fmt.Printf("panjang dataAngka1 adalah %d\n", len(dataAngka1))
	fmt.Printf("kapasitas dataAngka1 adalah %d\n", cap(dataAngka1))

	dataAngka1 = arrayAngka[1:3] //merubah panjang slice dengan cara di slice ulang
	fmt.Printf("isi dataAngka1 adalah %v\n", dataAngka1)
	fmt.Printf("panjang dataAngka1 adalah %d\n", len(dataAngka1))   //panjang slice berkurang
	fmt.Printf("kapasitas dataAngka1 adalah %d\n", cap(dataAngka1)) //kapasitas slice tidak berubah

	dataAngka1 = append(dataAngka1, 20, 21, 22, 23) //merubah panjang slice dengan cara menambahkan item
	fmt.Printf("isi dataAngka1 adalah %v\n", dataAngka1)
	fmt.Printf("panjang dataAngka1 adalah %d\n", len(dataAngka1))   //panjang slice bertambah
	fmt.Printf("kapasitas dataAngka1 adalah %d\n", cap(dataAngka1)) //kapasitas slice bertambah multiply by 5

}
