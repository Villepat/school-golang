package piscine

// Write a function Any that returns true, for a string slice :
// if, when that string slice is passed through an f function, at least one element returns true.

func Any(f func(string) bool, a []string) bool {
	var returnboolean bool
	for _, stringval := range a {
		if f(stringval) == true {
			returnboolean = true
		}
	}
	return returnboolean
}
