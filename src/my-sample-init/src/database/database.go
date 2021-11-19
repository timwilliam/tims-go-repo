package database

import "fmt"

var connection string

// package initialization
func init() {
	fmt.Println("Connection initialized and set!")
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
