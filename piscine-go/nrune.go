package piscine

func NRune(s string, n int) rune {
	natureRune := []rune(s)
	if n <= 0 || n > len(natureRune) {
		return 0
	} else {
		return natureRune[n-1]
	}
}
