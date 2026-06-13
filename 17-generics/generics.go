// ============================================================
// TOPIC: Generics in Go (Go 1.18+)
// ============================================================
// CONCEPTS COVERED:
//  1. Why generics? Problem without generics
//  2. Type parameters – basic syntax [T any]
//  3. Type constraints – interface-based constraints
//  4. Built-in constraints: any, comparable
//  5. constraints package (golang.org/x/exp/constraints) alternatives
//  6. Custom type constraints with unions (A | B)
//  7. Generic functions
//  8. Generic types (structs, slices, maps)
//  9. Multiple type parameters
// 10. Type inference – compiler can infer T
// 11. Generic Stack implementation
// 12. Generic Map, Filter, Reduce (functional helpers)
// 13. Constraints with methods (interface with methods)
// 14. Instantiating generic types
// ============================================================

package main

import (
	"fmt"
)

// ─────────────────────────────────────────
// 3 & 6. Custom type constraints
// ─────────────────────────────────────────

// Number matches any numeric type
type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

// Ordered types that support < > <= >=
type Ordered interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64 | string
}

// ─────────────────────────────────────────
// 11. Generic Stack
// ─────────────────────────────────────────

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}

func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// ─────────────────────────────────────────
// 8. Generic Pair type
// ─────────────────────────────────────────

type Pair[K, V any] struct {
	Key   K
	Value V
}

func NewPair[K, V any](k K, v V) Pair[K, V] {
	return Pair[K, V]{Key: k, Value: v}
}

// ─────────────────────────────────────────
// Generic Result type (like Rust's Result<T, E>)
// ─────────────────────────────────────────

type Result[T any] struct {
	value T
	err   error
}

func Ok[T any](value T) Result[T] {
	return Result[T]{value: value}
}

func Err[T any](err error) Result[T] {
	return Result[T]{err: err}
}

func (r Result[T]) IsOk() bool    { return r.err == nil }
func (r Result[T]) Unwrap() T     { return r.value }
func (r Result[T]) Error() error  { return r.err }

// ─────────────────────────────────────────
// 7. Generic functions
// ─────────────────────────────────────────

// Sum works with any Number type
func Sum[T Number](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

// Min returns the smaller of two Ordered values
func Min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max returns the larger of two Ordered values
func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Contains checks whether a slice contains a value (requires comparable)
func Contains[T comparable](slice []T, target T) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

// Keys extracts keys from any map
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Values extracts values from any map
func Values[K comparable, V any](m map[K]V) []V {
	vals := make([]V, 0, len(m))
	for _, v := range m {
		vals = append(vals, v)
	}
	return vals
}

// ─────────────────────────────────────────
// 12. Functional helpers: Map, Filter, Reduce
// ─────────────────────────────────────────

// Map applies transform to each element and returns new slice
func Map[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// Filter returns elements for which predicate is true
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce folds a slice into a single value
func Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U {
	acc := initial
	for _, v := range slice {
		acc = fn(acc, v)
	}
	return acc
}

// ─────────────────────────────────────────
// 13. Constraints with methods
// ─────────────────────────────────────────

type Stringer interface {
	String() string
}

func PrintAll[T Stringer](items []T) {
	for _, item := range items {
		fmt.Println(item.String())
	}
}

type Temperature struct {
	Value float64
	Unit  string
}

func (t Temperature) String() string {
	return fmt.Sprintf("%.2f°%s", t.Value, t.Unit)
}

func main() {

	// ─────────────────────────────────────────
	// 7. Generic functions
	// ─────────────────────────────────────────
	fmt.Println("═══ 7. Generic functions ═══")

	// Sum with different types
	ints := []int{1, 2, 3, 4, 5}
	floats := []float64{1.1, 2.2, 3.3}
	fmt.Printf("Sum(ints)   = %v\n", Sum(ints))    // 10. type inferred
	fmt.Printf("Sum(floats) = %.1f\n", Sum(floats))

	// Min & Max
	fmt.Printf("Min(3,7)    = %d\n", Min(3, 7))
	fmt.Printf("Max(\"apple\",\"banana\") = %s\n", Max("apple", "banana"))

	// ─────────────────────────────────────────
	// 4. comparable constraint
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 4. comparable (Contains) ═══")
	fmt.Printf("Contains([1,2,3], 2): %v\n", Contains([]int{1, 2, 3}, 2))
	fmt.Printf("Contains([\"a\",\"b\"], \"c\"): %v\n", Contains([]string{"a", "b"}, "c"))

	// ─────────────────────────────────────────
	// 9. Multiple type parameters (Keys/Values)
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 9. Multiple type params (Keys/Values) ═══")
	scores := map[string]int{"Alice": 95, "Bob": 87}
	fmt.Println("Keys:", Keys(scores))
	fmt.Println("Values:", Values(scores))

	// ─────────────────────────────────────────
	// 11. Generic Stack
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 11. Generic Stack ═══")
	intStack := &Stack[int]{}
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	fmt.Printf("Stack len: %d\n", intStack.Len())

	for !intStack.IsEmpty() {
		val, _ := intStack.Pop()
		fmt.Printf("  Popped: %d\n", val)
	}

	// String stack
	strStack := &Stack[string]{}
	strStack.Push("go")
	strStack.Push("generics")
	top, _ := strStack.Peek()
	fmt.Printf("String stack top: %q\n", top)

	// ─────────────────────────────────────────
	// 8. Generic Pair
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 8. Generic Pair ═══")
	p1 := NewPair("name", "Alice")
	p2 := NewPair(1, 3.14)
	fmt.Printf("p1: %v → %v\n", p1.Key, p1.Value)
	fmt.Printf("p2: %v → %v\n", p2.Key, p2.Value)

	// ─────────────────────────────────────────
	// Generic Result type
	// ─────────────────────────────────────────
	fmt.Println("\n═══ Generic Result type ═══")
	r1 := Ok(42)
	r2 := Err[int](fmt.Errorf("something failed"))
	fmt.Printf("r1 ok=%v, val=%d\n", r1.IsOk(), r1.Unwrap())
	fmt.Printf("r2 ok=%v, err=%v\n", r2.IsOk(), r2.Error())

	// ─────────────────────────────────────────
	// 12. Map, Filter, Reduce
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 12. Map / Filter / Reduce ═══")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	doubled := Map(numbers, func(n int) int { return n * 2 })
	fmt.Println("Doubled:", doubled)

	evens := Filter(numbers, func(n int) bool { return n%2 == 0 })
	fmt.Println("Evens:", evens)

	total := Reduce(numbers, 0, func(acc, n int) int { return acc + n })
	fmt.Println("Sum via Reduce:", total)

	// Chain: filter evens, square them, sum
	result := Reduce(
		Map(
			Filter(numbers, func(n int) bool { return n%2 == 0 }),
			func(n int) int { return n * n },
		),
		0,
		func(acc, n int) int { return acc + n },
	)
	fmt.Println("Sum of squares of evens:", result)

	// ─────────────────────────────────────────
	// 13. Constraints with methods
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 13. Constraints with methods ═══")
	temps := []Temperature{
		{Value: 100, Unit: "C"},
		{Value: 212, Unit: "F"},
		{Value: 373.15, Unit: "K"},
	}
	PrintAll(temps)
}
