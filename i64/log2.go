package i64

// Log2 returns log base 2 of n. It's the same as index of the highest
// bit set in n. n <= 0 return -1.
func Log2(n int64) int64 {
	if n <= 0 {
		return -1
	}
	for i := uint(56); i >= 0; i -= 8 {
		if n>>i > 0 {
			for i++; ; i++ {
				if n>>i == 0 {
					return int64(i)
				}
			}
		}
	}
	return -1
}
