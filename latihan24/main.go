package main

import (
	"fmt"

	"github.com/reza/praktikum/latihan10/fungsi"
	"github.com/reza/praktikum/latihan22/tampil"
)

func main() {
	tampil.Cetakjudul("Program Maps")
	var motor = make(map[string]string)
	motor["merk"] = "Honda"
	motor["model"] = "Vario"
	motor["tahun"] = "2018"
	fmt.Println("data maps  motor adalah", motor)
	fungsi.Garis()

	//mengubah elemen tahun
	motor["tahun"] = "2020"
	//mengubah elemen warna
	motor["warna"] = "Hitam"
	fmt.Println("data maps motor sekarang adalah", motor)

	//menghapus elemen model
	delete(motor, "model")
	fmt.Println("data maps motor sekarang adalah", motor)
	fungsi.Garis()

	motor["cc"] = ""
	nilai1, status1 := motor["merk"]  //cek dari key yang ada dan nilainya
	nilai2, status2 := motor["model"] //cek dari key yang tidak ada dan nilainya
	nilai3, status3 := motor["cc"]    //cek dari key yang ada dan nilainya
	_, status4 := motor["tahun"]      //hanya mengecek statusnya saja

	fmt.Println("nilai1 adalah", nilai1, "dan statusnya adalah", status1)
	fmt.Println("nilai2 adalah", nilai2, "dan statusnya adalah", status2)
	fmt.Println("nilai3 adalah", nilai3, "dan statusnya adalah", status3)
	fmt.Println("nilai4 statusnya adalah", status4)
}
