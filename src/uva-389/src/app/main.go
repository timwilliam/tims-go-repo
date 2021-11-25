/*
	Program		: uva 389 Basically Speaking
	Author		: timwilliam
	Compiled	: 2021/11/25
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"uva-389/src/helper"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// read until EOF is reached
	for scanner.Scan() {
		parsedInput := strings.Fields(scanner.Text())

		numberString := parsedInput[0]
		base, _ := strconv.Atoi(parsedInput[1])
		targetBase, _ := strconv.Atoi(parsedInput[2])

		// convert first to base 10 (decimal)
		numberBaseTen, _ := strconv.ParseInt(numberString, base, 64)

		// convert numberBaseTen to targetBase
		numberBaseX := helper.ToBase(numberBaseTen, targetBase)

		// print the results
		fmt.Printf("%7s\n", numberBaseX)
	}
}
