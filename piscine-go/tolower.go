package piscine

func ToLower(s string) string {
	x := []rune(s)
	temp := ""
	for _, character := range x {
		if character >= 'A' && character <= 'Z' {
			x := character + 32
			temp += string(x)
		} else {
			temp += string(character)
		}
	}
	return temp
}
