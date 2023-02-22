package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	x := []rune(s)
	for _, character := range x {
		z01.PrintRune(character)
	}
}
