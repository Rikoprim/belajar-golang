package main

import "fmt"

//function untuk menghentikan sebuah function ketika sistem error

func endApp() {
	message := recover() //recover disimpan di defer function
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
	fmt.Println("Aplikasi Selesai")
}

// recover untuk menangkap data panic
func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi Error")
	}

	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
}