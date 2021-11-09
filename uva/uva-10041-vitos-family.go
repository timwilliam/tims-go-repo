/*
	Problem		: uva 10041 Vito's Family
	Author		: timwilliam
	Compiled	: 2021/11/09
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func intAbs(value int) int {
	if value < 0 {
		return -value
	} else {
		return value
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// get N from user
	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())

	for t := 1; t <= T; t++ {
		// get number of relatives and street numbers
		scanner.Scan()
		inputString := scanner.Text()
		inputSlice := strings.Split(inputString, " ")

		nRelatives, _ := strconv.Atoi(inputSlice[0])
		var median, distanceSum int

		// store all the street numbers into streetNumbers slice
		streetNumbers := make([]int, nRelatives)
		for i, streetNumber := range inputSlice[1:] {
			streetNumbers[i], _ = strconv.Atoi(streetNumber)
		}

		// sort in increasing order
		sort.Ints(streetNumbers)

		// get the median
		if nRelatives%2 == 0 {
			median = nRelatives / 2
		} else {
			median = (nRelatives + 1) / 2
		}

		// calculate the minimal sum of distances
		distanceSum = 0
		for i := 0; i < nRelatives; i++ {
			distance := streetNumbers[median-1] - streetNumbers[i]
			distance = intAbs(distance)

			distanceSum += distance
		}

		// print the final result
		fmt.Println(distanceSum)
	}
}
