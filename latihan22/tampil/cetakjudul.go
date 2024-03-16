package tampil

import (
	"fmt"

	"github.com/reza/praktikum/latihan10/fungsi"
)

func Cetakjudul(judul string) {
	fungsi.Garis()
	fmt.Println("\t", judul)
	fungsi.Garis()
}
