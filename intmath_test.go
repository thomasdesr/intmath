package intmath

import (
	"intmath/i32"
	"intmath/i64"
	"intmath/intgr"
	"intmath/u32"
	"intmath/u64"
	"intmath/uintgr"
	"math"
	"testing"
)

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

// Using i64 package's simple exp function
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

func TestLog2(t *testing.T) {
	fmt.Println("Testing Log2")
	iMax := 1
	if iMax<<32 == 0 {
		iMax = 31
	} else {
		iMax = 63
	}
	for i := 0; i < iMax; i++ {
		if intgr.Log2(1<<uint(i)) != i {
			t.Logf("intgr.Log2(1<<%v) == %v\n", i, intgr.Log2(1<<uint(i)))
			t.FailNow()
		}
	}

	for i := uint(0); i < uint(iMax); i++ {
		if uintgr.Log2(1<<i) != i {
			t.Logf("uintgr.Log2(1<<%v) == %v\n", i, uintgr.Log2(1<<uint(i)))
			t.Fail()
		}
	}
	for i := int32(0); i < 31; i++ {
		if i32.Log2(1<<uint(i)) != i {
			t.Logf("i32.Log2(1<<%v) == %v\n", i, i32.Log2(1<<uint(i)))
			t.Fail()
		}
	}
	for i := int64(0); i < 63; i++ {
		if i64.Log2(1<<uint(i)) != i {
			t.Logf("i64.Log2(1<<%v) == %v\n", i, i64.Log2(1<<uint(i)))
			t.Fail()
		}
	}
	for i := uint32(0); i < 32; i++ {
		if u32.Log2(1<<uint(i)) != i {
			t.Logf("u32.Log2(1<<%v) == %v\n", i, u32.Log2(1<<uint(i)))
			t.Fail()
		}
	}
	for i := uint64(0); i < 64; i++ {
		if u64.Log2(1<<uint(i)) != i {
			t.Logf("u64.Log2(1<<%v) != %v\n", i, u64.Log2(1<<uint(i)))
			t.Fail()
		}
	}
}

func TestSqrt(t *testing.T) {
	fmt.Println("Testing Sqrt")
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
