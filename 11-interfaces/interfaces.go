// ============================================================
// TOPIC: Interfaces in Go
// ============================================================
// CONCEPTS COVERED:
//  1. Interface declaration
//  2. Implicit implementation (no 'implements' keyword)
//  3. Interface as function parameter (polymorphism)
//  4. Empty interface (interface{} / any)
//  5. Type assertion – extract concrete type from interface
//  6. Type assertion with comma-ok (safe)
//  7. Type switch
//  8. Interface embedding (composing interfaces)
//  9. Standard library interfaces: fmt.Stringer, io.Reader, error
// 10. Nil interface vs nil concrete type (gotcha!)
// 11. Interface comparison
// 12. Checking if a value implements an interface (compile-time)
// ============================================================

package main

import (
	"fmt"
	"math"
	"strings"
)

// ─────────────────────────────────────────
// 1. Interface declaration
// An interface is a set of method signatures
// ─────────────────────────────────────────

type Shape interface {
	Area() float64
	Perimeter() float64
}

// ─────────────────────────────────────────
// 8. Interface embedding – compose from smaller interfaces
// ─────────────────────────────────────────

type Namer interface {
	Name() string
}

type Describer interface {
	Namer           // embeds Namer
	Describe() string
}

// ─────────────────────────────────────────
// 9. Standard library interfaces
// ─────────────────────────────────────────

// fmt.Stringer – any type with String() string
// error        – any type with Error() string

// ─────────────────────────────────────────
// 2. Concrete types that implicitly implement Shape
// No 'implements' keyword required – duck typing
// ─────────────────────────────────────────

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Circle also implements fmt.Stringer
func (c Circle) String() string {
	return fmt.Sprintf("Circle(r=%.2f)", c.Radius)
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Triangle struct {
	A, B, C float64 // three sides
}

func (t Triangle) Area() float64 {
	s := (t.A + t.B + t.C) / 2 // semi-perimeter (Heron's formula)
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

// ─────────────────────────────────────────
// 8. Types implementing composed interface
// ─────────────────────────────────────────

type Animal struct {
	species string
}

func (a Animal) Name() string {
	return a.species
}

func (a Animal) Describe() string {
	return fmt.Sprintf("I am a %s.", a.species)
}

// ─────────────────────────────────────────
// 9. Custom error type (implements error interface)
// ─────────────────────────────────────────

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error on %q: %s", e.Field, e.Message)
}

// ─────────────────────────────────────────
// 3. Function that accepts any Shape (polymorphism)
// ─────────────────────────────────────────

func printShape(s Shape) {
	fmt.Printf("  %T → Area=%.2f, Perimeter=%.2f\n", s, s.Area(), s.Perimeter())
}

func totalArea(shapes []Shape) float64 {
	total := 0.0
	for _, s := range shapes {
		total += s.Area()
	}
	return total
}

func main() {

	// ─────────────────────────────────────────
	// 2 & 3. Implicit implementation & polymorphism
	// ─────────────────────────────────────────
	fmt.Println("═══ 2 & 3. Implicit implementation & polymorphism ═══")
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 4, Height: 6},
		Triangle{A: 3, B: 4, C: 5},
	}
	for _, s := range shapes {
		printShape(s)
	}
	fmt.Printf("Total area: %.2f\n", totalArea(shapes))

	// ─────────────────────────────────────────
	// 4. Empty interface – holds any value
	// interface{} is the same as 'any' (Go 1.18+)
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 4. Empty interface (any) ═══")
	var anything interface{}
	anything = 42
	fmt.Println("int:   ", anything)
	anything = "hello"
	fmt.Println("string:", anything)
	anything = Circle{Radius: 3}
	fmt.Println("Circle:", anything)

	// Slice of any type
	mixed := []any{1, "two", true, 3.14, Circle{Radius: 2}}
	for _, v := range mixed {
		fmt.Printf("  %T: %v\n", v, v)
	}

	// ─────────────────────────────────────────
	// 5 & 6. Type assertion
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 5. Type assertion ═══")
	var s Shape = Circle{Radius: 7}

	// Direct assertion (panics if wrong type!)
	c := s.(Circle)
	fmt.Printf("Asserted Circle radius: %.2f\n", c.Radius)

	// Safe assertion with comma-ok
	fmt.Println("\n═══ 6. Safe type assertion (comma-ok) ═══")
	if circle, ok := s.(Circle); ok {
		fmt.Printf("It's a Circle with radius %.2f\n", circle.Radius)
	}
	if _, ok := s.(Rectangle); !ok {
		fmt.Println("It's NOT a Rectangle")
	}

	// ─────────────────────────────────────────
	// 7. Type switch
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 7. Type switch ═══")
	values := []interface{}{42, "hello", true, 3.14, Circle{Radius: 1}}
	for _, v := range values {
		switch t := v.(type) {
		case int:
			fmt.Printf("int: %d\n", t)
		case string:
			fmt.Printf("string: %q (upper: %s)\n", t, strings.ToUpper(t))
		case bool:
			fmt.Printf("bool: %v\n", t)
		case Circle:
			fmt.Printf("Circle: area=%.2f\n", t.Area())
		default:
			fmt.Printf("unknown type %T: %v\n", t, t)
		}
	}

	// ─────────────────────────────────────────
	// 8. Interface embedding
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 8. Interface embedding ═══")
	var d Describer = Animal{species: "Dog"}
	fmt.Println(d.Name())
	fmt.Println(d.Describe())

	// ─────────────────────────────────────────
	// 9. Standard library interfaces
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 9. fmt.Stringer and error ═══")
	// fmt.Stringer
	circ := Circle{Radius: 4}
	fmt.Println(circ) // automatically calls circ.String()

	// Custom error
	err := &ValidationError{Field: "email", Message: "invalid format"}
	fmt.Println("Error:", err)
	// The error interface is just: type error interface { Error() string }

	// ─────────────────────────────────────────
	// 10. Nil interface gotcha
	// A nil interface{} is NOT the same as an interface holding a nil pointer
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 10. Nil interface gotcha ═══")
	var i1 interface{} = nil              // truly nil
	var p *Circle = nil
	var i2 interface{} = p               // interface holds (type=*Circle, value=nil)

	fmt.Printf("i1 == nil: %v\n", i1 == nil) // true
	fmt.Printf("i2 == nil: %v\n", i2 == nil) // FALSE! interface is not nil
	fmt.Printf("i2 type: %T, value: %v\n", i2, i2)

	// ─────────────────────────────────────────
	// 12. Compile-time interface check
	// Assign to _ using the interface type to verify implementation
	// ─────────────────────────────────────────
	// This line causes a compile error if Circle does NOT implement Shape
	var _ Shape = Circle{}       // value receiver check
	var _ Shape = Rectangle{}    // value receiver check
	fmt.Println("\n═══ 12. Compile-time interface check ═══")
	fmt.Println("Circle and Rectangle correctly implement Shape (verified at compile time)")
}
