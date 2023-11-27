package piscine

func BasicAtoi(s string) int {
	storageint := 0
	counter := 0
	r := []rune(s)
	for _, char := range r {
		for i := 0; i < int(char)-48; i++ {
			counter++
		}
		storageint = storageint*10 + counter
		counter = 0

	}
	return storageint
}
