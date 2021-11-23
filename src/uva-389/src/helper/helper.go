package helper

import (
	"strconv"
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

func ParseUserInput(userInput string) Input {
	var parsedUserInput Input
	parsedString := strings.Split(userInput, " ")

	// e.g.: 1111000 2 10 [number, base, target base]
	parsedUserInput.Number = parsedString[0]
	parsedUserInput.Base = parsedString[1]
	parsedUserInput.TargetBase = parsedString[2]

	return parsedUserInput
}

// from base to decimal (base 10)
func ToDecimal(numberString string, base int) int64 {
	decimal, _ := strconv.ParseInt(numberString, base, 64)
	return decimal
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

	// reverse the order of the resulting string
	return ReverseString(outputString.String())
}

func ReverseString(originalString string) string {
	r := []rune(originalString)
	var reversedString []rune

	for i := len(r) - 1; i >= 0; i-- {
		reversedString = append(reversedString, r[i])
	}

	return string(reversedString)
}

func PadString(originalString string) string {
	var paddedString []rune

	// DIGIT by default is set to 7
	for i := 0; i < DIGIT; i++ {
		if i >= DIGIT-len(originalString) {
			break
		} else {
			paddedString = append(paddedString, '0')
		}
	}

	return string(paddedString) + originalString
}

func TrimToDigit(paddedString string) string {
	trimLocation := len(paddedString) - DIGIT
	return paddedString[trimLocation:]
}

func PrintOutput(numberBase string) string {
	paddedString := PadString(numberBase)

	if len(paddedString) > DIGIT {
		paddedString = TrimToDigit(paddedString)
	}

	return paddedString
}
