package intgr

// Sqrt returns the square root of x. x < 0 returns -1. Based on a C implementation of 
// Newton's Method using bitshifting, originally found here:
// http://www.codecodex.com/wiki/Calculate_an_integer_square_root#C
func Sqrt(x int) (r int) {
	if x < 0 {
		return -1
	}
	// p starts at the highest power of four less or equal to x
	p := 1
	for p < x {
		p <<= 2
	}
	p >>= 2

	for p != 0 {
		if x >= r+p {
			x -= r + p
			r += p << 1
		}
		r >>= 1
		p >>= 2
	}
}
