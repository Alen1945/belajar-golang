package main

import (
	"fmt"
	"strconv"
)

var firstname = "Dwi Alendra Ipang"
var lastName string

// const untuk variable konstan
const tahunLahir = 1997

func main() {
	lastName = "Selly"
	/*
		operator := untuk deklarasi variable didalam function
	*/
	umur := 23
	fmt.Println("hello perkenalkan nama saya " + firstname + " " + lastName)
	fmt.Println("saya berumur", umur, "tahun")
	fmt.Println("saya kelahiran " + strconv.Itoa(tahunLahir))
}
