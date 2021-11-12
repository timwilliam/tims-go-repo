package main

import "fmt"

func endApp() {
	fmt.Println("Application finishes!")

	// recover needs to be executed last
	// recover needs to be placed in defer function
	message := recover()
	fmt.Println("Error occured with message:", message)
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("ERROR OCCURS!")
	}
}

func main() {
	// runApp(true)
	runApp(true)
}
