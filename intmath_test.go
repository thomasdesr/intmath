package intmath

import (
	"fmt"
	"intmath/i32"
	"intmath/i64"
	//	"intmath/intgr"
	//	"intmath/u32"
	//	"intmath/u64"
	//	"intmath/uintgr"
	"math"
	"testing"
)

// Bitmasking Abs as I found it in package bitbucket.org/SyntaxK/imath
func BenchmarkAbs1(b *testing.B) {
	fmt.Println("Testing Abs version a")
	fmt.Println(int64(b.N))
	for x := int32(0); -x < int32(b.N); x-- {
		_ = x ^ x>>31 + x>>31&1
	}
}

// Bitmasking Abs that seems simpler to me...
func BenchmarkAbs2(b *testing.B) {
	fmt.Println("Testing Abs version b")
	fmt.Println(int64(b.N))
	for x := int32(0); -x < int32(b.N); x-- {
		_ = x ^ x>>31 - x>>31
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

// Using the standard Math package exp function and casting from/to int64
func BenchmarkExpFloat64ToInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := int32(0); j < 15; j++ {
			_ = int32(math.Pow(float64(j), float64(j>>1)))
		}
	}
}
