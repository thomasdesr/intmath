package intgr

// Log2 returns log base 2 of n. It's the same as index of the highest
// bit set in n. n <= 0 return -1.
func Log2(n int) int {
	if n <= 0 {
		return -1
	}
	for i := 1; ; i++ {
		if n>>i == 0 {
			return i
		}
	}

}
