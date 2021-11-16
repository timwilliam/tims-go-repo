/*
	Problem		: uva 10055 Hashmat the brave warrior
	Author		: timwilliam
	Compiled	: 2021/11/09
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
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		userInput := scanner.Text()
		numOfSoldiers := strings.Split(userInput, " ")

		hashmatArmy, _ := strconv.Atoi(numOfSoldiers[0])
		enemyArmy, _ := strconv.Atoi(numOfSoldiers[1])

		// hashmatArmy number is always lower than enemy
		if hashmatArmy > enemyArmy {
			hashmatArmy = hashmatArmy + enemyArmy
			enemyArmy = hashmatArmy - enemyArmy
			hashmatArmy = hashmatArmy - enemyArmy
		}

		fmt.Println(enemyArmy - hashmatArmy)
	}
}
