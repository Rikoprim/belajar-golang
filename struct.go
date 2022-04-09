package main

import "fmt"
// Deklarasi field
type Customer struct {
	Name, Address string
	Age			  int
}

func main() {
	var eko Customer
	eko.Name = "Eko"
	eko.Address = "Kediri"
	eko.Age = 30

	fmt.Println(eko)
	fmt.Println(eko.Age)

	// struct literal
	joko := Customer{
		Name: "Joko",
		Address: "Malang",
		Age: 25,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Jakarta", 25}
	fmt.Println(budi)
}