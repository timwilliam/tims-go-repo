package helper

import (
	"strings"
)

type Input struct {
	Number, Base, TargetBase string
}

// only display up to seven digit as per problem statement
var DIGIT int = 7

var baseMap map[int]string = map[int]string{
	0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9",
	10: "A", 11: "B", 12: "C", 13: "D", 14: "E", 15: "F",
}

// from decimal to base 2-16
func ToBase(number int64, targetBase int) string {
	var outputString strings.Builder
	var remainder int64

	// convert, and store result in outputString
	for number > 0 {
		remainder = number % int64(targetBase)
		outputString.WriteString(baseMap[int(remainder)])
		number /= int64(targetBase)
	}

	if len(outputString.String()) <= 7 {
		// reverse the order of the resulting string
		return ReverseString(outputString.String())
	} else {
		// return error if number of digit exceeds DIGIT (default = 7)
		return "ERROR"
	}

}

func ReverseString(originalString string) string {
	r := []rune(originalString)
	var reversedString []rune

	for i := len(r) - 1; i >= 0; i-- {
		reversedString = append(reversedString, r[i])
	}

	return string(reversedString)
}
