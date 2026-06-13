// ============================================================
// TOPIC: Variables in Go
// ============================================================
// CONCEPTS COVERED:
//  1. var keyword – explicit type declaration
//  2. var keyword – with type inference
//  3. Short variable declaration (:=)
//  4. Multiple variable declaration (single var block)
//  5. var block (grouped declarations)
//  6. Constants (const) – typed and untyped
//  7. iota – auto-incrementing constant generator
//  8. Zero values – default values when not initialized
//  9. Multiple assignment & swap
// 10. Blank identifier (_) to discard values
// 11. Scope: package-level vs function-level variables
// ============================================================

package main

import "fmt"

// ─────────────────────────────────────────
// 11. Package-level variables
// These are accessible throughout the entire file/package
// ─────────────────────────────────────────
var appName string = "GoLearner"   // explicit type + value
var appVersion = "1.0.0"           // type inferred as string
var totalUsers int                  // zero value = 0

// ─────────────────────────────────────────
// 6. Constants – value cannot change at runtime
// ─────────────────────────────────────────
const Pi = 3.14159       // untyped constant (flexible)
const MaxSize int = 100  // typed constant

// ─────────────────────────────────────────
// 7. iota – auto-incrementing values for related constants
// ─────────────────────────────────────────
const (
	Sunday    = iota // 0
	Monday           // 1
	Tuesday          // 2
	Wednesday        // 3
	Thursday         // 4
	Friday           // 5
	Saturday         // 6
)

// iota with expressions
const (
	_  = iota             // skip 0
	KB = 1 << (10 * iota) // 1 << 10 = 1024
	MB                    // 1 << 20 = 1048576
	GB                    // 1 << 30 = 1073741824
)

func main() {

	// ─────────────────────────────────────────
	// 1. var keyword – explicit type
	// ─────────────────────────────────────────
	var a int = 69
	var b float64 = 69.69
	var c string = "Hello, Go!"
	var d rune = '🔥' // rune = int32, holds a Unicode code point
	var e byte = 'G'  // byte = uint8, holds an ASCII character
	var f bool = true

	fmt.Println("── Explicit var declarations ──")
	fmt.Printf("a (int)    : %v\n", a)
	fmt.Printf("b (float64): %v\n", b)
	fmt.Printf("c (string) : %v\n", c)
	fmt.Printf("d (rune)   : %c  (code point: %d)\n", d, d)
	fmt.Printf("e (byte)   : %c  (ASCII: %d)\n", e, e)
	fmt.Printf("f (bool)   : %v\n\n", f)

	// ─────────────────────────────────────────
	// 2. var with type inference (no explicit type)
	// ─────────────────────────────────────────
	var score = 95        // Go infers this as int
	var ratio = 0.85      // Go infers this as float64
	var lang = "Go"       // Go infers this as string

	fmt.Println("── var with type inference ──")
	fmt.Printf("score : %v  (%T)\n", score, score)
	fmt.Printf("ratio : %v  (%T)\n", ratio, ratio)
	fmt.Printf("lang  : %v  (%T)\n\n", lang, lang)

	// ─────────────────────────────────────────
	// 3. Short variable declaration (:=)
	// ONLY usable inside functions; cannot be used at package level
	// ─────────────────────────────────────────
	name := "Gopher"        // string
	age := 5                // int
	isRunning := false      // bool
	temperature := 36.6     // float64

	fmt.Println("── Short variable declaration (:=) ──")
	fmt.Printf("name        : %v  (%T)\n", name, name)
	fmt.Printf("age         : %v  (%T)\n", age, age)
	fmt.Printf("isRunning   : %v  (%T)\n", isRunning, isRunning)
	fmt.Printf("temperature : %v  (%T)\n\n", temperature, temperature)

	// ─────────────────────────────────────────
	// 4. Multiple variables on one line
	// ─────────────────────────────────────────
	var x, y, z int = 1, 2, 3
	p, q := "hello", 42 // mixed types with :=

	fmt.Println("── Multiple variables on one line ──")
	fmt.Printf("x=%d, y=%d, z=%d\n", x, y, z)
	fmt.Printf("p=%q, q=%d\n\n", p, q)

	// ─────────────────────────────────────────
	// 5. var block (grouped declarations)
	// ─────────────────────────────────────────
	var (
		firstName string = "John"
		lastName  string = "Doe"
		zipCode   int    = 400001
		active    bool   = true
	)

	fmt.Println("── var block ──")
	fmt.Printf("Name: %s %s, Zip: %d, Active: %v\n\n", firstName, lastName, zipCode, active)

	// ─────────────────────────────────────────
	// 6 & 7. Constants and iota
	// ─────────────────────────────────────────
	fmt.Println("── Constants ──")
	fmt.Printf("Pi      = %v\n", Pi)
	fmt.Printf("MaxSize = %v\n\n", MaxSize)

	fmt.Println("── iota – Days of week ──")
	fmt.Printf("Sunday=%d, Monday=%d, Tuesday=%d, Saturday=%d\n\n",
		Sunday, Monday, Tuesday, Saturday)

	fmt.Println("── iota – Byte sizes ──")
	fmt.Printf("KB = %d\n", KB)
	fmt.Printf("MB = %d\n", MB)
	fmt.Printf("GB = %d\n\n", GB)

	// ─────────────────────────────────────────
	// 8. Zero values – defaults when not initialized
	// ─────────────────────────────────────────
	var zeroInt int
	var zeroFloat float64
	var zeroString string
	var zeroBool bool

	fmt.Println("── Zero Values ──")
	fmt.Printf("int    : %v\n", zeroInt)
	fmt.Printf("float64: %v\n", zeroFloat)
	fmt.Printf("string : %q\n", zeroString) // %q shows empty quotes ""
	fmt.Printf("bool   : %v\n\n", zeroBool)

	// ─────────────────────────────────────────
	// 9. Multiple assignment & variable swap
	// ─────────────────────────────────────────
	m, n := 10, 20
	fmt.Printf("Before swap: m=%d, n=%d\n", m, n)
	m, n = n, m // Go allows this elegant swap in one line
	fmt.Printf("After swap : m=%d, n=%d\n\n", m, n)

	// ─────────────────────────────────────────
	// 10. Blank identifier (_) – discards unwanted values
	// Very common when functions return multiple values
	// ─────────────────────────────────────────
	result, _ := divide(10, 3) // discard the remainder
	fmt.Printf("10 / 3 = %d  (remainder discarded with _)\n\n", result)

	// ─────────────────────────────────────────
	// 11. Package-level variable usage
	// ─────────────────────────────────────────
	fmt.Println("── Package-level variables ──")
	fmt.Printf("App: %s v%s, Users: %d\n", appName, appVersion, totalUsers)
}

// Helper function returning two values (quotient and remainder)
func divide(a, b int) (int, int) {
	return a / b, a % b
}
