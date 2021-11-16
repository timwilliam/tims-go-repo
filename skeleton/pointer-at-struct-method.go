package main

import "fmt"

type Man struct {
	Name string
}

// struct method
func (man *Man) married() {
	man.Name = "Mr." + man.Name
}

func main() {
	name := Man{"Ethan Hunt"}
	name.married()

	fmt.Println(name)
}
