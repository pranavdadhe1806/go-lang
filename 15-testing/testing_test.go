// ============================================================
// TOPIC: Testing in Go – Actual Test File
// ============================================================
// This file demonstrates how _test.go files look.
// Run: go test -v ./15-testing/
// Run benchmarks: go test -bench=. ./15-testing/
// ============================================================

package main

import (
	"testing"
)

// ─────────────────────────────────────────
// 2. Basic test function
// Must be named TestXxx; receives *testing.T
// ─────────────────────────────────────────
func TestAdd(t *testing.T) {
	// 3. t.Errorf – report failure but continue
	got := Add(2, 3)
	want := 5
	if got != want {
		t.Errorf("Add(2, 3) = %d; want %d", got, want)
	}
}

// ─────────────────────────────────────────
// 5. Table-driven test (idiomatic Go pattern)
// Define test cases as a slice of structs
// ─────────────────────────────────────────
func TestAdd_TableDriven(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"positive", 2, 3, 5},
		{"negative", -1, -2, -3},
		{"zero", 0, 0, 0},
		{"mixed", -5, 10, 5},
	}

	// 6. Subtests with t.Run()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Add(tc.a, tc.b)
			if got != tc.want {
				t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

// ─────────────────────────────────────────
// Testing functions that return errors
// ─────────────────────────────────────────
func TestDivide(t *testing.T) {
	// Happy path
	got, err := Divide(10, 2)
	if err != nil {
		t.Fatalf("Divide(10, 2) unexpected error: %v", err) // 4. t.Fatalf stops test
	}
	if got != 5.0 {
		t.Errorf("Divide(10, 2) = %f; want 5.0", got)
	}

	// Error path
	_, err = Divide(10, 0)
	if err == nil {
		t.Error("Divide(10, 0) expected error, got nil")
	}
}

// ─────────────────────────────────────────
// Table-driven test for IsPalindrome
// ─────────────────────────────────────────
func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"racecar", true},
		{"hello", false},
		{"madam", true},
		{"go", false},
		{"", true},
		{"a", true},
		{"A man a plan a canal Panama", false}, // spaces break it
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got := IsPalindrome(tc.input)
			if got != tc.want {
				t.Errorf("IsPalindrome(%q) = %v; want %v", tc.input, got, tc.want)
			}
		})
	}
}

// ─────────────────────────────────────────
// 7. Test helper function
// t.Helper() marks this as a helper so failures point to the caller
// ─────────────────────────────────────────
func assertEqual(t *testing.T, got, want int) {
	t.Helper() // marks this function as a test helper
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestFibonacci(t *testing.T) {
	assertEqual(t, Fibonacci(0), 0)
	assertEqual(t, Fibonacci(1), 1)
	assertEqual(t, Fibonacci(5), 5)
	assertEqual(t, Fibonacci(10), 55)
}

// ─────────────────────────────────────────
// 11. Skipping tests
// ─────────────────────────────────────────
func TestExpensive(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping expensive test in short mode") // run with -short flag to skip
	}
	// Expensive computation would go here
}

// ─────────────────────────────────────────
// 8. Benchmark functions
// Run with: go test -bench=. ./...
// ─────────────────────────────────────────
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(100, 200) // b.N is calibrated by the test runner
	}
}

func BenchmarkFibonacci20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(20)
	}
}

// ─────────────────────────────────────────
// 9. Example functions
// Also act as runnable documentation; output is verified
// ─────────────────────────────────────────
func ExampleAdd() {
	fmt.Println(Add(1, 2))
	// Output:
	// 3
}

func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome("racecar"))
	fmt.Println(IsPalindrome("hello"))
	// Output:
	// true
	// false
}
