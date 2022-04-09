package main

import "fmt"

func main() {
	//sebuah function/blockspok yang dapet berinteraksi dengan value lainnya
	counter := 0
	increment := func() {
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
}