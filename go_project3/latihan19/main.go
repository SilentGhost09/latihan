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

// Go Function Multiple Return Values
func LuasDanKellPersegi(panjang, lebar int) (luas int, kell int) {
	luas = panjang * lebar
	kell = (2 * panjang) + (2 * lebar)
	return
}

func CetakJudul() {
	fungsi.Garis()
	fmt.Println("Program Menghitung Luas Dan Keliling Persegi")
	fungsi.Garis()
}

func main() {
	CetakJudul()
	panjangPersegi, lebarPersegi := 10, 4
	fmt.Println("Panjang persegi adalah \t\t:", panjangPersegi)
	fmt.Println("Lebar persegi adalah \t\t:", lebarPersegi)

	//Buat variable sekaligus 2 untuk menampung return value dari fungsi yang juga ada 2
	var luasPersegi, kellPersegi int = LuasDanKellPersegi(panjangPersegi, lebarPersegi)
	fmt.Println("Luas persegi adalah \t\t:", luasPersegi)
	fmt.Println("Keliling persegi adalah \t:", kellPersegi)
	fungsi.Garis()
}
