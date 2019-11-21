package main

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func percent(fraction, total int) int {
	if 0 == total {
		return 0
	}

	return int(float32(100) / float32(total) * float32(fraction))
}
