package piscine

// Write a function CountIf that returns, the number of elements of a string slice, for which the f function returns true.

func CountIf(f func(string) bool, tab []string) int {
	var storageint int
	for _, val := range tab {
		if f(val) == true {
			storageint += 1
		}
	}
	return storageint
}
