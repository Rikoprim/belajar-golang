package main

import "fmt"

func main() {
	//Perbandingan == > < !=
	var name1 = "Riko"
	var name2 = "Budo"

	var result bool = name1 == name2
	fmt.Println(result)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}