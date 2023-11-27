package piscine

func SplitWhiteSpaces(s string) []string {
	var a []string
	var b string

	for _, c := range s {
		if c != ' ' && c != '\t' && c != '\n' {
			b += string(c)
		} else if c == ' ' || c == '\t' || c == '\n' {
			if b != "" {
				a = append(a, b)
				b = ""
			}
		}
	}
	if b != "" {
		a = append(a, b)
	}
	return a
}
