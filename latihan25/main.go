package main

import (
	"fmt"

	"github.com/reza/praktikum/latihan10/fungsi"
	"github.com/reza/praktikum/latihan22/tampil"
)

func main() {
	tampil.Cetakjudul("Program Maps")
	var motor1 = map[string]string{"merk": "Honda", "model": "Vario", "tahun": "2018"}
	//maps references
	//jika sebuah map mereferensikan ke map yang lain maka mengubah nilai map
	//tersebut akan mengubah nilai map asalnya.
	motor2 := motor1

	fmt.Println("data motor1 adalah ", motor1)
	fmt.Println("data motor2 adalah ", motor2)

	motor2["tahun"] = "2020"
	fungsi.Garis()

	fmt.Println("data motor1 sekarang adalah ", motor1)
	fmt.Println("data motor2 sekarang adalah ", motor2)
	fungsi.Garis()

	//looping di dalam map
	for kunci, nilai := range motor1 {
		fmt.Printf("%v:%v, ", kunci, nilai)
	}
	fmt.Print("\n")
}
