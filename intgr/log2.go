package intgr

// Log2 returns log base 2 of n. It's the same as index of the highest
// bit set in n. n <= 0 return -1.
func Log2(n int) int {
	if n <= 0 {
		return -1
	}
	var i uint
	if n<<32 == 0 {
		i = 24
	} else {
		i = 56
	}
	for ; i >= 0; i -= 8 {
		if n>>i > 0 {
			for i++; ; i++ {
				if n>>i == 0 {
					return int(i)
				}
			}
		}
	}
	return -1
}
