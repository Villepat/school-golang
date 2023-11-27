package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb < 0 || nb > 63 {
		result = 0
	} else {
		for i := 1; i <= nb; i++ {
			result *= i
		}
	}
	return result
}
