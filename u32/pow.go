package u32

// Pow returns x**y, the base-x exponential of y.
func Pow(x, y uint32) (r uint32) {
	if x == 0 {
		return
	}
	r = 1
	if x == 1 || y == 0 {
		return
	}
	for y > 0 {
		if y&1 == 1 {
			r *= x
		}
		x *= x
		y >>= 1
	}
	return
}
