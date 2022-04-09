package main

import "fmt"

func getFullName() (firstname string, middlename string, lastname string) {
	firstname = "Rudi"
	middlename = "Maria"
	lastname = "Diba"

	return
}

func main() {
	a, b, c := getFullName()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}