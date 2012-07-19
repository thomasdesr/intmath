package i64

// Pow returns x**y, the base-x exponential of y.
func Pow(x, y int64) (r int64) {
	if x == 0 || y < 0 {
		return
	}
	r = 1
	if x == 1 || y == 0 {
		return
	}
	if x < 0 {
		x = -x
		if y&1 == 1 {
			r = -1
		}
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
