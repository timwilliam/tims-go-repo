/*
	Problem		: uva 10226 Hardwood Species
	Author		: timwilliam
	Compiled	: 2021/11/08
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	treePop := map[string]int{}
	var speciesCount, treeCount int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	T, _ := strconv.Atoi(scanner.Text())
	fmt.Println()

	for i := 1; i <= T; i++ {
		speciesCount = 0
		treeCount = 0

		// get user input and store them in the treePop map
		for {
			scanner.Scan()
			inputString := scanner.Text()

			if inputString == "" {
				break
			}

			treeCount += 1
			if count, exist := treePop[inputString]; exist {
				treePop[inputString] = count + 1
			} else {
				speciesCount += 1
				treePop[inputString] = 1
			}
		}

		// create a slice to hold all the treeSpecies and then sort the slice
		treeSpecies := make([]string, 0, speciesCount)
		for k := range treePop {
			treeSpecies = append(treeSpecies, k)
		}
		sort.Strings(treeSpecies)

		// go through each of the key in the treePop map and print out the percentage
		for _, k := range treeSpecies {
			percentage := float32(treePop[k]) / float32(treeCount) * 100.0
			fmt.Printf("%s %3.4f\n", k, percentage)

			// clear the map (might not be the best idea w/ resource utilization)
			delete(treePop, k)
		}

		fmt.Println()
	}
}
