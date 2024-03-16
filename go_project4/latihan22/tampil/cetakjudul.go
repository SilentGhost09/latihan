package tampil

import (
	"fmt"

	"github.com/reza/praktikum/go_project2/latihan10/fungsi"
)

func Cetakjudul(judul string) {
	fungsi.Garis()
	fmt.Println("\t", judul)
	fungsi.Garis()
}
