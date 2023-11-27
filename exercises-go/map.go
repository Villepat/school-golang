package piscine

// Write a function Map that, for an int slice, applies a function of this type func(int) bool on each element of that slice and
// returns a slice of all the return values.

func Map(f func(int) bool, a []int) []bool {
	storageslice := []bool{}
	for _, val := range a {
		storageslice = append(storageslice, f(val))
	}
	return storageslice
}
