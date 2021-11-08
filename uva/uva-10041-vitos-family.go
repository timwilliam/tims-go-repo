package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// get N from user
	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())

	for i := 1; i <= T; i++ {

	}
}
