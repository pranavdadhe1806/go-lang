// ============================================================
// TOPIC: Functions in Go
// ============================================================
// CONCEPTS COVERED:
//  1. Basic function – single return value
//  2. Function with multiple parameters
//  3. Multiple return values
//  4. Named return values (naked return)
//  5. Variadic functions (...T)
//  6. Anonymous functions (function literals)
//  7. Immediately Invoked Function Expression (IIFE)
//  8. Functions as first-class values (stored in variables)
//  9. Higher-order functions (passing functions as arguments)
// 10. Returning functions (closures)
// 11. Closures – capturing variables from outer scope
// 12. Recursive functions
// 13. defer inside functions
// 14. init() function – runs before main()
// 15. Blank return (when all named return values are set)
// ============================================================

package main

import (
	"fmt"
	"math"
)

// ─────────────────────────────────────────
// 14. init() – automatically called before main()
// Useful for setup/initialization logic
// ─────────────────────────────────────────
func init() {
	fmt.Println("=== init() runs before main() ===\n")
}

func main() {

	// ─────────────────────────────────────────
	// 1. Basic function – single return
	// ─────────────────────────────────────────
	fmt.Println("═══ 1. Basic function ═══")
	fmt.Println(greet("Gopher"))

	// ─────────────────────────────────────────
	// 2. Function with multiple parameters
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 2. Multiple parameters ═══")
	sum := add(10, 20)
	fmt.Printf("10 + 20 = %d\n", sum)

	// When consecutive params share a type, we can shorten: (a, b int)
	fmt.Printf("3 + 4 = %d\n", addShort(3, 4))

	// ─────────────────────────────────────────
	// 3. Multiple return values
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 3. Multiple return values ═══")
	q, r := divmod(17, 5)
	fmt.Printf("17 / 5 = %d remainder %d\n", q, r)

	// Discard unwanted return values with _
	quotient, _ := divmod(20, 4)
	fmt.Printf("20 / 4 = %d (remainder ignored)\n", quotient)

	// ─────────────────────────────────────────
	// 4. Named return values
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 4. Named return values ═══")
	hyp := hypotenuse(3, 4)
	fmt.Printf("hypotenuse(3,4) = %.2f\n", hyp)

	// ─────────────────────────────────────────
	// 5. Variadic functions
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 5. Variadic functions ═══")
	fmt.Printf("sum(1,2,3)     = %d\n", sumAll(1, 2, 3))
	fmt.Printf("sum(10,20,30,40) = %d\n", sumAll(10, 20, 30, 40))

	// Pass a slice to variadic using spread operator ...
	nums := []int{5, 10, 15, 20}
	fmt.Printf("sum(slice...) = %d\n", sumAll(nums...))

	// ─────────────────────────────────────────
	// 6. Anonymous functions (function literals)
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 6. Anonymous function ═══")
	square := func(n int) int {
		return n * n
	}
	fmt.Printf("square(9) = %d\n", square(9))

	// ─────────────────────────────────────────
	// 7. Immediately Invoked Function Expression (IIFE)
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 7. IIFE ═══")
	result := func(a, b int) int {
		return a * b
	}(6, 7) // called immediately with (6, 7)
	fmt.Printf("6 * 7 = %d\n", result)

	// ─────────────────────────────────────────
	// 8. Functions as first-class values
	// Functions can be stored in variables, slices, maps
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 8. Functions as values ═══")
	ops := map[string]func(int, int) int{
		"add": func(a, b int) int { return a + b },
		"sub": func(a, b int) int { return a - b },
		"mul": func(a, b int) int { return a * b },
	}
	for opName, fn := range ops {
		fmt.Printf("%s(10, 3) = %d\n", opName, fn(10, 3))
	}

	// ─────────────────────────────────────────
	// 9. Higher-order functions – receiving a function as argument
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 9. Higher-order functions ═══")
	doubled := applyToAll([]int{1, 2, 3, 4}, func(n int) int { return n * 2 })
	fmt.Println("Doubled:", doubled)

	squared := applyToAll([]int{1, 2, 3, 4}, func(n int) int { return n * n })
	fmt.Println("Squared:", squared)

	// ─────────────────────────────────────────
	// 10 & 11. Closures – function + captured environment
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 10 & 11. Closures ═══")
	counter := makeCounter()    // counter() remembers its own `count`
	fmt.Println(counter())      // 1
	fmt.Println(counter())      // 2
	fmt.Println(counter())      // 3

	// Each call to makeCounter() creates an INDEPENDENT closure
	counter2 := makeCounter()
	fmt.Println(counter2())     // 1  (fresh counter)

	// Adder closure
	add5 := makeAdder(5)
	add10 := makeAdder(10)
	fmt.Printf("add5(3)  = %d\n", add5(3))
	fmt.Printf("add10(3) = %d\n", add10(3))

	// ─────────────────────────────────────────
	// 12. Recursive functions
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 12. Recursion ═══")
	fmt.Printf("factorial(5) = %d\n", factorial(5))
	fmt.Printf("fibonacci(8) = %d\n", fibonacci(8))

	// ─────────────────────────────────────────
	// 13. defer inside function
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 13. defer ═══")
	greetWithDefer()
}

// ─────────────────────────────────────────
// Function definitions
// ─────────────────────────────────────────

// 1. Basic function
func greet(name string) string {
	return "Hello, " + name + "!"
}

// 2. Multiple parameters – explicit
func add(a int, b int) int {
	return a + b
}

// 2. Multiple parameters – shorthand (shared type)
func addShort(a, b int) int {
	return a + b
}

// 3. Multiple return values
func divmod(a, b int) (int, int) {
	return a / b, a % b
}

// 4. Named return values + naked return
// The return variables are declared in the signature
func hypotenuse(a, b float64) (result float64) {
	result = math.Sqrt(a*a + b*b)
	return // "naked return" – returns named value `result`
}

// 5. Variadic function – accepts zero or more ints
func sumAll(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// 9. Higher-order function – takes a function as argument
func applyToAll(nums []int, transform func(int) int) []int {
	result := make([]int, len(nums))
	for i, n := range nums {
		result[i] = transform(n)
	}
	return result
}

// 10 & 11. Closure – returns a function that captures state
func makeCounter() func() int {
	count := 0 // captured by the returned function
	return func() int {
		count++
		return count
	}
}

// Closure factory – returns an adder with a baked-in value
func makeAdder(x int) func(int) int {
	return func(y int) int {
		return x + y // x is captured from makeAdder's scope
	}
}

// 12. Recursion – factorial
func factorial(n int) int {
	if n <= 1 {
		return 1 // base case
	}
	return n * factorial(n-1) // recursive call
}

// 12. Recursion – fibonacci
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 13. defer inside a named function
func greetWithDefer() {
	defer fmt.Println("  defer: Goodbye!")   // runs last
	fmt.Println("  Hello from greetWithDefer")
}
