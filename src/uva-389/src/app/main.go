/*
	Program		: uva 389 Basically Speaking
	Author		: timwilliam
	Compiled	: 2021/11/23
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"uva-389/src/helper"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// read until EOF is reached
	for scanner.Scan() {
		rawInput := scanner.Text()
		parsedInput := helper.ParseUserInput(rawInput)

		base, _ := strconv.Atoi(parsedInput.Base)
		targetBase, _ := strconv.Atoi(parsedInput.TargetBase)

		// convert first to decimal (base 10)
		number10 := helper.ToDecimal(parsedInput.Number, base)

		// convert number10 to number at target base
		numberBase := helper.ToBase(number10, targetBase)

		fmt.Println(helper.PrintOutput(numberBase))
	}
}
