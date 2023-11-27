package piscine

func MakeRange(min, max int) []int {
	lele := max - min
	if min >= max {
		return nil
	}
	returnvalue := make([]int, lele)

	if min < max {
		for i := min; i < max; i++ {
			if i >= min {
				returnvalue[i-min] = i
			}
		}
	}
	return returnvalue
}
