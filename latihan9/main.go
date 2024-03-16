package main

/*
Nama : Reza Hidayatulloh
Nim : 3420210019
Prodi : Teknik Informatika
*/
import (
	"fmt"
)

func garis() {
	fmt.Println("=================================")
}

func main() {
	/*
		slice - mirip seperti array tetapi lebih fleksible
		panjang dari slice bisa bertembah dan berkurang
		contoh deklarasi slice_name := []datatype{values}
	*/
	garis()

	mataKuliah1 := []string{}
	fmt.Println("panjang slice mataKuliah1 adalah : ", len(mataKuliah1))
	fmt.Println("isi dari mataKuliah1 adalah :", mataKuliah1)

	garis()

	mataKuliah2 := []string{"pemrograman mobile", "web technology", "kapita selekta", "skripsi"}
	fmt.Println("panjang slice mataKuliah2 adalah : ", len(mataKuliah2))
	fmt.Println("isi dari mataKuliah2 adalah : ", mataKuliah2)

	garis()

	//membuat slice dari array
	dataNilai := [6]int{10, 11, 12, 13, 14, 15}
	sliceNilai := dataNilai[2:4]

	//%v - print the value in the default format
	fmt.Printf("nilai dari dataNilai adalah : %v\n", dataNilai)
	fmt.Printf("nilai dari sliceNilai adalah : %v\n", sliceNilai)
	fmt.Printf("kapasitas dari sliceNilai adalah : %d\n", cap(sliceNilai))

	garis()
}
