package piscine

func ActiveBits(n int) int {
	storebits := make([]int, 0)
	counter := 0
	for n > 0 {
		storebits = append(storebits, n%2)
		n /= 2
	}
	for _, ones := range storebits {
		if ones == 1 {
			counter += 1
		}
	}
	return counter
}
