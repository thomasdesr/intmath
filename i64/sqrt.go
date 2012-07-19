package i64

// Sqrt returns the square root of x. x < 0 returns -1. Based on a C implementation of 
// Newton's Method using bitshifting, originally found here:
// http://www.codecodex.com/wiki/Calculate_an_integer_square_root#C
func Sqrt(x int64) (r int64) {
	if x < 0 {
		return -1
	}
	// p starts at the highest power of four less or equal to x
	//Fast way to make p highest power of 4 <= x
	var n uint
	p := x
	if p >= 1<<32 {
		p >>= 32
		n = 32
	}
	if p >= 1<<16 {
		p >>= 16
		n += 16
	}
	if p >= 1<<8 {
		p >>= 8
		n += 8
	}
	if p >= 1<<4 {
		p >>= 4
		n += 4
	}
	if p >= 1<<2 {
		p >>= 2
		n += 2
	}
	p = 1 << n
	var b int64
	for ; p != 0; p >>= 2 {
		b = r | p
		r >>= 1
		if x >= b {
			x -= b
			r |= p
		}
	}
	return
}
