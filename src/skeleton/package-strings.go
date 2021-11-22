package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("tim william", "tim"))
	fmt.Println(strings.Split("tim william", " "))

	fmt.Println(strings.ToLower("Tim William"))
	fmt.Println(strings.ToUpper("tim william"))
	fmt.Println(strings.ToTitle("tim william"))

	fmt.Println(strings.Trim("     tim william   ", " "))

	fmt.Println(strings.ReplaceAll("tim tim tim william", "tim", "nottim"))
}
