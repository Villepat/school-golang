package piscine

func IsNumeric(s string) bool {
	counter := true
	for _, character := range []byte(s) {
		if character < 48 || character > 57 {
			counter = false
		}
	}
	return counter
}
