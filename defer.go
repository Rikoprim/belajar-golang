package main

import "fmt"
// penjadwalan function setelah function diesekusi walaupun function sebelumnya error
func logging() {
	fmt.Println("Selesai Memanggil Function")
}

func runApplication(value int) {
	defer logging() // defer selalu diatas
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("result", result)
}

func main() {
	runApplication(0)
}