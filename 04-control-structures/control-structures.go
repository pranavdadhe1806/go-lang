package main

import (
	"fmt"
)

func main() {
	fmt.Println("--- If/Else Statements ---")
	age := 20
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}

	// If with a short statement
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	fmt.Println("\n--- For Loops ---")
	// Standard for loop
	fmt.Print("Standard loop: ")
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// For loop as a "while" loop
	fmt.Print("While-style loop: ")
	n := 1
	for n < 5 {
		fmt.Print(n, " ")
		n *= 2
	}
	fmt.Println()

	fmt.Println("\n--- Switch Statements ---")
	day := "Tuesday"
	switch day {
	case "Monday":
		fmt.Println("Start of the work week")
	case "Tuesday", "Wednesday", "Thursday":
		fmt.Println("Mid-week days")
	case "Friday":
		fmt.Println("TGIF!")
	default:
		fmt.Println("It's the weekend!")
	}
}
