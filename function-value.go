package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	goodbye := getGoodBye
	fmt.Println(goodbye("Luki"))
	result := goodbye("Rudi")
	fmt.Println(result)
}