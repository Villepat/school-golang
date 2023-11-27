package piscine

func Swap(a *int, b *int) {
	var x int
	x = *a
	*a = *b
	*b = x
}
