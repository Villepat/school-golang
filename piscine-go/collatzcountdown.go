package piscine

func CollatzCountdown(start int) int {
	count := 0
	if start <= 0 {
		count = -1
	} else if start > 1 {
		for i := 0; start != 1; i++ {
			count++
			if start%2 == 0 && start != 1 {
				start = start / 2
			} else if (start%2) == 1 && start != 1 {
				start = start * 3
				start = start + 1
			}
		}
	}
	return count
}
