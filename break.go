package main

import "fmt"

func main() {
	// break untuk menghentikan seluruh perulangan
	for i := 0; i < 10; i++ {
		fmt.Println("Perulangan ke ", i)
		if i == 5 {
			break
		}
	}
}