package main

import "fmt"

func main() {
	// continue menghentikan perulangan dan menyeleksi dahulu lalu dilanjutkan
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Println("Perulangan ke ", i)
	}
}