/*
	Problem		: uva 10252 Common Permutation
	Author		: timwilliam
	Compiled	: 2021/11/14
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func getCharOccurence(stringInput string) (charOccurence map[string]int) {
	// initialize the map
	charOccurence = make(map[string]int)

	for i := 0; i < len(stringInput); i++ {
		charOccurence[string(stringInput[i])]++
	}

	// e.g.: for the string 'abb' charOccurence is ['a':a; 'b':2]
	return charOccurence
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		stringA := scanner.Text()

		scanner.Scan()
		stringB := scanner.Text()

		charOccurenceA := getCharOccurence(stringA)
		charOccurenceB := getCharOccurence(stringB)

		// get char that is common for both stringA, and string B
		var commonPermutation []string
		for char, _ := range charOccurenceA {
			if _, exist := charOccurenceB[char]; exist {
				commonPermutation = append(commonPermutation, char)
			} else {
				continue
			}
		}

		// sort the charOccurence slice
		sort.Sort(sort.StringSlice(commonPermutation))

		// print the common char based on its number of occurences and alphabetical order
		for i := 0; i < len(commonPermutation); i++ {
			for j := 0; j < charOccurenceA[commonPermutation[i]]; j++ {
				fmt.Printf("%s", commonPermutation[i])
			}
		}
		fmt.Println()
	}
}
