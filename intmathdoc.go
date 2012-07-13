// Package intmath provides a simple integer math library.
//
// The standard Go math library uses the float64 type for everything. 
// This often results in unnecessary casting, resulting in both more 
// verbose and (possibly) slower code. This library aims to become 
// an extension to the Go math library by providing math packages for
// the int32, uint32, int64 and uint64 types.
//
// Note that the dedicated FPU in your computer can be so fast that 
// casting to and from float64 would be faster than using integer 
// math. In other words, the standard math library can in theory 
// outperform these libraries despite repeated casting to and from
// float64 - benchmark code will be included so you can easily test
// this for yourself.
//
// Not all math functions will be copied, as not all of them make 
// sense in this context. At the moment, the following functions have been implemented:
//
//  Abs(x)     (only for signed types, obviously)
//  GCD(a, b)  (adapted from github.com/cznic/mathutil)
//  Min(x, y)
//  Max(x, y)
//  Pow(x, y)
//  Sqrt(x)
//  Log2(n)
//
// The intmath package itself is just a dummy package used for
// documentation - the real library is in the subpackages. Import
// the relevant type to use math functions tailored to that type.
//
// As a rule, all functions take and return types matching the type
// of the library, so functions in u64 only takes and returns uint64. 
// This both reduces verbosity and avoids the need for functions that 
// explicitly name the type they apply to. This makes it easy to 
// understand what code like this is doing:
// 
// z := u64.Pow(x, y)
//
// Think of it as fake function overloading.
package intmath
