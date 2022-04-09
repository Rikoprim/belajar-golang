package main

import "fmt"

type Filterku func(string) string // untuk alias functions

func sayHelloWithFilter(name string, filter Filterku) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Eko", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
}