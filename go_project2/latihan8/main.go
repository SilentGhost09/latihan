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
	//arrays initialize only specific element
	//pemberian nilai hanya di elemen ke 1 dan ke 3 saja
	dataNilai := [5]int{1: 98, 3: 70}
	fmt.Println("data nilai adalah : ", dataNilai)

	//perubahan nilai salah satu elemen array
	dataNilai[0] = 77
	fmt.Println("data nilai sekarang adalah : ", dataNilai)

	//salah satu cara deklarasi multiple variable
	var (
		nilaiA int = 66
		nilaiB int = 90
	)

	//memasukan nilai sebuah variable ke dalam elemen array
	dataNilai[2] = nilaiA
	dataNilai[4] = nilaiB
	fmt.Println("data nilai lengkap adalah : ", dataNilai)
}
