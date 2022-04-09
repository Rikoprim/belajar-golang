package main

import (
	"workspace/package/helper"
	"fmt"
)

func main() {
	helper.SayHello("Riko")
	// helper.sayGoodbye("Riko") //error karena depan kecil
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) //error karena depan kecil
}