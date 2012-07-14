package u64

// Sqrt returns the square root of x. Based on a C implementation of 
// Newton's Method using bitshifting, originally found here:
// http://www.codecodex.com/wiki/Calculate_an_integer_square_root#C
func Sqrt(x uint64) (r uint64) {
	// p starts at the highest power of four less or equal to x
	p := uint64(1 << 62)
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
