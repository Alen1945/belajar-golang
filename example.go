package main

import (
	"fmt"
)

func deretAngka(batas int) (total int) {
	total = 0
	for i := 1; i <= batas; i++ {
		total = total + i
	}
	return
}
func main() {

	for i := 1; i <= 10; i++ {
		fmt.Println("ini adalah langkah ke", i)
	}
	batas := 100
	fmt.Printf("jumlah deret 1 sampai %d adalah %d \n", batas, deretAngka(batas))
}
