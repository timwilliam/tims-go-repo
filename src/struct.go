package main

import "fmt"

type Customer struct {
	Name, Country string
	Age           int
}

func main() {
	var sample Customer

	sample.Name = "Tim"
	sample.Age = 17
	sample.Country = "Taiwan"

	fmt.Println(sample)

	sample2 := Customer{
		Name:    "Budi",
		Age:     30,
		Country: "Indonesia",
	}
	fmt.Println(sample2)

	sample3 := Customer{"Citra", "Indonesia", 22}
	fmt.Println(sample3)
}
