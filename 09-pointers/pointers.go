// ============================================================
// TOPIC: Pointers in Go
// ============================================================
// CONCEPTS COVERED:
//  1. What is a pointer? (&, * operators)
//  2. Declare a pointer variable
//  3. Dereferencing a pointer
//  4. nil pointer
//  5. new() – allocate zero-valued heap memory
//  6. Pointer to struct
//  7. Passing pointer to a function (pass-by-reference effect)
//  8. Returning pointer from a function (heap allocation)
//  9. Pointer to pointer (**T)
// 10. Pointer arithmetic – NOT allowed in Go
// 11. When to use pointers vs values
// 12. Slice & map are already reference types (no explicit pointer needed)
// ============================================================

package main

import "fmt"

// ─────────────────────────────────────────
// Simple struct for demos
// ─────────────────────────────────────────
type Point struct {
	X, Y int
}

func main() {

	// ─────────────────────────────────────────
	// 1. The basics: & (address-of) and * (dereference)
	// & → takes the address of a variable
	// * → dereferences a pointer (gets the value it points to)
	// ─────────────────────────────────────────
	fmt.Println("═══ 1. & and * operators ═══")
	x := 42
	p := &x          // p is a pointer to x; type is *int
	fmt.Printf("x  = %d\n", x)
	fmt.Printf("&x = %v  (address)\n", &x)
	fmt.Printf("p  = %v  (same address)\n", p)
	fmt.Printf("*p = %d  (value at address)\n", *p)

	// ─────────────────────────────────────────
	// 2. Declare a pointer variable
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 2. Pointer declaration ═══")
	var ptr *int // zero value of a pointer is nil
	fmt.Printf("ptr (before) = %v\n", ptr)
	n := 100
	ptr = &n
	fmt.Printf("ptr (after)  = %v, *ptr = %d\n", ptr, *ptr)

	// ─────────────────────────────────────────
	// 3. Dereferencing – modify the original value through a pointer
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 3. Dereferencing ═══")
	val := 5
	pVal := &val
	fmt.Printf("Before: val = %d\n", val)
	*pVal = 99 // modifies val through the pointer
	fmt.Printf("After : val = %d  (changed via *pVal)\n", val)

	// ─────────────────────────────────────────
	// 4. nil pointer – zero value of any pointer type
	// Dereferencing nil causes a runtime PANIC
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 4. nil pointer ═══")
	var nilPtr *int
	fmt.Printf("nilPtr = %v, is nil? %v\n", nilPtr, nilPtr == nil)
	// Always guard before dereferencing
	if nilPtr != nil {
		fmt.Println(*nilPtr)
	} else {
		fmt.Println("nilPtr is nil – safe guard prevented panic")
	}

	// ─────────────────────────────────────────
	// 5. new() – allocates zeroed memory on the heap
	// Returns a pointer to the new value
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 5. new() ═══")
	numPtr := new(int)         // *int, value is 0
	fmt.Printf("*numPtr = %d (zero value)\n", *numPtr)
	*numPtr = 77
	fmt.Printf("*numPtr = %d (after assignment)\n", *numPtr)

	// ─────────────────────────────────────────
	// 6. Pointer to struct
	// Go automatically dereferences: pt.X == (*pt).X
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 6. Pointer to struct ═══")
	pt := &Point{X: 3, Y: 4}
	fmt.Printf("Point = %+v\n", pt)
	pt.X = 10              // auto-dereference, same as (*pt).X = 10
	fmt.Printf("After pt.X=10: %+v\n", *pt)

	// ─────────────────────────────────────────
	// 7. Passing a pointer to a function
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 7. Pointer as function argument ═══")
	score := 50
	fmt.Printf("Before increment: %d\n", score)
	increment(&score) // pass address
	fmt.Printf("After  increment: %d\n", score) // 51

	// Compare with value parameter (no mutation)
	tryIncrement(score)
	fmt.Printf("After tryIncrement (value): %d\n", score) // still 51

	// ─────────────────────────────────────────
	// 8. Returning pointer from a function
	// Go is safe: the variable escapes to the heap automatically
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 8. Returning pointer from function ═══")
	newPoint := createPoint(5, 10)
	fmt.Printf("newPoint = %+v\n", *newPoint)

	// ─────────────────────────────────────────
	// 9. Pointer to pointer (**T)
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 9. Pointer to pointer ═══")
	a := 42
	b := &a    // *int
	c := &b    // **int
	fmt.Printf("a   = %d\n", a)
	fmt.Printf("*b  = %d\n", *b)
	fmt.Printf("**c = %d\n", **c)
	**c = 999
	fmt.Printf("a after **c=999: %d\n", a)

	// ─────────────────────────────────────────
	// 10. Pointer arithmetic – NOT ALLOWED in Go
	// Go does NOT allow ptr++ or ptr+1 like C/C++
	// For unsafe pointer arithmetic use unsafe.Pointer (advanced)
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 10. Pointer arithmetic – NOT allowed ═══")
	fmt.Println("Go intentionally disallows pointer arithmetic for safety.")
	fmt.Println("Use slice indexing instead of manual pointer offsets.")

	// ─────────────────────────────────────────
	// 11. When to use pointers
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 11. When to use pointers ═══")
	fmt.Println("Use pointers when:")
	fmt.Println("  • Function needs to modify the original variable")
	fmt.Println("  • Struct is large and copying is expensive")
	fmt.Println("  • You want to represent 'no value' (nil)")
	fmt.Println("Use values when:")
	fmt.Println("  • Small types (int, bool, small structs)")
	fmt.Println("  • You want immutability / no side-effects")

	// ─────────────────────────────────────────
	// 12. Slices and maps are already reference types
	// You rarely need a *[]T or *map[K]V
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 12. Slices/Maps are reference types ═══")
	sl := []int{1, 2, 3}
	addToSlice(sl) // slice header is copied, but underlying array shared
	fmt.Println("Slice after addToSlice:", sl) // [99 2 3] – first element changed
}

// ─────────────────────────────────────────
// Helper functions
// ─────────────────────────────────────────

// 7. Pointer receiver – mutates the original
func increment(n *int) {
	*n++
}

// 7. Value receiver – works on a copy
func tryIncrement(n int) {
	n++ // local copy only
}

// 8. Returns a pointer; variable escapes to heap safely
func createPoint(x, y int) *Point {
	p := Point{X: x, Y: y}
	return &p // safe in Go – compiler handles escape analysis
}

// 12. Modifying slice via function (no *[] needed)
func addToSlice(s []int) {
	s[0] = 99 // modifies underlying array
}
