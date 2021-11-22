package main

import (
	"fmt"
	"regexp"
)

// utility to find regular expression, made based on Google RE2

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("e([a-z])o")

	fmt.Println(regex.MatchString("tim"))
	fmt.Println(regex.MatchString("Tim"))

	search := regex.FindAllString("andi budi citra tim", 1)
	fmt.Println(search)
}
