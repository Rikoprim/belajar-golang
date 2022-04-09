package main

import "fmt"

func main() {
	var name = "Budi"

	if name == "Eko" {
		fmt.Println("Hello Eko")
	} else {
		fmt.Println("Hi, Kenalan Yuks")
	}

	//if panjang huruf (short statement)

	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}