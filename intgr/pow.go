package intgr

// Pow returns x**y, the base-x exponential of y.
func Pow(x, y int) int {
	if x < 0 {
		x = -x
		if Abs(y%2) == 1 {
			return -Pow(x, y)
		}
	}
	if x == 1 || y == 0 {
		return 1
	}
	if x == 0 || y < 0 {
		return 0
	}
	r := int32(1)
	for y > 0 {
		if y%2 == 0 {
			r *= x
		}
		x *= x
		y >>= 1
	}
	return r
}
