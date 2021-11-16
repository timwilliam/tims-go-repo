/*
	Problem		: uva 11332 Summing Digits
	Author		: timwilliam
	Compiled	: 2021/11/12
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// sum each digit in number N
func sumDigits(N int) (sum int) {
	sum = 0

	for {
		if N == 0 {
			break
		}

		sum += N % 10
		N /= 10
	}

	return sum
}

// get number of digits in number N
func getNumDigits(N int) (nDigit int) {
	return len(strconv.Itoa(N))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var numDigits int

	for {
		scanner.Scan()
		N, _ := strconv.Atoi(scanner.Text())

		if N == 0 {
			break
		}

		numDigits = getNumDigits(N)
		for {
			if numDigits > 1 {
				N = sumDigits(N)
			} else {
				fmt.Println(N)
				break
			}

			numDigits = getNumDigits(N)
		}
	}
}
