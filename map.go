package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Eko",
		"address": "Kediri",
	}
	person["title"] = "Programmer"

	fmt.Println(person)

	//delete map
	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Riko"
	book["ups"] = "Salah"

	delete(book, "ups")
	fmt.Println(book)
}