package main

import "fmt"

func random() interface{} {
	// return "Oops!"
	// return 100
	return true
}

func main() {
	var result interface{} = random()
	// var resultString string = result.(string)

	switch value := result.(type) {
	case int:
		fmt.Println("Value", value, "is int")
	case string:
		fmt.Println("Value", value, "is string")
	default:
		fmt.Println("Unknown type")
	}

}
