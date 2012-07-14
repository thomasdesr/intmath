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

func BenchmarkSqrtIntgr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := uint(0); j < 64; j++ {
			_ = intgr.Sqrt(int(1) << j)
		}
	}
}

func BenchmarkSqrtFloat64ToIntgr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := uint(0); j < 64; j++ {
			_ = int(math.Sqrt(float64(int(1) << j)))
		}
	}
}

func BenchmarkSqrtInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := uint(0); j < 64; j++ {
			_ = i64.Sqrt(int64(1) << j)
		}
	}
}

func BenchmarkSqrtFloat64ToInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := uint(0); j < 64; j++ {
			_ = int64(math.Sqrt(float64(int64(1) << j)))
		}
	}
}

func BenchmarkSqrtInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := uint(0); j < 32; j++ {
			_ = i32.Sqrt(int32(1) << j)
		}
	}
}

func BenchmarkSqrtFloat64ToInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := uint(0); j < 32; j++ {
			_ = int32(math.Sqrt(float64(int32(1) << j)))
		}
	}
}

func BenchmarkSqrtUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := uint(0); j < 64; j++ {
			_ = u64.Sqrt(uint64(1) << j)
		}
	}
}

func BenchmarkSqrtFloat64ToUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := uint(0); j < 64; j++ {
			_ = uint64(math.Sqrt(float64(uint64(1) << j)))
		}
	}
}

func BenchmarkSqrtUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := uint(0); j < 32; j++ {
			_ = u32.Sqrt(uint32(1) << j)
		}
	}
}

func BenchmarkSqrtFloat64ToUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := uint(0); j < 32; j++ {
			_ = int64(math.Sqrt(float64(uint(1) << j)))
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
