package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	num := n
	storageslice := []rune{} // curlys for empty slices
	for num > 0 {
		storageslice = append(storageslice, rune(num%10))
		num /= 10
	}
	for i := 0; i < len(storageslice); i++ {
		for j := 0; j < len(storageslice); j++ {
			if storageslice[i] > storageslice[j] {
				temp := storageslice[i]
				storageslice[i] = storageslice[j]
				storageslice[j] = temp
			}
		}
	}
	if len(storageslice) == 0 {
		z01.PrintRune('0')
	}
	for i := len(storageslice) - 1; i >= 0; i-- {
		z01.PrintRune(storageslice[i] + 48)
	}
}
