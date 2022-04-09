package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpEko NoKTP = "2313123543453535"
	var marriedstatus Married = true
	fmt.Println(noKtpEko)
	fmt.Println(marriedstatus)
}