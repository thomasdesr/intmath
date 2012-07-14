package intmath

import (
	"code.google.com/p/intmath/i32"
	"code.google.com/p/intmath/i64"
	"code.google.com/p/intmath/intgr"
	"code.google.com/p/intmath/u32"
	"code.google.com/p/intmath/u64"
	"code.google.com/p/intmath/uintgr"
	"math"
	"testing"
)

// ================================================
// ==== Benchmarking different Abs algorithms ====
// ================================================

// Bitmasking Abs as I found it in package bitbucket.org/SyntaxK/imath. int32 version
func BenchmarkAbs32bit1(b *testing.B) {
	for x := int32(0); -x < int32(b.N); x-- {
		_ = x ^ x>>31 + x>>31&1
	}
}

// Bitmasking Abs that seems simpler to me, used in library. int32 version
// Appears to be of similar speed on my 32bit system
func BenchmarkAbs32bit2(b *testing.B) {
	for x := int32(0); -x < int32(b.N); x-- {
		_ = x ^ x>>31 - x>>31
	}
}

// Bitmasking Abs as I found it in package bitbucket.org/SyntaxK/imath. int64 version
func BenchmarkAbs64bit1(b *testing.B) {
	for x := int64(0); -x < int64(b.N); x-- {
		_ = x ^ x>>63 + x>>63&1
	}
}

// Bitmasking Abs that seems simpler to me, used in library. int64 version
// Appears to be faster in speed on my 32bit system, presumably because 
// int64 is not natively supported, making the other operation more complex
func BenchmarkAbs64bit2(b *testing.B) {
	for x := int64(0); -x < int64(b.N); x-- {
		_ = x ^ x>>63 - x>>63
	}
}

// ================================================
// ==== Benchmarking different Epx algorithms ====
// ================================================

// Using i64 package's simple exp function
func BenchmarkExpInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := int64(0); j < 15; j++ {
			_ = i64.Pow(j, j)
		}
	}
}

// Using the standard Math package exp function and casting from/to int64
func BenchmarkExpFloat64ToInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := int64(0); j < 15; j++ {
			_ = int64(math.Pow(float64(j), float64(j)))
		}
	}
}

// Using i32 package's simple exp function
func BenchmarkExpInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := int32(0); j < 15; j++ {
			_ = i32.Pow(j, j>>1)
		}
	}

}

// Using the standard Math package exp function and casting from/to int32
func BenchmarkExpFloat64ToInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := int32(0); j < 15; j++ {
			_ = int32(math.Pow(float64(j), float64(j>>1)))
		}
	}
}

// ================================================
// ==== Benchmarking different Log2 algorithms ====
// ================================================

func BenchmarkLog2NaiveU64(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := uint64(1); i < 64; i++ {
			_ = log2U64(1 << i)
		}

	}
}

func BenchmarkLog2U64(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := uint64(1); i < 64; i++ {
			_ = loopU64Log2(1 << i)
		}
	}
}

func BenchmarkLog2bitmaskU64(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := uint64(1); i < 64; i++ {
			_ = bitmaskU64Log2(1 << i)
		}
	}
}

func log2U64(n uint64) (r uint64) {
	for n > 1 {
		n >>= 1
		r++
	}
	return
}

func loopU64Log2(n uint64) uint64 {
	for i := uint64(56); i >= 0; i -= 8 {
		if n>>i > 0 {
			for ; ; i++ {
				if n>>i == 1 {
					return i
				}
			}
		}
	}
	return 0
}

func bitmaskU64Log2(n uint64) uint64 {
	// uint is always set to the native word width, and it is therefore 
	// logical to expect that the operation is significantly faster on 
	// 32bit systems if we use it instead of uint64. On a netbook with
	// an Atom N280 CPU the benchmark confirm that this speeds up the
	// algorithm by 25%. The unnecessary cast on 64bit systems was 
	// considered worth the trade-off.
	var r, v uint
	if n >= 1<<32 {
		r += 32
		v = uint(n >> 32)
	} else {
		v = uint(n)
	}
	if v >= 1<<16 {
		r += 16
		v >>= 16
	}
	if v >= 1<<8 {
		r += 8
		v >>= 8
	}
	if v >= 1<<4 {
		r += 4
		v >>= 4
	}
	if v >= 1<<2 {
		r += 2
		v >>= 2
	}
	r += v >> 1
	return uint64(r)
}

func BenchmarkLog2NaiveU32(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := uint32(1); i < 32; i++ {
			_ = log2U32(1 << i)
		}

	}
}

func BenchmarkLog2U32(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := uint32(1); i < 32; i++ {
			_ = loopU32Log2(1 << i)
		}
	}
}

func BenchmarkLog2bitmaskU32(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := uint32(1); i < 32; i++ {
			_ = bitmaskU32Log2(1 << i)
		}
	}
}

func log2U32(n uint32) (r uint32) {
	for n > 1 {
		n >>= 1
		r++
	}
	return
}

func loopU32Log2(n uint32) uint32 {
	for i := uint32(25); i >= 0; i -= 8 {
		if n>>i > 0 {
			for ; ; i++ {
				if n>>i == 1 {
					return i
				}
			}
		}
	}
	return 0
}

