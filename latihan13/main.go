package main

/*
Nama : Reza Hidayatulloh
Nim : 3420210019
Prodi : Teknik Informatika
*/
import (
	"fmt"

	"github.com/reza/praktikum/latihan10/fungsi"
)

/*
efisiensi memori
ketika menggunakan slice, go menuangkan seluruh elemen slice ke dalam memori.
jika isi data slice terlalu besar dan kita hanya perlu beberapa elemen saja maka
lebih baik kita meng copy elemen tersebut ke slice baru menggunakan copy().
ini akan mengurangi penggunaan memory pada program
*/

func main() {
	fungsi.Garis()
	//sebuah slice besar
	dataAngka := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	fmt.Printf("isi dataAngka \t :  %v\n", dataAngka)
	fmt.Printf("panjang dataAngka \t :  %d\n", len(dataAngka))
	fmt.Printf("kapasitas dataAngka \t :  %d\n", cap(dataAngka))
	fungsi.Garis()

	//copy dataAngka dari elemen ke 2 sampai panjang elemen dataAngka dikurang 10
	dataCut := dataAngka[2 : len(dataAngka)-10]
	fmt.Printf("isi dataCut \t : %v\n", dataCut)

	//buat slice dengan panjang berupa panjang dari dataCut
	dataCopy := make([]int, len(dataCut))

	//copy isi dataCut ke dataCopy
	copy(dataCopy, dataCut)

	fmt.Printf("isi dataAngka \t :  %v\n", dataCopy)
	fmt.Printf("panjang dataAngka \t :  %d\n", len(dataCopy))
	fmt.Printf("kapasitas dataAngka \t :  %d\n", cap(dataCopy))
	fungsi.Garis()
}
