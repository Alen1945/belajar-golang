package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 5
	var c int
	c = a + b
	fmt.Printf("%d + %d = %d \n", a, b, c)

	c = a - b
	fmt.Printf("%d - %d = %d \n", a, b, c)

	c = a * b
	fmt.Printf("%d * %d = %d \n", a, b, c)

	c = a / b
	fmt.Printf("%d : %d = %d\n", a, b, c)

	c = a % b
	fmt.Printf("Sisa Pembagian dari %d dan %d adalah %d\n", a, b, c)
}
