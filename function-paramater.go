package main

import "fmt"

func sayHello(firstname string, lastname string) {
	fmt.Println("Hallo", firstname, lastname)
}

func main() {
	firstname := "Riko"
	sayHello(firstname, "Santoso")
}