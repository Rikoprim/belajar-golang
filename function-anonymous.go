package main

import "fmt"

type Blacklistku func(string) bool

func registerUser(name string, blacklist Blacklistku) {
	if blacklist(name) {
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool { //function dimasukkan kedalam variabel
		return name == "admin"
	}
	registerUser("admin", blacklist)
	registerUser("eko", blacklist)
}