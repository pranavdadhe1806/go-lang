// ============================================================
// TOPIC: Testing in Go
// ============================================================
// CONCEPTS COVERED:
//  1. Go testing package overview
//  2. Basic test function (TestXxx)
//  3. t.Error / t.Errorf – mark failure, continue
//  4. t.Fatal / t.Fatalf – mark failure, stop test immediately
//  5. Table-driven tests (idiomatic Go testing pattern)
//  6. Subtests with t.Run()
//  7. Test helpers with t.Helper()
//  8. Benchmark functions (BenchmarkXxx)
//  9. Example functions (ExampleXxx) – also serve as docs
// 10. Test setup and teardown (TestMain)
// 11. Skipping tests (t.Skip)
// ============================================================
//
// NOTE: This file contains RUNNABLE demo code showing testing concepts.
// Actual test files must end in _test.go and live alongside the code.
// Run tests with: go test ./...
// Run with race detector: go test -race ./...
// Run benchmarks: go test -bench=. ./...
// Run specific test: go test -run TestAdd ./...
// ============================================================

package main

import "fmt"

// ─────────────────────────────────────────
// Functions that would normally be tested
// ─────────────────────────────────────────

// Add returns the sum of two integers
func Add(a, b int) int {
	return a + b
}

// Divide returns (result, error); errors on divide-by-zero
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// IsPalindrome returns true if s reads the same forwards and backwards
func IsPalindrome(s string) bool {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

// Fibonacci returns the nth Fibonacci number
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// ─────────────────────────────────────────
// main – demonstrate the functions
// (in a real project the functions would be in their own package,
//  and the tests would be in *_test.go files)
// ─────────────────────────────────────────
func main() {
	fmt.Println("═══ Testing Concepts Demo ═══")
	fmt.Println()

	// 2. What a basic test checks
	fmt.Println("─── Add ───")
	fmt.Printf("Add(2, 3)   = %d\n", Add(2, 3))
	fmt.Printf("Add(-1, 1)  = %d\n", Add(-1, 1))
	fmt.Printf("Add(0, 0)   = %d\n", Add(0, 0))

	// 3 & 4. Error path
	fmt.Println("\n─── Divide ───")
	result, err := Divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}

	_, err = Divide(5, 0)
	fmt.Println("Divide by zero error:", err)

	// 5. Table-driven style (used in *_test.go)
	fmt.Println("\n─── IsPalindrome (table-driven style) ───")
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
	}
	for _, tc := range tests {
		got := IsPalindrome(tc.input)
		status := "✓"
		if got != tc.want {
			status = "✗"
		}
		fmt.Printf("  %s IsPalindrome(%q) = %v (want %v)\n",
			status, tc.input, got, tc.want)
	}

	// 8. Fibonacci (used in benchmarks)
	fmt.Println("\n─── Fibonacci ───")
	for _, n := range []int{0, 1, 5, 10, 20} {
		fmt.Printf("  Fib(%2d) = %d\n", n, Fibonacci(n))
	}

	// Print what the actual test files would look like
	fmt.Println()
	fmt.Println("══════════════════════════════════════════")
	fmt.Println("  See testing_test.go for actual test code")
	fmt.Println("══════════════════════════════════════════")
}
