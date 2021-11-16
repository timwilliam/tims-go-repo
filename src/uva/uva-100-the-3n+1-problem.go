/*
	Problem		: uva 100 The 3n+1 problem
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

func algo(n int) (cycleCount int) {
	cycleCount = 1

	for {
		// fmt.Printf("%d ", n)

		if n == 1 {
			break
		}

		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}

		cycleCount++
	}

	// fmt.Println()
	return cycleCount
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		inputString := scanner.Text()
		inputSlice := strings.Split(inputString, " ")

		cycleLength := 1
		maxCycleLength := 1

		i, _ := strconv.Atoi(inputSlice[0])
		j, _ := strconv.Atoi(inputSlice[1])

		// swap if i with j if i > j
		if i > j {
			temp := i
			i = j
			j = temp
		}

		for idx := i; idx <= j; idx++ {
			cycleLength = algo(idx)

			if maxCycleLength <= cycleLength {
				maxCycleLength = cycleLength
			}
		}

		fmt.Printf("%d %d %d\n", i, j, maxCycleLength)
	}
}
