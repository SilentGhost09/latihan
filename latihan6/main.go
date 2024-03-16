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
	//array
	/*
		cara mendeklarasikan array pertama

		var array_name = [length]datetype{value} --> panjang array di definisikan
		var array_name = [...]datetype{value} --> panjang array bebas

		cara mendeklarasikan array kedua

		array_name := [length]datetype{value} --> panjang array di definisikan
		array_name := [...]datetype{value} --> panjang array bebas
	*/
	var buah = [2]string{"anggur", "nanas"}
	fmt.Printf("buah yang pertama adalah %s dan buah yang kedua adalah %s \n", buah[0], buah[1])

	/*
		integer formating verbs = "%d" = base 10
		string formating verbs = "%s" = print the value as plain string
		boolean formating verbs = "%t" = value of the boolean oprator in true or false format
		float formating verbs = "%d" = default width, precision 2
	*/
}
