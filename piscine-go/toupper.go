package piscine

func ToUpper(s string) string {
	x := []rune(s)
	temp := ""
	for _, character := range x {
		if character >= 'a' && character <= 'z' {
			x := character - 32
			temp += string(x)
		} else {
			temp += string(character)
		}
	}
	return temp
}
