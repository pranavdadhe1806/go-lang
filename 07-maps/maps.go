// ============================================================
// TOPIC: Maps in Go
// ============================================================
// CONCEPTS COVERED:
//  1. Map declaration (var) – nil map
//  2. Map literal – inline initialization
//  3. make() to create an empty map
//  4. CRUD: add / read / update / delete entries
//  5. Check if a key exists (comma-ok idiom)
//  6. Iterating a map with for-range
//  7. Map with struct values
//  8. Map of slices
//  9. Counting occurrences with a map
// 10. Deleting all keys (clear / re-assign)
// 11. Maps are reference types (shared under the hood)
// 12. Nested maps (map of maps)
// 13. len() on a map
// ============================================================

package main

import (
	"fmt"
	"sort"
)

func main() {

	// ─────────────────────────────────────────
	// 1. Nil map (declaration without initialization)
	// Reading from a nil map is safe (returns zero value)
	// Writing to a nil map causes a panic!
	// ─────────────────────────────────────────
	fmt.Println("═══ 1. Nil map ═══")
	var nilMap map[string]int
	fmt.Printf("nilMap = %v, nil? %v\n", nilMap, nilMap == nil)
	// Accessing a missing key returns zero value – NO panic
	fmt.Printf("nilMap[\"x\"] = %d\n", nilMap["x"])
	// nilMap["x"] = 1  // ← would panic: assignment to entry in nil map

	// ─────────────────────────────────────────
	// 2. Map literal
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 2. Map literal ═══")
	capitals := map[string]string{
		"India":  "New Delhi",
		"USA":    "Washington D.C.",
		"France": "Paris",
		"Japan":  "Tokyo",
	}
	fmt.Println("capitals:", capitals)

	// ─────────────────────────────────────────
	// 3. make() – create a ready-to-use empty map
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 3. make() ═══")
	ages := make(map[string]int)
	ages["Alice"] = 30
	ages["Bob"] = 25
	fmt.Println("ages:", ages)

	// ─────────────────────────────────────────
	// 4. CRUD operations
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 4. CRUD ═══")
	m := map[string]int{"a": 1, "b": 2}

	// Create
	m["c"] = 3
	fmt.Println("After add 'c':", m)

	// Read
	fmt.Printf("Read m[\"a\"] = %d\n", m["a"])

	// Update
	m["a"] = 100
	fmt.Println("After update 'a':", m)

	// Delete
	delete(m, "b")
	fmt.Println("After delete 'b':", m)

	// ─────────────────────────────────────────
	// 5. Comma-ok idiom – check if key exists
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 5. Comma-ok (key existence) ═══")
	val, ok := capitals["India"]
	fmt.Printf("capitals[\"India\"]: val=%q, exists=%v\n", val, ok)

	val, ok = capitals["Germany"]
	fmt.Printf("capitals[\"Germany\"]: val=%q, exists=%v\n", val, ok)

	// Safe access pattern
	if city, found := capitals["France"]; found {
		fmt.Println("Capital of France:", city)
	}

	// ─────────────────────────────────────────
	// 6. Iterating a map with for-range
	// NOTE: Map iteration order is NOT guaranteed in Go
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 6. Iterating (sorted for determinism) ═══")
	keys := make([]string, 0, len(capitals))
	for k := range capitals {
		keys = append(keys, k)
	}
	sort.Strings(keys) // sort keys for consistent output
	for _, k := range keys {
		fmt.Printf("  %s → %s\n", k, capitals[k])
	}

	// ─────────────────────────────────────────
	// 7. Map with struct values
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 7. Map with struct values ═══")
	type Person struct {
		Name string
		Age  int
	}
	people := map[string]Person{
		"emp1": {Name: "Alice", Age: 30},
		"emp2": {Name: "Bob", Age: 25},
	}
	fmt.Println("emp1:", people["emp1"])
	fmt.Println("emp2:", people["emp2"])

	// ─────────────────────────────────────────
	// 8. Map of slices
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 8. Map of slices ═══")
	courses := map[string][]string{
		"Alice": {"Math", "Physics"},
		"Bob":   {"History"},
	}
	courses["Alice"] = append(courses["Alice"], "Chemistry")
	for student, subj := range courses {
		fmt.Printf("  %s: %v\n", student, subj)
	}

	// ─────────────────────────────────────────
	// 9. Counting occurrences
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 9. Counting occurrences ═══")
	words := []string{"go", "rust", "go", "python", "go", "rust"}
	freq := make(map[string]int)
	for _, w := range words {
		freq[w]++ // if key absent, zero-value 0 is used before incrementing
	}
	fmt.Println("Word frequency:", freq)

	// ─────────────────────────────────────────
	// 10. Deleting all entries
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 10. Clear map ═══")
	tiny := map[string]int{"x": 1, "y": 2}
	fmt.Println("Before clear:", tiny)
	for k := range tiny {
		delete(tiny, k)
	}
	fmt.Println("After clear:", tiny)

	// ─────────────────────────────────────────
	// 11. Maps are reference types
	// Both variables point to the SAME underlying map
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 11. Maps are reference types ═══")
	original := map[string]int{"key": 1}
	alias := original // NOT a copy – both point to same map
	alias["key"] = 999
	fmt.Println("original[\"key\"]:", original["key"]) // 999 – both changed

	// ─────────────────────────────────────────
	// 12. Nested maps (map of maps)
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 12. Nested maps ═══")
	config := map[string]map[string]string{
		"database": {
			"host": "localhost",
			"port": "5432",
		},
		"cache": {
			"host": "localhost",
			"port": "6379",
		},
	}
	for section, settings := range config {
		fmt.Printf("[%s]\n", section)
		for k, v := range settings {
			fmt.Printf("  %s = %s\n", k, v)
		}
	}

	// ─────────────────────────────────────────
	// 13. len() on a map
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 13. len() ═══")
	fmt.Printf("Number of capitals: %d\n", len(capitals))
}