func bitmaskU32Log2(n uint32) (r uint32) {
	if n >= 1<<16 {
		r += 16
		n >>= 16
	}
	if n >= 1<<8 {
		r += 8
		n >>= 8
	}
	if n >= 1<<4 {
		r += 4
		n >>= 4
	}
	if n >= 1<<2 {
		r += 2
		n >>= 2
	}
	r += n >> 1
	return
}

// ====================================================
// ==== Testing Log2 algorithm used in the library ====
// ====================================================

func TestLog2(t *testing.T) {
	t.Log("Testing Log2\n")
	iMax := uint(1)
	if iMax<<32 == 0 {
		iMax = 31
	} else {
		iMax = 63
	}
	for i := uint(0); i < iMax; i++ {
		for j := uint(0); j < i; j++ {
			if intgr.Log2((1<<i)+(1<<j)) != int(i) {
				t.Logf("intgr.Log2(%v) == %v\n", (1<<i)+(1<<j), intgr.Log2((1<<i)+(1<<j)))
				t.FailNow()
			}
		}
	}
	if intgr.Log2(-1) != -1 {
		t.Logf("intgr.Log2(-1) == %v\n", intgr.Log2(-1))
		t.FailNow()
	}

	for i := uint(0); i < iMax; i++ {
		for j := uint(0); j < i; j++ {
			if uintgr.Log2((1<<i)+(1<<j)) != i {
				t.Logf("intgr.Log2(%v) == %v\n", (1<<i)+(1<<j), uintgr.Log2((1<<i)+(1<<j)))
				t.Fail()
			}
		}
	}

	for i := uint32(0); i < uint32(31); i++ {
		for j := uint32(0); j < i; j++ {
			if i32.Log2((1<<i)+(1<<j)) != int32(i) {
				t.Logf("i32.Log2(%v) == %v\n", (1<<i)+(1<<j), i32.Log2((1<<i)+(1<<j)))
				t.FailNow()
			}
		}
	}
	if i32.Log2(-1) != int32(-1) {
		t.Logf("i32.Log2(-1) == %v\n", i32.Log2(-1))
		t.FailNow()
	}

	for i := uint32(0); i < uint32(31); i++ {
		for j := uint32(0); j < i; j++ {
			if u32.Log2((1<<i)+(1<<j)) != i {
				t.Logf("u32.Log2(%v) == %v\n", (1<<i)+(1<<j), u32.Log2((1<<i)+(1<<j)))
				t.FailNow()
			}
		}
	}

	for i := uint64(0); i < uint64(63); i++ {
		for j := uint64(0); j < i; j++ {
			if i64.Log2((1<<i)+(1<<j)) != int64(i) {
				t.Logf("i64.Log2(%v) == %v\n", (1<<i)+(1<<j), i64.Log2((1<<i)+(1<<j)))
				t.FailNow()
			}
		}
	}
	if i64.Log2(-1) != int64(-1) {
		t.Logf("i64.Log2(-1) == %v\n", i64.Log2(-1))
		t.FailNow()
	}

	for i := uint64(0); i < uint64(63); i++ {
		for j := uint64(0); j < i; j++ {
			if u64.Log2((1<<i)+(1<<j)) != i {
				t.Logf("u64.Log2(%v) == %v\n", (1<<i)+(1<<j), u64.Log2((1<<i)+(1<<j)))
				t.FailNow()
			}
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Logf("Testing Sqrt\n")
	i32Max := int32(0x7FFF)
	i64Max := int64(0x7FFFFFFF)
	u32Max := uint32(0xFFFF)
	u64Max := uint64(0xFFFFFFFF)

	iMax := 1
	var uMax uint
	if iMax<<32 == 0 {
		iMax = int(i32Max)
		uMax = uint(u32Max)
	} else {
		iMax = int(i64Max)
		uMax = uint(u64Max)
	}

	for i := iMax; i > 0; i >>= 1 {
		if intgr.Sqrt(i*i) != i {
			t.Logf("intgr.Sqrt(%X*%X) == %X\n", i, i, intgr.Sqrt(i))
			t.Fail()
		}
	}
	for i := uMax; i > 0; i >>= 1 {
		if uintgr.Sqrt(i*i) != i {
			t.Logf("uintgr.Sqrt(%X*%X) == %X\n", i, i, uintgr.Sqrt(i))
			t.Fail()
		}
	}
	for i := i32Max; i > 0; i >>= 1 {
		if i32.Sqrt(i*i) != i {
			t.Logf("i32.Sqrt(%X*%X) == %X\n", i, i, i32.Sqrt(i))
			t.Fail()
		}
	}
	for i := u32Max; i > 0; i >>= 1 {
		if u32.Sqrt(i*i) != i {
			t.Logf("u32.Sqrt(%X*%X) == %X\n", i, i, u32.Sqrt(i))
			t.Fail()
		}
	}
	for i := i64Max; i > 0; i >>= 1 {
		if i64.Sqrt(i*i) != i {
			t.Logf("i64.Sqrt(%X*%X) == %X\n", i, i, i64.Sqrt(i))
			t.Fail()
		}
	}
	for i := u64Max; i > 0; i >>= 1 {
		if u64.Sqrt(i*i) != i {
			t.Logf("u64.Sqrt(%X*%X) == %X\n", i, i, u64.Sqrt(i))
			t.Fail()
		}
	}
}
