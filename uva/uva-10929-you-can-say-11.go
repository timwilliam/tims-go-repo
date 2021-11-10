/*
	Problem		: uva 10929 You can say 11
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

func isEleven(N int) (flag bool) {
	if N%11 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		nString := strings.TrimSpace(scanner.Text())
		fmt.Println(nString) // trim all whitespace in the input

		N, _ := strconv.Atoi(nString) // get the number N

		if N == 0 {
			break
		}

		flag := isEleven(N)

		if flag {
			fmt.Printf("%s is a multiple of 11.\n", nString)
		} else {
			fmt.Printf("%s is not a multiple of 11.\n", nString)
		}
	}
}
