package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

var sourceString = "Hello, OTUS!"

func main() {
	reversedStr := stringutil.Reverse(sourceString)
	fmt.Println(reversedStr)
}
