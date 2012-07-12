package uintgr

// Pow returns x**y, the base-x exponential of y.
func Pow(x, y uint32) uint32 {
	if x == 0 {
		return 0
	}
	if x == 1 || y == 0 {
		return 1
	}
	r := uint32(1)
	for y > 0 {
		if y%2 == 0 {
			r *= x
		}
		x *= x
		y >>= 1
	}
	return r
}
