package ascii

import (
	"strings"
)

/*
Takes input argument string and matches key of character by casting it into an integer and appending the equivalent
ascii-art character into a slice of string, then joins said slice by empty string element into returned string.
*/

func PrintAsciiArt(n string, y map[int][]string) string {
	stringSlc := []string{}

	for j := 0; j < len(y[32]); j++ { // Len(y[32]) is the number of elements for the first ascii character which equals eight (lines).
		for _, char := range n {
			stringSlc = append(stringSlc, y[int(char)][j])
		}
		stringSlc = append(stringSlc, "\n")
	}
	joinedString := strings.Join(stringSlc, "")
	return joinedString
}
