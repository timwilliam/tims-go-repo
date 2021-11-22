package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Enter database hostname")
	username := flag.String("username", "root", "Enter database username")
	password := flag.String("password", "root", "Enter database password")

	flag.Parse()
	fmt.Println(*host, *username, *password)

	// to run, use the following
	// $ go run package-flag.go -host=nb-local -username=tim -password=nasigoreng
}
