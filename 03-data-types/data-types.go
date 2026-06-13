// ============================================================
// TOPIC: Data Types in Go
// ============================================================
// CONCEPTS COVERED:
//  1. Boolean (bool)
//  2. Signed integers: int8, int16, int32, int64, int
//  3. Unsigned integers: uint8, uint16, uint32, uint64, uint
//  4. Float types: float32, float64
//  5. Complex types: complex64, complex128
//  6. String type and raw string literals
//  7. Byte (alias for uint8) – single ASCII character
//  8. Rune (alias for int32) – single Unicode code point
//  9. uintptr – integer large enough to hold a pointer value
// 10. Type inference with :=
// 11. Explicit type conversion (casting)
// 12. Zero values – default values of each type
// ============================================================

package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	// ─────────────────────────────────────────
	// 1. BOOLEAN
	// ─────────────────────────────────────────
	var isGoAwesome bool = true
	var isBoring bool = false
	fmt.Println("── Boolean ──")
	fmt.Printf("isGoAwesome: %v | isBoring: %v\n\n", isGoAwesome, isBoring)

	// ─────────────────────────────────────────
	// 2. INTEGER TYPES
	// ─────────────────────────────────────────
	// Signed integers
	var i8 int8 = 127                    // -128 to 127
	var i16 int16 = 32767                // -32768 to 32767
	var i32 int32 = 2147483647          // -2^31 to 2^31-1
	var i64 int64 = 9223372036854775807 // -2^63 to 2^63-1
	var i int = 42                       // platform-dependent (32 or 64 bit)

	// Unsigned integers
	var u8 uint8 = 255                     // 0 to 255
	var u16 uint16 = 65535                 // 0 to 65535
	var u32 uint32 = 4294967295           // 0 to 2^32-1
	var u64 uint64 = 18446744073709551615 // 0 to 2^64-1
	var u uint = 42                        // platform-dependent

	fmt.Println("── Signed Integers ──")
	fmt.Printf("int8   : %v  (size: %v bytes)\n", i8, unsafe.Sizeof(i8))
	fmt.Printf("int16  : %v  (size: %v bytes)\n", i16, unsafe.Sizeof(i16))
	fmt.Printf("int32  : %v  (size: %v bytes)\n", i32, unsafe.Sizeof(i32))
	fmt.Printf("int64  : %v  (size: %v bytes)\n", i64, unsafe.Sizeof(i64))
	fmt.Printf("int    : %v  (size: %v bytes)\n\n", i, unsafe.Sizeof(i))

	fmt.Println("── Unsigned Integers ──")
	fmt.Printf("uint8  : %v  (size: %v bytes)\n", u8, unsafe.Sizeof(u8))
	fmt.Printf("uint16 : %v  (size: %v bytes)\n", u16, unsafe.Sizeof(u16))
	fmt.Printf("uint32 : %v  (size: %v bytes)\n", u32, unsafe.Sizeof(u32))
	fmt.Printf("uint64 : %v  (size: %v bytes)\n", u64, unsafe.Sizeof(u64))
	fmt.Printf("uint   : %v  (size: %v bytes)\n\n", u, unsafe.Sizeof(u))

	// ─────────────────────────────────────────
	// 3. FLOAT TYPES
	// ─────────────────────────────────────────
	var f32 float32 = math.Pi           // ~6-7 decimal digits precision
	var f64 float64 = math.Pi           // ~15-16 decimal digits precision

	fmt.Println("── Float Types ──")
	fmt.Printf("float32: %.10f  (size: %v bytes)\n", f32, unsafe.Sizeof(f32))
	fmt.Printf("float64: %.10f  (size: %v bytes)\n\n", f64, unsafe.Sizeof(f64))

	// ─────────────────────────────────────────
	// 4. COMPLEX TYPES
	// ─────────────────────────────────────────
	var c64 complex64 = complex(3, 4)   // 32-bit real + 32-bit imaginary
	var c128 complex128 = complex(5, 6) // 64-bit real + 64-bit imaginary

	fmt.Println("── Complex Types ──")
	fmt.Printf("complex64 : %v  real=%v imag=%v\n", c64, real(c64), imag(c64))
	fmt.Printf("complex128: %v  real=%v imag=%v\n\n", c128, real(c128), imag(c128))

	// ─────────────────────────────────────────
	// 5. STRING
	// ─────────────────────────────────────────
	var greeting string = "Hello, Gopher! 🐹"
	multiLine := `This is a
raw string literal
spanning multiple lines.`

	fmt.Println("── String ──")
	fmt.Printf("greeting   : %v\n", greeting)
	fmt.Printf("length     : %v bytes\n", len(greeting))
	fmt.Printf("multi-line :\n%v\n\n", multiLine)

	// ─────────────────────────────────────────
	// 6. BYTE (alias for uint8)
	// ─────────────────────────────────────────
	var b byte = 'G' // stores a single ASCII character
	fmt.Println("── Byte (uint8 alias) ──")
	fmt.Printf("byte value : %v  char: %c\n\n", b, b)

	// ─────────────────────────────────────────
	// 7. RUNE (alias for int32)
	// ─────────────────────────────────────────
	var r rune = '🔥' // stores a single Unicode code point (UTF-32)
	fmt.Println("── Rune (int32 alias) ──")
	fmt.Printf("rune value : %v  char: %c\n\n", r, r)

	// ─────────────────────────────────────────
	// 8. UINTPTR
	// ─────────────────────────────────────────
	x := 42
	ptr := uintptr(unsafe.Pointer(&x))
	fmt.Println("── uintptr ──")
	fmt.Printf("uintptr of x: %v\n\n", ptr)

	// ─────────────────────────────────────────
	// 9. TYPE INFERENCE  (:=)
	// ─────────────────────────────────────────
	inferred := 100        // int
	inferredF := 3.14      // float64
	inferredS := "Go!"     // string
	inferredB := true      // bool

	fmt.Println("── Type Inference ──")
	fmt.Printf("inferred int    : %v  (%T)\n", inferred, inferred)
	fmt.Printf("inferred float  : %v  (%T)\n", inferredF, inferredF)
	fmt.Printf("inferred string : %v  (%T)\n", inferredS, inferredS)
	fmt.Printf("inferred bool   : %v  (%T)\n\n", inferredB, inferredB)

	// ─────────────────────────────────────────
	// 10. TYPE CONVERSION
	// ─────────────────────────────────────────
	var original int = 42
	converted := float64(original)
	backToInt := int(converted)

	fmt.Println("── Type Conversion ──")
	fmt.Printf("original (int)     : %v\n", original)
	fmt.Printf("converted (float64): %v\n", converted)
	fmt.Printf("back to int        : %v\n\n", backToInt)

	// ─────────────────────────────────────────
	// 11. ZERO VALUES (default values)
	// ─────────────────────────────────────────
	var zeroInt int
	var zeroFloat float64
	var zeroString string
	var zeroBool bool

	fmt.Println("── Zero Values ──")
	fmt.Printf("int    zero value: %v\n", zeroInt)
	fmt.Printf("float64 zero value: %v\n", zeroFloat)
	fmt.Printf("string zero value: %q\n", zeroString)
	fmt.Printf("bool   zero value: %v\n", zeroBool)
}
