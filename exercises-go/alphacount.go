package piscine

func AlphaCount(s string) int {
	counter := 0
	for _, character := range []byte(s) {
		if character >= 'a' && character <= 'z' {
			counter += 1
		}
	}
	for _, character := range []byte(s) {
		if character >= 'A' && character <= 'Z' {
			counter += 1
		}
	}
	return counter
}
