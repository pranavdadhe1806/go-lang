// ============================================================
// TOPIC: Structs in Go
// ============================================================
// CONCEPTS COVERED:
//  1. Basic struct declaration
//  2. Struct instantiation – field names
//  3. Struct instantiation – positional (not recommended)
//  4. Struct zero value
//  5. Accessing and modifying fields
//  6. Pointer to struct & automatic dereferencing
//  7. Anonymous structs (inline, one-off usage)
//  8. Struct embedding (composition over inheritance)
//  9. Promoted fields from embedded struct
// 10. Struct with methods (see also: 10-methods)
// 11. Struct tags (used with encoding/json, db ORM, etc.)
// 12. Comparing structs
// 13. Struct as function arguments (value copy vs pointer)
// 14. Constructor function pattern (idiomatic Go)
// 15. Nested structs
// ============================================================

package main

import (
	"encoding/json"
	"fmt"
)

// ─────────────────────────────────────────
// 1. Basic struct declaration
// ─────────────────────────────────────────
type Person struct {
	Name string
	Age  int
	City string
}

// ─────────────────────────────────────────
// 11. Struct tags – metadata for external packages
// encoding/json uses these to map fields to JSON keys
// ─────────────────────────────────────────
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price,omitempty"` // omit if zero
}

// ─────────────────────────────────────────
// 8. Embedded structs (composition)
// Address is embedded inside Employee
// ─────────────────────────────────────────
type Address struct {
	Street string
	City   string
	Zip    string
}

type Employee struct {
	Person  // embedded – promotes all Person fields
	Address // embedded – promotes all Address fields
	Company string
	Salary  float64
}

// ─────────────────────────────────────────
// 15. Nested structs (explicit, not embedded)
// ─────────────────────────────────────────
type Order struct {
	ID       int
	Customer Person  // named field, not embedded
	Total    float64
}

// ─────────────────────────────────────────
// 14. Constructor function (idiomatic pattern)
// ─────────────────────────────────────────
func NewPerson(name string, age int, city string) *Person {
	return &Person{
		Name: name,
		Age:  age,
		City: city,
	}
}

// ─────────────────────────────────────────
// 10. Method on a struct (value receiver)
// ─────────────────────────────────────────
func (p Person) Greet() string {
	return fmt.Sprintf("Hi, I'm %s from %s, age %d.", p.Name, p.City, p.Age)
}

// Method with pointer receiver (can modify the struct)
func (p *Person) HaveBirthday() {
	p.Age++
}

