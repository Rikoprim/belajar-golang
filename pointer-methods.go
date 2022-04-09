package main

import "fmt"

type Man struct {
	Name string
}

// * sebagai penanda pointer untuk menuju data Man
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
	fmt.Println("Di Method", man.Name)
}

func main() {
	eko := Man{"Eko"}
	eko.Married()

	fmt.Println(eko.Name)
}

