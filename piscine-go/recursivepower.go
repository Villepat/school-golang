package piscine

func RecursivePower(nb int, power int) int {
	var result int = 1
	if power < 0 {
		return 0
	} else if power > 0 {
		result = nb * RecursivePower(nb, power-1)
	}
	return result
}
