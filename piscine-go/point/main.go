package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func PrintNbrInOrder(n int) {
	num := n
	storageslice := []rune{} // curlys for empty slices
	for num > 0 {
		storageslice = append(storageslice, rune(num%10))
		num /= 10
	}
	if len(storageslice) == 0 {
		z01.PrintRune('0')
	}
	for i := len(storageslice) - 1; i >= 0; i-- {
		z01.PrintRune(storageslice[i] + 48)
	}
}

func PrintStr(s string) {
	x := []rune(s)
	for _, character := range x {
		z01.PrintRune(character)
	}
}

func main() {
	points := &point{}

	setPoint(points)

	PrintStr("x = ")
	PrintNbrInOrder(points.x)
	PrintStr(", y = ")
	PrintNbrInOrder(points.y)
	z01.PrintRune('\n')
}
