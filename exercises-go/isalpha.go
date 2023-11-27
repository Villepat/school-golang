package piscine

func IsAlpha(s string) bool {
	counter := true
	for _, character := range []byte(s) {
		if character < 48 || character > 57 && character < 65 || character > 90 && character < 97 || character > 122 {
			counter = false
		}
	}
	return counter
}
