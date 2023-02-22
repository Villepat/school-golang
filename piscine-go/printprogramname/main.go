package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	getname := os.Args[0]
	soar := []rune(getname)
	for idx, char := range soar {
		if idx > 1 {
			z01.PrintRune(rune(char))
		}
	}
	z01.PrintRune('\n')
}
