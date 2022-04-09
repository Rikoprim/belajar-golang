package main

import "fmt"

func getFullName() (string, string) {
	return "Rudi", "Santoso"
}

func main() {
	firstname, lastname, _ := getFullName() //_ untuk menghilngkan variable
	fmt.Println(firstname, lastname)
}