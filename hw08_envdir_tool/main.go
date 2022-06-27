package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprint(os.Stdout, "wrong arguments count")
		os.Exit(1)
	}
	dir := os.Args[1]
	cmd := os.Args[2:]

	environments, err := ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	code := RunCmd(cmd, environments)
	os.Exit(code)
}
