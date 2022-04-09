package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke", counter)
	// 	counter++
	// }

	//for loop dengan statement (lebih simple)
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	//perulangan dengan slice
	slice := []string{"Eko", "Mulai", "Finish"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//perulangan dengan range
	// for i, value := range slice {
	// 	fmt.Println("index", i, "=", value)
	// }

	for _, value := range slice { //_untuk menghilangkan variable yg tidak digunakan
		fmt.Println(value)
	}

	//perulangan dengan make
	person := make(map[string]string)

	person["name"] = "Eko"
	person["title"] ="Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}