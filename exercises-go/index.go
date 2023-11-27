package piscine

func Index(s string, toFind string) int {
	a := []rune(s)
	if len(s) == 0 {
		return -1
	}
	if len(toFind) < 1 {
		return 0
	}
	for idx1, character1 := range s {
		if character1 == rune(toFind[0]) {
			temp := a[idx1 : idx1+len(toFind)]
			if toFind == string(temp) {
				return idx1
			}
		}
	}
	return -1
}
