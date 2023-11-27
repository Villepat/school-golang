package piscine

import "fmt"

func UltimatePointOne(n ***int) {
	***n = 1
}

func main() {
	a := 5
	b := &a
	c := &b
	d := &c
	fmt.Println(a)

	UltimatePointOne(d)
	fmt.Println(a)
}
