package piscine

func AppendRange(min, max int) []int {
	var returnvalue []int
	for i := min; i < max; i++ {
		returnvalue = append(returnvalue, i)
	}
	return returnvalue
}
