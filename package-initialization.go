package main

import (
	"workspace/package/database"
	"fmt"
)

// _ untuk mematikan function yg tidak dipanggil

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}