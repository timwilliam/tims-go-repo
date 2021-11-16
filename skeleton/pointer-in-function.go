package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeAddress(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{"Medan", "Sumut", ""}

	fmt.Println(address)
	changeAddress(&address)
	fmt.Println(address)
}
