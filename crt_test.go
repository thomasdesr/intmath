package intmath

import (
	"code.google.com/p/intmath/i32"
	"code.google.com/p/intmath/i64"
	"code.google.com/p/intmath/intgr"
	"code.google.com/p/intmath/u32"
	"code.google.com/p/intmath/u64"
	"code.google.com/p/intmath/uintgr"
	//"math"
	"testing"
)

func TestCrt(t *testing.T) {
	t.Logf("Testing Crt\n")
	// 32 bits -> 10 bits max number
	// 64 bits -> 21 bits max number
	i32Max := int32(1<<10 - 1)
	i64Max := int64(1<<21 - 1)
	u32Max := uint32(1<<10 - 1)
	u64Max := uint64(1<<21 - 1)

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
		k := intgr.Crt(i * i * i)
		if k != i {
			t.Logf("intgr.Crt(%X*%X*%X) == %X\n", i, i, i, k)
			t.Fail()
		}
		k = intgr.Crt(-(i * i * i))
		if k != -i {
			t.Logf("intgr.Crt(%X*%X*%X) == %X\n", -i, -i, -i, k)
			t.Fail()
		}
	}
	for i := uMax; i > 0; i >>= 1 {
		k := uintgr.Crt(i * i * i)
		if k != i {
			t.Logf("uintgr.Crt(%X*%X*%X) == %X\n", i, i, i, k)
			t.Fail()
		}
	}
	for i := i32Max; i > 0; i >>= 1 {
		k := i32.Crt(i * i * i)
		if k != i {
			t.Logf("i32.Crt(%X*%X*%X) == %X\n", i, i, i, k)
			t.Fail()
		}
		k = i32.Crt(-(i * i * i))
		if k != -i {
			t.Logf("i32.Crt(%X*%X*%X) == %X\n", -i, -i, -i, k)
			t.Fail()
		}
	}
	for i := u32Max; i > 0; i >>= 1 {
		k := u32.Crt(i * i * i)
		if k != i {
			t.Logf("u32.Crt(%X*%X*%X) == %X\n", i, i, i, k)
			t.Fail()
		}
	}
	for i := i64Max; i > 0; i >>= 1 {
		k := i64.Crt(i * i * i)
		if k != i {
			t.Logf("i64.Crt(%X*%X*%X) == %X\n", i, i, i, k)
			t.Fail()
		}
		k = i64.Crt(-(i * i * i))
		if k != -i {
			t.Logf("i64.Crt(%X*%X*%X) == %X\n", -i, -i, -i, k)
			t.Fail()
		}
	}
	for i := u64Max; i > 0; i >>= 1 {
		k := u64.Crt(i * i * i)
		if k != i {
			t.Logf("u64.Crt(%X*%X*%X) == %X\n", i, i, i, k)
			t.Fail()
		}
	}
}
