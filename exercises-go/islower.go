package piscine

func IsLower(s string) bool {
	counter := true
	for _, character := range []byte(s) {
		if character < 97 || character > 122 { // decimal ascii
			counter = false
		}
	}
	return counter
}
