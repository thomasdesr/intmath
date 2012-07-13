package i64

// Log2 returns log base 2 of n. It's the same as index of the highest
// bit set in n. n <= 0 return -1.
func Log2(n uint) uint {
	for i := uint(1); ; i++ {
		if n>>i == 0 {
			return i
		}
	}
}
