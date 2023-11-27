package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	for _, name := range arguments {
		kek := []rune(name)
		for _, ru := range kek {
			z01.PrintRune(rune(ru))
		}
		z01.PrintRune('\n')
	}
}
