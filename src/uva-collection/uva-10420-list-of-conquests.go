/*
	Problem		: uva 10420 List of Conquests
	Author		: timwilliam
	Compiled	: 2021/11/10
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

func main() {
	listOfConquests := make(map[string][]string)
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	conquestCountry := []string{}

	for i := 0; i < n; i++ {
		scanner.Scan()
		inputString := scanner.Text()
		inputSlice := strings.Fields(inputString) // split string around instance of consecutive white space

		// map the input
		conquestName := strings.Join(inputSlice[1:], " ")
		countryName := inputSlice[0]

		listOfConquests[countryName] = append(listOfConquests[countryName], conquestName)
	}

	// doing it here now so that we don't get duplicate country name, might not be the most efficient
	for countryName := range listOfConquests {
		conquestCountry = append(conquestCountry, countryName)
	}
	sort.Strings(conquestCountry) // sort alphabetically

	for _, country := range conquestCountry {
		fmt.Printf("%s %d\n", country, len(listOfConquests[country]))
	}
}
