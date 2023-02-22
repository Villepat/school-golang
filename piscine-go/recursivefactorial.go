package piscine

func RecursiveFactorial(nb int) int {
	if nb > 63 || nb < 0 {
		return 0
	}

	if nb >= 1 && nb < 64 {
		return nb * RecursiveFactorial(nb-1)
	} else {
		return 1
	}
}
