/*
	Problem		: uva 10035 Primary Arithmetic
	Author		: timwilliam
	Compiled	: 2021/11/10
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func storeToIntSlice(numberString string) []int {
	numberInt, _ := strconv.Atoi(numberString)
	numberLen := len(numberString)
	numberSlice := make([]int, numberLen)

	// store number into slice, in reverse order
	// e.g.: 123 --> [3 2 1]
	for i := 0; i < numberLen; i++ {
		numberSlice[i] = numberInt % 10
		numberInt /= 10
	}

	return numberSlice
}

func sumAndGetCarryCount(numberOne []int, numberTwo []int) int {
	carryCount := 0

	maxNumberLen := len(numberOne)
	if maxNumberLen < len(numberTwo) {
		maxNumberLen = len(numberTwo)
	}

	sum := make([]int, maxNumberLen+1)
	carry := 0

	for i := 0; i <= maxNumberLen; i++ {
		if i < len(numberOne) && i < len(numberTwo) {
			sum[i] = numberOne[i] + numberTwo[i] + carry
		} else if i >= len(numberOne) && i != maxNumberLen {
			sum[i] = numberTwo[i] + carry
		} else if i >= len(numberTwo) && i != maxNumberLen {
			sum[i] = numberOne[i] + carry
		} else {
			sum[i] = carry
		}

		if sum[i] >= 10 {
			sum[i] -= 10
			carry = 1
			carryCount++
		} else {
			carry = 0
		}

	}

	// fmt.Println(sum)
	return carryCount
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		input := scanner.Text()
		inputSlice := strings.Split(input, " ")

		if inputSlice[0] == "0" && inputSlice[1] == "0" {
			break
		}

		numberOne := storeToIntSlice(inputSlice[0])
		numberTwo := storeToIntSlice(inputSlice[1])
		carryCount := sumAndGetCarryCount(numberOne, numberTwo)

		if carryCount == 0 {
			fmt.Println("No carry operation.")
		} else if carryCount == 1 {
			fmt.Println("1 carry operation.")
		} else {
			fmt.Printf("%d carry operations.\n", carryCount)
		}
	}
}
