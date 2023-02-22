package piscine

func ConcatParams(args []string) string {
	var sentence string
	for i := 0; i < len(args); i++ {
		if i < len(args)-1 {
			sentence += args[i]
			sentence += "\n"
		} else {
			sentence += args[i]
		}
	}
	return sentence
}
