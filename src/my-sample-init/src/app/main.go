package main

import (
	"fmt"
	"my-sample-init/src/database"
	// can also use blank identifier like this
	// _ "my-sample-init/src/database"
	// this is so that we can run the init() function without actually having to execute any of the function in the package
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
