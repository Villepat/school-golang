package piscine

func Compact(ptr *[]string) int {
	var arrYAY []string
	for i := 0; i < len(*ptr); i++ {
		if (*ptr)[i] != "" {
			arrYAY = append(arrYAY, (*ptr)[i])
		}
	}
	*ptr = arrYAY
	return len(arrYAY)
}
