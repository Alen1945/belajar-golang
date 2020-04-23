package main

import (
	"fmt"
)

func ubahNilai(angka *int) {
	*angka = 2000
	fmt.Println(*angka)
}

func main() {
	nilai := 1000
	//ubah nilai diluar scope
	ubahNilai(&nilai)
	fmt.Println("Nilainya sekarang", nilai)
}
