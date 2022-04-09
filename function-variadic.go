package main

import "fmt"

func sumAll(numbers ...int) int { //penjumlahan keseluruhan
	total := 0

	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	total := sumAll(10, 23, 24, 59, 50)
	fmt.Println(total)

	//variadic dengan slice
	slice := []int{10, 10, 10, 10, 10, 10}
	total = sumAll(slice...)
	fmt.Println(total)
}