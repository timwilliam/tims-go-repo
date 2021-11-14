package main

import "fmt"

type Customer struct {
	Name string
}

func (customer Customer) sayHello() {
	fmt.Println("Hello", customer.Name)
}

func (customer Customer) sayBye() {
	fmt.Println("Bye", customer.Name)
}

func main() {
	customer := Customer{"tim"}

	customer.sayHello()
	customer.sayBye()
}
