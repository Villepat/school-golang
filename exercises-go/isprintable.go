package piscine

func IsPrintable(s string) bool {
	counter := true
	for _, character := range []byte(s) {
		if character < 31 {
			counter = false
		}
	}
	return counter
}
