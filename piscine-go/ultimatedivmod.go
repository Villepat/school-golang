package piscine

func UltimateDivMod(a *int, b *int) {
	x := *a / *b
	y := *a % *b
	*a = x
	*b = y
}
