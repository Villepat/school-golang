package piscine

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	if y > 0 && x > 0 {
		for i := 1; i <= x; i++ {
			if i == 1 {
				z01.PrintRune('A')
			} else if i == x {
				z01.PrintRune('C')
			} else {
				z01.PrintRune('B')
			}
		}

		z01.PrintRune('\n')
		for i := 2; i < y; i++ {
			for j := 1; j <= x; j++ {
				if x == 1 {
					z01.PrintRune('B')
					z01.PrintRune('\n')
				} else if j < 2 {
					z01.PrintRune('B')
				} else if j < x {
					z01.PrintRune(' ')
				} else if j == x {
					z01.PrintRune('B')
					z01.PrintRune('\n')
				}
			}
		}

		if y > 1 {
			for i := 1; i <= x; i++ {
				if i == 1 {
					z01.PrintRune('A')
				} else if i == x {
					z01.PrintRune('C')
				} else {
					z01.PrintRune('B')
				}
			}
		}
		z01.PrintRune('\n')
	}
}
