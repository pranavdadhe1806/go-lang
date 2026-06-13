// ============================================================
// TOPIC: Methods in Go
// ============================================================
// CONCEPTS COVERED:
//  1. Method vs function – syntax difference
//  2. Value receiver method
//  3. Pointer receiver method (can mutate)
//  4. Methods on non-struct types (custom type)
//  5. Method set rules (value vs pointer)
//  6. Chaining methods (fluent interface)
//  7. Embedding and method promotion
//  8. Method expressions (function value from method)
//  9. Method values (bound method)
// 10. Stringer interface (fmt.Stringer) via String() method
// ============================================================

package main

import (
	"fmt"
	"math"
)

// ─────────────────────────────────────────
// 1. Types used in examples
// ─────────────────────────────────────────

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

// 4. Method on a non-struct custom type
type Celsius float64
type Fahrenheit float64

// ─────────────────────────────────────────
// 2. Value receiver – does NOT modify the receiver
// Use when the method doesn't need to mutate state
// ─────────────────────────────────────────

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// ─────────────────────────────────────────
// 10. Stringer interface – any type with String() string
// is automatically used by fmt package for printing
// ─────────────────────────────────────────
func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle(%.1f x %.1f)", r.Width, r.Height)
}

// ─────────────────────────────────────────
// 3. Pointer receiver – MODIFIES the receiver
// Use when method must mutate the struct
// ─────────────────────────────────────────

func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// ─────────────────────────────────────────
// 6. Chaining – each method returns *Rectangle
// (fluent interface pattern)
// ─────────────────────────────────────────

func (r *Rectangle) SetWidth(w float64) *Rectangle {
	r.Width = w
	return r
}

func (r *Rectangle) SetHeight(h float64) *Rectangle {
	r.Height = h
	return r
}

// Circle methods
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle(r=%.1f)", c.Radius)
}

// ─────────────────────────────────────────
// 4. Methods on custom non-struct types
// ─────────────────────────────────────────

func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func (f Fahrenheit) ToCelsius() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%.2f°C", float64(c))
}

// ─────────────────────────────────────────
// 7. Embedding and method promotion
// ─────────────────────────────────────────

type Animal struct {
	Name string
}

func (a Animal) Speak() string {
	return a.Name + " makes a sound."
}

type Dog struct {
	Animal  // embedded – Speak() is promoted
	Breed string
}

// Dog can override the promoted method
func (d Dog) Speak() string {
	return d.Name + " says: Woof!"
}

// ─────────────────────────────────────────
// Builder pattern using chained pointer receivers
// ─────────────────────────────────────────

type QueryBuilder struct {
	table  string
	where  string
	limit  int
}

func NewQuery(table string) *QueryBuilder {
	return &QueryBuilder{table: table}
}

func (q *QueryBuilder) Where(cond string) *QueryBuilder {
	q.where = cond
	return q
}

func (q *QueryBuilder) Limit(n int) *QueryBuilder {
	q.limit = n
	return q
}

func (q *QueryBuilder) Build() string {
	sql := "SELECT * FROM " + q.table
	if q.where != "" {
		sql += " WHERE " + q.where
	}
	if q.limit > 0 {
		sql += fmt.Sprintf(" LIMIT %d", q.limit)
	}
	return sql
}

func main() {

	// ─────────────────────────────────────────
	// 2. Value receiver methods
	// ─────────────────────────────────────────
	fmt.Println("═══ 2. Value receiver ═══")
	rect := Rectangle{Width: 5, Height: 3}
	fmt.Printf("Area      : %.2f\n", rect.Area())
	fmt.Printf("Perimeter : %.2f\n", rect.Perimeter())

	// ─────────────────────────────────────────
	// 3. Pointer receiver methods
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 3. Pointer receiver (mutating) ═══")
	fmt.Printf("Before scale: %v\n", rect)
	rect.Scale(2) // Go auto-takes address: (&rect).Scale(2)
	fmt.Printf("After scale : %v\n", rect)

	// ─────────────────────────────────────────
	// 4. Methods on custom types
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 4. Methods on custom types ═══")
	boiling := Celsius(100)
	fmt.Printf("%v → %v°F\n", boiling, boiling.ToFahrenheit())

	freezing := Fahrenheit(32)
	fmt.Printf("%.1f°F → %v\n", float64(freezing), freezing.ToCelsius())

	// ─────────────────────────────────────────
	// 5. Value vs pointer method set
	// A *T has access to both value and pointer receiver methods
	// A  T  has access ONLY to value receiver methods
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 5. Method set ═══")
	r1 := Rectangle{Width: 4, Height: 2}    // value
	r2 := &Rectangle{Width: 4, Height: 2}   // pointer

	r1.Scale(3)  // Go implicitly takes &r1 – works
	r2.Scale(3)  // direct pointer – works
	fmt.Printf("r1 area after scale: %.2f\n", r1.Area())
	fmt.Printf("r2 area after scale: %.2f\n", r2.Area())

	// ─────────────────────────────────────────
	// 6. Method chaining (fluent interface)
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 6. Method chaining ═══")
	r := new(Rectangle)
	r.SetWidth(10).SetHeight(5) // each call returns *Rectangle
	fmt.Printf("Chained rect area: %.2f\n", r.Area())

	// ─────────────────────────────────────────
	// Fluent builder pattern
	// ─────────────────────────────────────────
	query := NewQuery("users").Where("age > 18").Limit(10).Build()
	fmt.Println("\nSQL:", query)

	// ─────────────────────────────────────────
	// 7. Embedding and method promotion / override
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 7. Embedding & promotion ═══")
	genericAnimal := Animal{Name: "Cat"}
	fmt.Println(genericAnimal.Speak()) // Animal method

	dog := Dog{Animal: Animal{Name: "Rex"}, Breed: "Labrador"}
	fmt.Println(dog.Speak())           // Dog's overridden method
	fmt.Println(dog.Animal.Speak())    // Access original via explicit path

	// ─────────────────────────────────────────
	// 8. Method expressions
	// Treat a method as a function value; receiver is first argument
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 8. Method expression ═══")
	areaFn := Rectangle.Area // type: func(Rectangle) float64
	fmt.Printf("areaFn(rect): %.2f\n", areaFn(rect))

	// ─────────────────────────────────────────
	// 9. Method values (bound to a specific receiver)
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 9. Method value (bound) ═══")
	c := Circle{Radius: 5}
	boundArea := c.Area // bound to c; no receiver argument needed
	fmt.Printf("boundArea(): %.2f\n", boundArea())

	// ─────────────────────────────────────────
	// 10. fmt.Stringer – automatic via String() method
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 10. fmt.Stringer ═══")
	fmt.Println(rect)     // calls rect.String() automatically
	fmt.Println(c)        // calls c.String() automatically
}
