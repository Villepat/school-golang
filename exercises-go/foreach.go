package piscine

// function that takes two arguments
// first argument has to be a function expecting an int
// second argument has to be a slice of int
// inside ForEach f and a stand for the argument names to be used within ForEach
func ForEach(f func(int), a []int) {
	for _, val := range a {
		f(val)
	}
}
