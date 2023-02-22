package piscine

func TrimAtoi(s string) int {
	var num int
	soar := []rune(s)
	kek := ""
	isnega := false
	isNum := false
	for _, char := range soar {
		if char >= '0' && char <= '9' {
			kek += string(char)
			isNum = true
		} else if char == '-' && !isNum {
			isnega = true
		}
	}
	if isnega {
		for _, oo := range kek {
			num = num*10 + int(oo) - 48
		}
		return num * -1
	}
	for _, oo := range kek {
		num = num*10 + int(oo) - 48
	}
	return num
}
