// ============================================================
// TOPIC: Error Handling in Go
// ============================================================
// CONCEPTS COVERED:
//  1. The built-in error interface
//  2. errors.New() – create simple errors
//  3. fmt.Errorf() – format errors with context
//  4. Returning and checking errors (idiomatic Go pattern)
//  5. Custom error types (struct implementing error)
//  6. Sentinel errors – predefined package-level errors
//  7. errors.Is() – check error chain for a specific sentinel
//  8. errors.As() – extract a specific error type from the chain
//  9. Error wrapping with %w (Go 1.13+)
// 10. Unwrapping errors with errors.Unwrap()
// 11. panic – unrecoverable errors / programming bugs
// 12. recover() – catch a panic in a deferred function
// 13. panic + recover pattern (for safe wrappers)
// 14. Multiple error values & first-error wins
// ============================================================

package main

import (
	"errors"
	"fmt"
	"strconv"
)

// ─────────────────────────────────────────
// 6. Sentinel errors – predefined at package level
// Callers compare against these using errors.Is()
// ─────────────────────────────────────────
var (
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrDivByZero    = errors.New("division by zero")
)

// ─────────────────────────────────────────
// 5. Custom error type
// Any type with an Error() string method satisfies the error interface
// ─────────────────────────────────────────
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed for %q: %s", e.Field, e.Message)
}

// ─────────────────────────────────────────
// Another custom error with wrapped cause
// ─────────────────────────────────────────
type DBError struct {
	Code    int
	Message string
	Cause   error
}

func (e *DBError) Error() string {
	return fmt.Sprintf("DB error %d: %s", e.Code, e.Message)
}

func (e *DBError) Unwrap() error {
	return e.Cause // enables errors.Is / errors.As to traverse the chain
}

func main() {

	// ─────────────────────────────────────────
	// 1 & 2. Basic error creation and handling
	// ─────────────────────────────────────────
	fmt.Println("═══ 1 & 2. Basic error handling ═══")
	result, err := safeDivide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = safeDivide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 2 = %d\n", result)
	}

	// ─────────────────────────────────────────
	// 3. fmt.Errorf – add context to an error
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 3. fmt.Errorf ═══")
	_, err = parseAge("-5")
	if err != nil {
		fmt.Println("Error:", err)
	}

	_, err = parseAge("abc")
	if err != nil {
		fmt.Println("Error:", err)
	}

	age, err := parseAge("25")
	if err == nil {
		fmt.Println("Parsed age:", age)
	}

	// ─────────────────────────────────────────
	// 4. Idiomatic Go error pattern
	// Check immediately, handle or return early
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 4. Idiomatic error pattern ═══")
	err = processUser("") // empty name triggers validation error
	if err != nil {
		fmt.Println("processUser error:", err)
	}

	// ─────────────────────────────────────────
	// 5. Custom error type
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 5. Custom error type ═══")
	err = validateAge(-1)
	if err != nil {
		fmt.Println("Custom error:", err)
	}

	// ─────────────────────────────────────────
	// 6 & 7. Sentinel errors and errors.Is()
	// errors.Is() checks the entire error chain
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 6 & 7. Sentinel errors + errors.Is() ═══")
	err = findUser(999)
	fmt.Println("err:", err)
	fmt.Println("Is ErrNotFound?", errors.Is(err, ErrNotFound)) // true

	err = findUser(1)
	fmt.Println("err:", err)
	fmt.Println("Is ErrNotFound?", errors.Is(err, ErrNotFound)) // false

	// ─────────────────────────────────────────
	// 8. errors.As() – extract concrete error type from chain
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 8. errors.As() ═══")
	err = validateAge(-5)
	var valErr *ValidationError
	if errors.As(err, &valErr) {
		fmt.Printf("Caught ValidationError – field: %q, msg: %q\n",
			valErr.Field, valErr.Message)
	}

	// ─────────────────────────────────────────
	// 9 & 10. Error wrapping with %w and Unwrap()
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 9 & 10. Wrapping + Unwrap ═══")
	err = wrappedOperation()
	fmt.Println("Wrapped error:", err)
	fmt.Println("Is ErrNotFound?", errors.Is(err, ErrNotFound))      // true via chain
	fmt.Println("Unwrapped:", errors.Unwrap(err))                     // one level up

	// errors.As traverses the chain
	var dbErr *DBError
	if errors.As(err, &dbErr) {
		fmt.Printf("DBError: code=%d, msg=%s\n", dbErr.Code, dbErr.Message)
	}

	// ─────────────────────────────────────────
	// 11. panic – signals unrecoverable state
	// Only use for programming bugs, never for business logic
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 11 & 12. panic + recover ═══")
	safePanic() // wrapped so the program doesn't crash

	// ─────────────────────────────────────────
	// 13. Safe wrapper with recover
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 13. safeRun wrapper ═══")
	err = safeRun(func() {
		mustDivide(10, 0) // will panic
	})
	if err != nil {
		fmt.Println("Recovered error:", err)
	}

	err = safeRun(func() {
		fmt.Println("No panic here, result:", mustDivide(10, 2))
	})
	if err != nil {
		fmt.Println("Recovered error:", err)
	}
}

// ─────────────────────────────────────────
// Helper functions
// ─────────────────────────────────────────

// 2. Returns (value, error) – the canonical Go pattern
func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivByZero
	}
	return a / b, nil
}

// 3. Using fmt.Errorf to wrap underlying errors with context
func parseAge(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		// wrap the underlying strconv error with context
		return 0, fmt.Errorf("parseAge: cannot convert %q: %w", s, err)
	}
	if n < 0 || n > 150 {
		return 0, fmt.Errorf("parseAge: age %d is out of valid range [0,150]", n)
	}
	return n, nil
}

// 4. Chained function calls with early return on error
func processUser(name string) error {
	if name == "" {
		return fmt.Errorf("processUser: %w", ErrUnauthorized)
	}
	// more processing…
	return nil
}

// 5. Returns a custom error type
func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Message: "must be non-negative"}
	}
	return nil
}

// 6. Function that returns a sentinel error
func findUser(id int) error {
	if id != 1 {
		// Wrap sentinel so errors.Is still works via chain
		return fmt.Errorf("findUser(id=%d): %w", id, ErrNotFound)
	}
	return nil
}

// 9. Multi-level wrapping
func wrappedOperation() error {
	inner := ErrNotFound
	dbErr := &DBError{Code: 500, Message: "query failed", Cause: inner}
	return fmt.Errorf("wrappedOperation: %w", dbErr)
}

// 11. mustDivide panics instead of returning an error
func mustDivide(a, b int) int {
	if b == 0 {
		panic("mustDivide: division by zero")
	}
	return a / b
}

// 12. recover() inside a deferred call
func safePanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	panic("something went terribly wrong!")
}

// 13. Generic safe wrapper – catches panics and returns as errors
func safeRun(fn func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic caught: %v", r)
		}
	}()
	fn()
	return nil
}
