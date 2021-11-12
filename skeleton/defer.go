package main

import "fmt"

func logging() {
	fmt.Println("App has finished running!")
}

func runApp() {
	// defer function will be executed after the calling function finishes
	// will always be executed even though an error occurs in the calling function
	defer logging()
	fmt.Println("Hello World!")
}

func main() {
	runApp()
}
