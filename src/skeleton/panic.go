package main

import "fmt"

func endApp() {
	fmt.Println("Application finishes!")
}

func runApp(error bool) {
	defer endApp()

	if error {
		// panic can be used to stop a running program
		// panic is usually called when there is an error in the program
		// when panic is called, program will stop, but defer function will still be executed
		panic("ERROR OCCURS!")
	}

	fmt.Println("Application is running ...")
}

func main() {
	// runApp(true)
	runApp(false)
}
