package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string

	s = "hello"
	var c = 'x'

	var sb strings.Builder
	sb.WriteString(s)
	sb.WriteRune(c)
	fmt.Println(sb.String())
}
