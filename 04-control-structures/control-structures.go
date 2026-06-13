// ============================================================
// TOPIC: Control Structures in Go
// ============================================================
// CONCEPTS COVERED:
//  1. if / else if / else
//  2. if with short initializer statement
//  3. Nested if
//  4. for – standard 3-clause loop (like C)
//  5. for – while-style (condition only)
//  6. for – infinite loop with break
//  7. for – range over slice
//  8. for – range over string (runes)
//  9. for – range over map
// 10. continue & break statements
// 11. Labeled break (breaking outer loops)
// 12. switch – basic value matching
// 13. switch – multiple values per case
// 14. switch – with short initializer
// 15. switch – without expression (replaces if-else chains)
// 16. switch – type switch
// 17. fallthrough keyword
// 18. goto statement (rarely used, shown for completeness)
// 19. defer statement (execution order)
// 20. defer with a loop (LIFO order)
// ============================================================

package main

import "fmt"

func main() {

	// ─────────────────────────────────────────
	// 1. Basic if / else if / else
	// ─────────────────────────────────────────
	fmt.Println("═══ 1. if / else if / else ═══")
	age := 20
	if age >= 18 {
		fmt.Println("Adult")
	} else if age >= 13 {
		fmt.Println("Teenager")
	} else {
		fmt.Println("Child")
	}

	// ─────────────────────────────────────────
	// 2. if with a short initializer statement
	// The variable (num) is scoped ONLY to the if block
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 2. if with short initializer ═══")
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Printf("%d has 1 digit\n", num)
	} else {
		fmt.Printf("%d has multiple digits\n", num)
	}
	// fmt.Println(num) // ERROR: num is out of scope here

	// ─────────────────────────────────────────
	// 3. Nested if
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 3. Nested if ═══")
	score := 85
	if score >= 50 {
		if score >= 90 {
			fmt.Println("Grade: A")
		} else if score >= 75 {
			fmt.Println("Grade: B")
		} else {
			fmt.Println("Grade: C")
		}
	} else {
		fmt.Println("Grade: F (Fail)")
	}

	// ─────────────────────────────────────────
	// 4. Standard for loop (3-clause)
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 4. Standard for loop ═══")
	for i := 0; i < 5; i++ {
		fmt.Printf("i = %d\n", i)
	}

	// ─────────────────────────────────────────
	// 5. for as while loop (condition only)
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 5. for as while loop ═══")
	n := 1
	for n < 64 {
		fmt.Printf("n = %d\n", n)
		n *= 2
	}

	// ─────────────────────────────────────────
	// 6. Infinite loop with break
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 6. Infinite loop with break ═══")
	count := 0
	for {
		if count >= 3 {
			break // exit the loop
		}
		fmt.Printf("count = %d\n", count)
		count++
	}

	// ─────────────────────────────────────────
	// 10. continue & break
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 10. continue & break ═══")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // skip even numbers
		}
		if i == 7 {
			break // stop at 7
		}
		fmt.Printf("odd: %d\n", i)
	}

	// ─────────────────────────────────────────
	// 7. for range over slice
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 7. for range over slice ═══")
	fruits := []string{"apple", "banana", "cherry"}
	for index, fruit := range fruits {
		fmt.Printf("index %d: %s\n", index, fruit)
	}

	// Ignore index with _
	fmt.Println("-- only values --")
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	// ─────────────────────────────────────────
	// 8. for range over string (iterates runes, not bytes)
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 8. for range over string ═══")
	for i, ch := range "Go🔥" {
		fmt.Printf("byte-pos %d: %c  (rune value: %d)\n", i, ch, ch)
	}

	// ─────────────────────────────────────────
	// 9. for range over map
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 9. for range over map ═══")
	capitals := map[string]string{
		"India":  "New Delhi",
		"USA":    "Washington D.C.",
		"France": "Paris",
	}
	for country, capital := range capitals {
		fmt.Printf("%s → %s\n", country, capital)
	}

	// ─────────────────────────────────────────
	// 11. Labeled break – break out of outer loop
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 11. Labeled break ═══")
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				fmt.Println("Breaking outer loop at i=1,j=1")
				break outer // exits BOTH loops
			}
			fmt.Printf("  i=%d, j=%d\n", i, j)
		}
	}

	// ─────────────────────────────────────────
	// 12. switch – basic value matching
	// No need for break in Go! Falls through is NOT default.
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 12. Basic switch ═══")
	day := "Tuesday"
	switch day {
	case "Monday":
		fmt.Println("Start of work week")
	case "Tuesday":
		fmt.Println("Second day")
	case "Friday":
		fmt.Println("TGIF!")
	default:
		fmt.Println("Another day")
	}

	// ─────────────────────────────────────────
	// 13. switch – multiple values per case
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 13. Multiple values per case ═══")
	switch day {
	case "Saturday", "Sunday":
		fmt.Println("Weekend!")
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday")
	}

	// ─────────────────────────────────────────
	// 14. switch with short initializer
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 14. switch with initializer ═══")
	switch x := 42; {
	case x < 0:
		fmt.Println("negative")
	case x == 0:
		fmt.Println("zero")
	case x > 0:
		fmt.Println("positive")
	}

	// ─────────────────────────────────────────
	// 15. switch – without expression (replaces if-else chains)
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 15. switch without expression ═══")
	temp := 35
	switch {
	case temp < 0:
		fmt.Println("Freezing")
	case temp < 20:
		fmt.Println("Cold")
	case temp < 30:
		fmt.Println("Comfortable")
	default:
		fmt.Println("Hot!")
	}

	// ─────────────────────────────────────────
	// 16. Type switch – identify interface's dynamic type
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 16. Type switch ═══")
	values := []interface{}{42, "hello", true, 3.14}
	for _, v := range values {
		switch t := v.(type) {
		case int:
			fmt.Printf("int: %d\n", t)
		case string:
			fmt.Printf("string: %q\n", t)
		case bool:
			fmt.Printf("bool: %v\n", t)
		default:
			fmt.Printf("other type: %T = %v\n", t, t)
		}
	}

	// ─────────────────────────────────────────
	// 17. fallthrough – force next case to execute
	// (opposite of default behavior in Go)
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 17. fallthrough ═══")
	val := 1
	switch val {
	case 1:
		fmt.Println("case 1")
		fallthrough // explicitly falls into case 2
	case 2:
		fmt.Println("case 2 (via fallthrough)")
	case 3:
		fmt.Println("case 3")
	}

	// ─────────────────────────────────────────
	// 19. defer – deferred call runs after surrounding function returns
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 19. defer ═══")
	fmt.Println("start")
	defer fmt.Println("deferred: runs last")
	fmt.Println("end")

	// ─────────────────────────────────────────
	// 20. Multiple defers – execute in LIFO (stack) order
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 20. Multiple defers (LIFO order) ═══")
	for i := 0; i < 3; i++ {
		defer fmt.Printf("deferred %d\n", i)
	}
	fmt.Println("After the loop (defers fire after main returns)")
	// Output order will be: 2, 1, 0
}
