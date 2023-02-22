package piscine

func Abort(a, b, c, d, e int) int {
	s := []int{a, b, c, d, e}
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	return s[len(s)/2]
}
