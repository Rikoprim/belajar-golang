package main

import "fmt"

func main() {
	name := "JokoJokoJokoJoko" 
	
	switch name {
	case "Eko":
		fmt.Println("Hello EKo")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Hi, Boleh Kenalan")
	}

	//switch (short statement)
	switch length := len(name); length > 5 {
	case true: 
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}

	//switch tanpa kondisi
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}


}