package u64

// Log2 returns log base 2 of n. It's the same as index of the highest 
// bit set in n. n == 0 returns 0.
func Log2(n uint64) uint64 {
	// It's possible that a lookup-table approach is faster, 
	// but this is simpler code
	for i := uint64(1); ; i++ {
		if n>>i == 0 {
			return i
		}
	}
}
