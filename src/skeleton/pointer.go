package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Medan", "Sumatera_Utara", "Indonesia"}
	address2 := &address1
	address3 := &address1
	// var address3 *Address = &address1

	fmt.Println(address1)

	address2.City = "Binjai"

	*address2 = Address{"Malang", "Jawa_Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	address4 := new(Address)
	address5 := address4

	fmt.Println(address4)

	address5.Country = "Taiwan"

	fmt.Println(address4)
	fmt.Println(address5)
}
