package piscine

func LastRune(s string) rune {
	finalRune := []rune(s)
	return finalRune[len(finalRune)-1]
}
