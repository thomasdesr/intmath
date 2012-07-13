package u64

// Log2 returns log base 2 of n. It's the same as index of the highest
// bit set in n. n == 0 returns 0
func Log2(n uint64) uint64 {
	for i := uint64(56); i >= 0; i -= 8 {
		if n>>i > 0 {
			for i++; ; i++ {
				if n>>i == 0 {
					return i
				}
			}
		}
	}
	return 0
}
