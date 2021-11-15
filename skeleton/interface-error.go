package main

import (
	"errors"
	"fmt"
)

func division(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("Division by ZERO!")
	} else {
		return dividend / divisor, nil
	}
}

func main() {
	var quotient int
	var errorMessage error

	quotient, errorMessage = division(6, 3)
	fmt.Println(quotient, errorMessage)

	quotient, errorMessage = division(6, 0)
	fmt.Println(quotient, errorMessage)
}
