package uintgr

// Log2 returns log base 2 of n. It's the same as index of the highest
// bit set in n. n = 0 returns 0
func Log2(n uint) (i uint) {
	if n<<32 == 0 {
		i = 24
	} else {
		i = 56
	}
	for ; i >= 0; i -= 8 {
		if n>>i > 0 {
			for ; ; i++ {
				if n>>i == 1 {
					return
				}
			}
		}
	}
	return
}
