package i32

// Sqrt returns the square root of x. x < 0 returns -1. Based on a C implementation of 
// Newton's Method using bitshifting, originally found here:
// http://www.codecodex.com/wiki/Calculate_an_integer_square_root#C
func Sqrt(x int32) (r int32) {
	if x < 0 {
		return -1
	}
	// p starts at the highest power of four less or equal to x
	p := int32(1 << 30)
	for p > x {
		p >>= 2
	}

	for p != 0 {
		if x >= r+p {
			x -= r + p
			r += p << 1
		}
		r >>= 1
		p >>= 2
	}
	return
}
