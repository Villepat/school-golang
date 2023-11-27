package piscine

func Unmatch(a []int) int {
	for idx1, val1 := range a {
		count := 1
		for idx2, val2 := range a {
			if val1 == val2 && idx1 != idx2 { // check if there is a match in different elements of the string
				count++
			}
		}
		if count%2 == 1 {
			return val1
		}
	}
	return -1
}
