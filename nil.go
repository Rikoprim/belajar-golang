package main

import "fmt"
// untuk inisial data kosong menggunakan nil
// hanya tipe data interface, function, map, slice, pointer, channel
func NewMap(name string) map[string] string {
	if name == "" {
		return nil
	} else {
		return map[string] string {
			"name": name,
		}
	}
}

func main() {
	var person map[string]string = NewMap("Eko")

	if person == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person)
	}
}