package main

import "fmt"

//ketika menggunakan struct yg bnyak lebih baik gunakan pointer biar tidak overload memori
type Address struct {
	City, Province, Country string
}

//pass by reference kebalikan dari pass by value
// & untuk merubah data kepointer yang akan diubah

func ChangeCountryToIndonesia(address *Address){ //pointer dengan func
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1

	//pass by value (copy)
	address2.City = "Bandung"

	*address2 = Address{"Kediri", "Jawa Timur", "Indonesia"} //membuat data baru
	//ketika dikasih * maka address akan berubah semua

	fmt.Println(address1)
	fmt.Println(address2)

	var address4 *Address = new(Address)
	address4.City = "Jakarta"
	fmt.Println(address4)

	var alamat = Address {
		City: "Jombang",
		Province: "Jawa Timur",
		Country: "",
	}

	var alamatPointer *Address = &alamat
	ChangeCountryToIndonesia(alamatPointer)
	fmt.Println(alamat)
}