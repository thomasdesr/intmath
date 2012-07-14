package u64

// Pow returns x**y, the base-x exponential of y.
func Pow(x, y uint64) uint64 {
	if x == 0 {
		return 0
	}
	if x == 1 || y == 0 {
		return 1
	}
	r := uint64(1)
	for y > 0 {
		if y%2 == 1 {
			r *= x
		}
		x *= x
		y >>= 1
	}
	return r
}
