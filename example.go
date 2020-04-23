package main

import (
	"fmt"
)

func printText(text string) {
	fmt.Println(text)
}
func perkalian(angka1 int, angka2 int) int {
	return (angka1 * angka2)
}
func penjumlahan(angka1 int, angka2 int) (hasil int) {
	hasil = angka1 + angka2
	return
}
func baliknilai(angka1 int, angka2 int) (int, int) {
	return angka2, angka1
}
func main() {
	printText("ze ah haha hah ahahah")
	a := 2
	b := 4
	fmt.Printf("Hasil dari %d * %d adalah %d \n", a, b, perkalian(a, b))
	fmt.Printf("Hasil dari %d + %d adalah %d \n", a, b, penjumlahan(a, b))
	a, b = baliknilai(a, b)
	fmt.Println("nilai a setelah dibalik adalah", a)
	fmt.Println("nilai b setelah dibalik", b)
}
