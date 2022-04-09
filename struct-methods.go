package main

import "fmt"

//membuat function didalam struct
type Customer struct {
	Name, Address string
	Age			  int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func (a Customer) sayDoi() {
	fmt.Println("Hello Doi", a.Name)
}

func main() {
	var eko Customer
	eko.Name = "Eko"
	eko.Address = "Kediri"
	eko.Age = 30

	eko.sayHi("Joko")
	eko.sayDoi()
}