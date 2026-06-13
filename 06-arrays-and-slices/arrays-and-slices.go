// ============================================================
// TOPIC: Arrays and Slices in Go
// ============================================================
// CONCEPTS COVERED:
//  1. Array – declaration, fixed size
//  2. Array – initialization with values
//  3. Array – with implicit size (...)
//  4. Multi-dimensional arrays
//  5. Iterating arrays (index and range)
//  6. Slice – declaration and nil slice
//  7. Slice – created from array (shares underlying memory)
//  8. Slice – using make()
//  9. Slice – append() and growth
// 10. Slice – copy()
// 11. Slice of slices (2D slice)
// 12. Slice tricks: delete, insert, reverse
// 13. len() and cap() of slices
// 14. Slice literals
// 15. Passing slices to functions (reference semantics)
// 16. strings.Fields / strings.Split → slice of strings
// ============================================================

package main

import "fmt"

func main() {

	// ─────────────────────────────────────────
	// 1. Array – declaration with zero values
	// Arrays have a FIXED size, part of the type: [3]int ≠ [4]int
	// ─────────────────────────────────────────
	fmt.Println("═══ 1. Array declaration ═══")
	var scores [5]int // zero-initialized: [0 0 0 0 0]
	scores[0] = 10
	scores[1] = 20
	fmt.Println("scores:", scores)
	fmt.Printf("Length: %d\n", len(scores))

	// ─────────────────────────────────────────
	// 2. Array with values
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 2. Array with values ═══")
	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Println("primes:", primes)

	// ─────────────────────────────────────────
	// 3. Array with implicit size (...)
	// Compiler counts the elements
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 3. Array with ... size ═══")
	langs := [...]string{"Go", "Python", "Rust", "C"}
	fmt.Println("langs:", langs)
	fmt.Printf("Length: %d\n", len(langs))

	// ─────────────────────────────────────────
	// 4. Multi-dimensional arrays
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 4. Multi-dimensional array ═══")
	var matrix [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = i*3 + j + 1
		}
	}
	for _, row := range matrix {
		fmt.Println(row)
	}

	// ─────────────────────────────────────────
	// 5. Iterating arrays
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 5. Iterating arrays ═══")
	colors := [3]string{"red", "green", "blue"}
	// By index
	for i := 0; i < len(colors); i++ {
		fmt.Printf("  [%d] = %s\n", i, colors[i])
	}
	// By range
	for idx, val := range colors {
		fmt.Printf("  range[%d] = %s\n", idx, val)
	}

	// ─────────────────────────────────────────
	// 6. Slice – nil slice (zero value)
	// A slice is a dynamic view into an array
	// Anatomy: pointer | length | capacity
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 6. Nil slice ═══")
	var s []int // nil slice
	fmt.Printf("s = %v, len=%d, cap=%d, nil=%v\n", s, len(s), cap(s), s == nil)

	// ─────────────────────────────────────────
	// 7. Slice from array
	// syntax: array[low : high]  (low inclusive, high exclusive)
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 7. Slice from array ═══")
	arr := [6]int{10, 20, 30, 40, 50, 60}
	sl := arr[1:4] // elements at index 1, 2, 3
	fmt.Println("array :", arr)
	fmt.Println("slice [1:4]:", sl)

	// Modifying the slice modifies the underlying array!
	sl[0] = 999
	fmt.Println("After sl[0]=999:")
	fmt.Println("  slice:", sl)
	fmt.Println("  array:", arr) // arr[1] is now 999

	// ─────────────────────────────────────────
	// 8. Slice with make()
	// make([]T, length, capacity)
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 8. make() slice ═══")
	ms := make([]int, 5, 10) // len=5, cap=10
	fmt.Printf("make slice: %v  len=%d cap=%d\n", ms, len(ms), cap(ms))

	// ─────────────────────────────────────────
	// 9. append() – grows the slice
	// If capacity is exceeded, Go allocates a new, larger backing array
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 9. append() ═══")
	var ns []int
	for i := 1; i <= 5; i++ {
		ns = append(ns, i*10)
		fmt.Printf("  after append(%d): %v  len=%d cap=%d\n", i*10, ns, len(ns), cap(ns))
	}

	// Append multiple elements at once
	ns = append(ns, 60, 70, 80)
	fmt.Println("After appending 60,70,80:", ns)

	// Append one slice to another (spread operator)
	extra := []int{90, 100}
	ns = append(ns, extra...)
	fmt.Println("After appending extra...:", ns)

	// ─────────────────────────────────────────
	// 10. copy() – copies elements between slices
	// Copies min(len(dst), len(src)) elements
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 10. copy() ═══")
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src))
	n := copy(dst, src)
	fmt.Printf("Copied %d elements: %v\n", n, dst)

	// Modifying dst does NOT affect src
	dst[0] = 999
	fmt.Printf("src: %v  dst: %v\n", src, dst)

	// ─────────────────────────────────────────
	// 11. Slice of slices (2D slice / jagged slice)
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 11. Slice of slices ═══")
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][1] = "X"
	for _, row := range board {
		fmt.Println(row)
	}

	// ─────────────────────────────────────────
	// 12. Slice tricks
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 12. Slice tricks ═══")

	// Delete element at index 2
	data := []int{10, 20, 30, 40, 50}
	i := 2
	data = append(data[:i], data[i+1:]...)
	fmt.Println("After delete[2]:", data) // [10 20 40 50]

	// Insert 99 at index 1
	data = append(data[:1], append([]int{99}, data[1:]...)...)
	fmt.Println("After insert 99 at [1]:", data) // [10 99 20 40 50]

	// Reverse a slice
	rev := []int{1, 2, 3, 4, 5}
	for left, right := 0, len(rev)-1; left < right; left, right = left+1, right-1 {
		rev[left], rev[right] = rev[right], rev[left]
	}
	fmt.Println("Reversed:", rev)

	// ─────────────────────────────────────────
	// 13. len() vs cap()
	// len  = number of elements currently in the slice
	// cap  = total capacity of the underlying array from this pointer
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 13. len() and cap() ═══")
	base := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	view := base[2:5]
	fmt.Printf("base[2:5] = %v  len=%d cap=%d\n", view, len(view), cap(view))
	// cap = 10-2 = 8 (elements from index 2 to end of array)

	// ─────────────────────────────────────────
	// 14. Slice literal
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 14. Slice literal ═══")
	odds := []int{1, 3, 5, 7, 9} // No size → slice, not array
	fmt.Println("odds:", odds)
}

// 15. Slices are passed by reference
// Modifications inside the function affect the original slice
func doubleAll(nums []int) {
	for i := range nums {
		nums[i] *= 2
	}
}
