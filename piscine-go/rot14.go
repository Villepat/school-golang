package piscine

// from A to L it is possible to increase the size by 14
// from M to Z we have to -12 because increasing the size would give us wrong characters
func rotta(b rune) rune {
	if b >= 'A' && b < 'M' || b >= 'a' && b < 'm' {
		return b + 14
	}
	if b >= 'M' && b <= 'Z' || b >= 'm' && b <= 'z' {
		return b - 12
	}
	return b
}

func Rot14(str string) string {
	result := ""
	for _, res := range str { // range over the input string, convert the runes to string and add them to the result string
		result += string(rotta(res))
	}
	return result
}
