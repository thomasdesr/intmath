package i32

// Log2 returns log base 2 of n. It's the same as index of the highest
// bit set in n. n <= 0 return -1.
func Log2(n int32) int32 {
	if n <= 0 {
		return -1
	}
	for i := uint32(24); i >= 0; i -= 8 {
		if n>>i > 0 {
			for ; ; i++ {
				if n>>i == 0 {
					return int32(i)
				}
			}
		}
	}
	return -1
}
