package main

import (
	"fmt"
	"strings"
)

func main() {
	test := "This is a **test string"

	found := strings.Split(test, "**")
	for i, string := range found {
		fmt.Printf("%d: %s\n", i+1, string)
	}
}
