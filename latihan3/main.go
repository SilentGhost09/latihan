package main

/*
Nama : Reza Hidayatulloh
Nim : 3420210019
Prodi : Teknik Informatika
*/
import (
	"fmt"
	"math"
)

func main() {
	//go multiple variable declaration
	var NilaiA, NilaiB float64 = 2.7, 2.7

	//pemberian nilai dengan tanda ":="
	//pemberian nilai 16.0 agar value yang diberikan tidak dianggap int
	NilaiC := 16.0
	fmt.Println("nilai A adalah: ", NilaiA)
	fmt.Println("nilai floor A adalah: ", math.Floor(NilaiA))

	fmt.Println("nilai B adalah: ", NilaiB)
	fmt.Println("nilai floor B adalah: ", math.Ceil(NilaiB))

	fmt.Println("nilai C adalah: ", NilaiC)
	fmt.Println("nilai floor c adalah: ", math.Sqrt(NilaiC))
}
