package main

/*
Nama : Reza Hidayatulloh
Nim : 3420210019
Prodi : Teknik Informatika
*/
import "fmt"

func main() {
	var nama string = "reza"
	var umur int = 22
	var ukuranSepatu float32 = 42.5

	//const constName type = value
	const isPria bool = true

	fmt.Println("nama saya adalah: ", nama)
	fmt.Println("umur saya adalah: ", umur)
	//"\t artinya mencetak  karakter tab"
	fmt.Println("ukuran sepatu saya adalah \t\t\t:", ukuranSepatu)
	fmt.Println("jenis kelamin saya adalah (true = pria) \t:", isPria)
}
