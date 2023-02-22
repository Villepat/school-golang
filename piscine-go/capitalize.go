package piscine

func Capitalize(s string) string {
	slow := ToLower(s)
	var str string
	a := []rune(slow)
	if a[0] >= 'a' && a[0] <= 'z' {
		a[0] -= 32
	}
	for i := 0; i < len(a); i++ {
		if i < len(a)-1 && (a[i] < '0' || a[i] > '9') && (a[i] < 'A' || a[i] > 'Z') && (a[i] < 'a' || a[i] > 'z') {
			if a[i+1] >= 'a' && a[i+1] <= 'z' {
				a[i+1] -= 32
			}
		}
		str += string(a[i])
	}
	return str
}
