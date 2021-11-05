/*
	Problem		: uva 10222 Decode the Mad man
	Author		: timwilliam
	Compiled	: 2021/11/05

	Test Data (--> : input, <-- : output):
	--> KT''[ R[Y'G
	<-- hello world
	--> .d;t FoYt d]]'T UkDU Md,d,d Bd, gtbPgt T't]kD,U u[ HpJKU RpUK jPYdHHT Kdnp,j p.][fFPM't l[;T 'dUT .[NP,j ,[yuk []][F
	<-- make sure apple that banana can decide elephant to fight with giraffe having impossible joke late moving north oppos
	--> nDBBP,T N[.Pu RkTt'BKpDY  bDYgP[nDfBO'dy b[y[,Dyi JdfUY[P,Utfup,d' p,UTJo.T,u .oFbd'odY ,tyn[oF fIFut. ,tYo['[jPbd' yTF]PYdu[YTi F;T'Tud'  FT,F[YI oyP,DYi
	<-- vaccine vomit wheelchiar  cardiovascular coronary gastrointestinal integument muscaluar nervous system neruological respiratorey skeletal  sensory urinary
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	characters := "qwertyuiop[]\\asdfghjkl;'zxcvbnm,./"
	shift := 2 // 2 key immediately to its left

	// get N from user
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	for n := 1; n <= N; n++ {
		// get user input and convert to lower case
		scanner.Scan()
		inputString := scanner.Text()
		inputString = strings.ToLower(inputString)

		var outputString strings.Builder

		for i := 0; i < len(inputString); i++ {
			// ignore space
			if inputString[i] == ' ' {
				outputString.WriteString(" ")
				continue
			}

			// find the index in characters and write result to outputString
			charPos := strings.Index(characters, string(inputString[i]))
			outputString.WriteString(string(characters[charPos-shift]))
		}

		// print outputString
		fmt.Println(outputString.String())
		outputString.Reset()
	}
}
