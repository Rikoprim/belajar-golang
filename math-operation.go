package main

import "fmt"

func main() {
	//operation penjumlahan
	var result = 10 + 10
	fmt.Println(result)

	//operation perkalian
	var a = 10
	var b = 10
	var c = a * b
	fmt.Println(c)

	//penjumlahan augmented assigments
	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)

	//penjumlahan unary operator
	i++ // i = i + 1
	fmt.Println(i)
}