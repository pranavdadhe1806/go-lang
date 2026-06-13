// ============================================================
// TOPIC: Hello World in Go
// ============================================================
// CONCEPTS COVERED:
//  1. Package declaration (package main)
//  2. Import statement (single & grouped)
//  3. The main() entry-point function
//  4. fmt.Print  – prints without a newline
//  5. fmt.Println – prints with a newline
//  6. fmt.Printf  – formatted print (verbs: %s, %d, %v, %T …)
//  7. fmt.Sprintf – returns a formatted string (no print)
//  8. Multi-value fmt.Println
//  9. Escape sequences (\n, \t, \\, \")
// 10. Raw string literals (backtick strings)
// ============================================================

package main

import (
	"fmt"     // standard library package for formatted I/O
	"strings" // standard library for string utilities
)

func main() {

	// ─────────────────────────────────────────
	// 1. Basic Print Variations
	// ─────────────────────────────────────────

	// fmt.Print does NOT add a newline at the end
	fmt.Print("Hello, World!")
	fmt.Print(" (same line)\n") // we manually add \n here

	// fmt.Println always adds a newline at the end
	fmt.Println("Hello, Gopher! 🐹")

	// fmt.Println with multiple arguments – separated by spaces
	fmt.Println("Go", "is", "awesome!")

	// ─────────────────────────────────────────
	// 2. Formatted Printing with fmt.Printf
	// ─────────────────────────────────────────
	// Common format verbs:
	//   %v  – default format of the value
	//   %T  – type of the value
	//   %d  – integer (base 10)
	//   %f  – floating-point
	//   %s  – string
	//   %q  – quoted string
	//   %b  – binary
	//   %x  – hexadecimal

	name := "Gopher"
	age := 10
	version := 1.22

	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
	fmt.Printf("Go version: %.2f\n", version)         // 2 decimal places
	fmt.Printf("Type of name: %T\n", name)            // string
	fmt.Printf("Age in binary: %b, hex: %x\n", age, age)

	// ─────────────────────────────────────────
	// 3. fmt.Sprintf – build a string without printing
	// ─────────────────────────────────────────
	greeting := fmt.Sprintf("Welcome to Go, %s!", name)
	fmt.Println(greeting) // now we print it

	// ─────────────────────────────────────────
	// 4. Escape Sequences
	// ─────────────────────────────────────────
	fmt.Println("Newline:   line1\n\t\tline2 (tabbed)")
	fmt.Println("Quote:     \"Go is great\"")
	fmt.Println("Backslash: C:\\Users\\Gopher")

	// ─────────────────────────────────────────
	// 5. Raw String Literals (backtick ` `)
	// Raw strings preserve everything literally – no escape processing
	// ─────────────────────────────────────────
	raw := `This is a raw string.
It spans multiple lines.
Backslash \n is NOT a newline here.`
	fmt.Println(raw)

	// ─────────────────────────────────────────
	// 6. Using the strings package (bonus)
	// ─────────────────────────────────────────
	upper := strings.ToUpper("hello, go!")
	fmt.Println("Uppercase:", upper)

	repeated := strings.Repeat("Go! ", 3)
	fmt.Println("Repeated:", repeated)
}
