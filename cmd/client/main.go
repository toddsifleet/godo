package main

import (
	"os"
	"strings"

	"github.com/toddsifleet/godo/client"
)

func main() {
	if len(os.Args) < 2 {
		panic("NOT ENOUGH ARGUMENTS")
	}

	currentDirectory, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	searchType := os.Args[1]
	c, err := client.New(":1234", currentDirectory, searchType)
	if err != nil {
		panic(err)
	}
	searchString := ""
	if len(os.Args) > 2 {
		searchString = strings.Join(os.Args[2:], " ")
	}

	if err := c.Run(searchString); err != nil {
		panic(err)
	}
}
