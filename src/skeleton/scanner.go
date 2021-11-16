package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// get N from user
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	fmt.Println(N)

	// get inputString from user
	scanner.Scan()
	inputString := scanner.Text()
	fmt.Printf("user input: %q\n", inputString)
}
