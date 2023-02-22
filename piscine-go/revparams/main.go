package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	str := ""
	for _, firstarg := range args {
		str = firstarg + "\n" + str
	}
	for _, newArgOrder := range str {
		z01.PrintRune(newArgOrder)
	}
}
