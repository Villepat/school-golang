package piscine

func IsUpper(s string) bool {
	counter := true
	for _, character := range []byte(s) {
		if character < 65 || character > 90 {
			counter = false
		}
	}
	return counter
}
