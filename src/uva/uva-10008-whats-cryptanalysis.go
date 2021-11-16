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
	scanner := bufio.NewScanner(os.Stdin)

	// get N from user
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	// character map
	charMap := make(map[string]int)

	for i := 0; i < n; i++ {
		scanner.Scan()
		inputString := strings.ToUpper(scanner.Text())

		for i := 0; i < len(inputString); i++ {
			if inputString[i] >= 'A' && inputString[i] <= 'Z' {
				charMap[string(inputString[i])]++
			} else {
				continue
			}
		}
	}

	// copy map into slice for sorting
	type keyValue struct {
		char  string
		count int
	}

	var kv []keyValue
	for char, count := range charMap {
		kv = append(kv, keyValue{char, count})
	}

	// sort the charMap based on count, and then alphabetically
	sort.Slice(kv, func(i, j int) bool {
		if kv[i].count > kv[j].count {
			return true
		} else if kv[i].count < kv[j].count {
			return false
		}
		return kv[i].char < kv[j].char
	})

	for i := 0; i < len(kv); i++ {
		fmt.Printf("%s %d\n", kv[i].char, kv[i].count)
	}
}