func main() {

	// ─────────────────────────────────────────
	// 2. Struct instantiation with field names
	// ─────────────────────────────────────────
	fmt.Println("═══ 2. Struct with field names ═══")
	p1 := Person{Name: "Alice", Age: 30, City: "Mumbai"}
	fmt.Println("p1:", p1)

	// ─────────────────────────────────────────
	// 3. Positional instantiation (not recommended)
	// Must supply ALL fields in order
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 3. Positional instantiation ═══")
	p2 := Person{"Bob", 25, "Delhi"}
	fmt.Println("p2:", p2)

	// ─────────────────────────────────────────
	// 4. Zero value struct
	// Each field gets its zero value (0, "", false)
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 4. Zero value struct ═══")
	var p3 Person // all fields zero-valued
	fmt.Printf("p3: %+v\n", p3) // %+v prints field names

	// ─────────────────────────────────────────
	// 5. Accessing and modifying fields
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 5. Accessing & modifying fields ═══")
	p1.City = "Pune"
	fmt.Printf("Updated p1.City = %s\n", p1.City)
	fmt.Printf("p1.Name = %s, p1.Age = %d\n", p1.Name, p1.Age)

	// ─────────────────────────────────────────
	// 6. Pointer to struct
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 6. Pointer to struct ═══")
	pp := &Person{Name: "Charlie", Age: 22, City: "Bangalore"}
	// Go automatically dereferences: pp.Name == (*pp).Name
	fmt.Println("Name via pointer:", pp.Name)
	pp.Age = 23 // auto-dereference
	fmt.Println("Updated age:", pp.Age)

	// ─────────────────────────────────────────
	// 7. Anonymous structs
	// Useful for one-off data shapes, test fixtures, JSON decoding
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 7. Anonymous struct ═══")
	config := struct {
		Host string
		Port int
	}{
		Host: "localhost",
		Port: 8080,
	}
	fmt.Printf("config: %+v\n", config)

	// ─────────────────────────────────────────
	// 8 & 9. Struct embedding and promoted fields
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 8 & 9. Embedding & promoted fields ═══")
	emp := Employee{
		Person:  Person{Name: "Diana", Age: 28, City: "Hyderabad"},
		Address: Address{Street: "MG Road", City: "Hyderabad", Zip: "500001"},
		Company: "GoTech",
		Salary:  75000,
	}

	// Promoted fields – access directly without specifying embedded struct
	fmt.Println("Name (promoted):", emp.Name)       // emp.Person.Name
	fmt.Println("Street (promoted):", emp.Street)   // emp.Address.Street
	fmt.Println("Company:", emp.Company)
	// When fields conflict, use explicit path
	fmt.Println("emp.Address.City:", emp.Address.City)

	// ─────────────────────────────────────────
	// 10. Using methods
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 10. Methods ═══")
	alice := Person{Name: "Alice", Age: 30, City: "Mumbai"}
	fmt.Println(alice.Greet())
	alice.HaveBirthday() // pointer receiver – modifies alice
	fmt.Printf("After birthday: %d\n", alice.Age)

	// ─────────────────────────────────────────
	// 11. Struct tags with encoding/json
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 11. Struct tags / JSON ═══")
	prod := Product{ID: 1, Name: "Laptop", Price: 999.99}
	jsonBytes, _ := json.Marshal(prod)
	fmt.Println("JSON:", string(jsonBytes))

	// Unmarshal JSON back to struct
	jsonStr := `{"id":2,"name":"Phone"}`
	var prod2 Product
	json.Unmarshal([]byte(jsonStr), &prod2)
	fmt.Printf("Unmarshaled: %+v\n", prod2) // Price is 0 (omitempty skips it in output)

	// ─────────────────────────────────────────
	// 12. Comparing structs
	// Two structs are equal if all their fields are equal
	// Structs with incomparable fields (slice, map) can NOT be compared with ==
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 12. Comparing structs ═══")
	pa := Person{Name: "Alice", Age: 30, City: "Mumbai"}
	pb := Person{Name: "Alice", Age: 30, City: "Mumbai"}
	pc := Person{Name: "Bob", Age: 25, City: "Delhi"}
	fmt.Printf("pa == pb: %v\n", pa == pb) // true
	fmt.Printf("pa == pc: %v\n", pa == pc) // false

	// ─────────────────────────────────────────
	// 13. Struct as function argument
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 13. Struct as argument ═══")
	original := Person{Name: "Eve", Age: 20, City: "Chennai"}
	makeOlder(original)        // value copy – original unchanged
	fmt.Println("After makeOlder (value):", original.Age) // still 20

	makeOlderPtr(&original)   // pointer – original IS changed
	fmt.Println("After makeOlderPtr (pointer):", original.Age) // 21

	// ─────────────────────────────────────────
	// 14. Constructor function
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 14. Constructor ═══")
	frank := NewPerson("Frank", 35, "Kolkata")
	fmt.Printf("Constructed: %+v\n", *frank)

	// ─────────────────────────────────────────
	// 15. Nested structs (explicit field)
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 15. Nested struct ═══")
	order := Order{
		ID:       101,
		Customer: Person{Name: "Grace", Age: 27, City: "Jaipur"},
		Total:    249.99,
	}
	fmt.Printf("Order #%d by %s, Total: $%.2f\n",
		order.ID, order.Customer.Name, order.Total)
}

// Helper functions for concept 13
func makeOlder(p Person) {
	p.Age++ // changes only the local copy
}

func makeOlderPtr(p *Person) {
	p.Age++ // changes the original via pointer
}
