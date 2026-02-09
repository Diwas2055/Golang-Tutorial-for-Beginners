// Chapter 01.5: Operators
// This chapter demonstrates the different types of operators in Go:
// 1. Arithmetic Operators
// 2. Comparison Operators
// 3. Logical Operators
// 4. Assignment Operators
package main

import "fmt"

func main() {
	// 1. Arithmetic Operators
	// Used for basic math calculations.
	a, b := 10, 3
	fmt.Printf("Arithmetic: %v + %v = %v\n", a, b, a+b)  // Addition
	fmt.Printf("Arithmetic: %v - %v = %v\n", a, b, a-b)  // Subtraction
	fmt.Printf("Arithmetic: %v * %v = %v\n", a, b, a*b)  // Multiplication
	fmt.Printf("Arithmetic: %v / %v = %v\n", a, b, a/b)  // Division (integer division)
	fmt.Printf("Arithmetic: %v %% %v = %v\n", a, b, a%b) // Modulus (remainder)

	// Increment and Decrement (statements, not expressions in Go)
	count := 1
	count++ // count = count + 1
	fmt.Printf("Increment: count is now %v\n", count)
	count-- // count = count - 1
	fmt.Printf("Decrement: count is now %v\n", count)

	// 2. Comparison Operators
	// Used to compare values, returns a boolean (true/false).
	fmt.Printf("Comparison: %v == %v is %v\n", a, b, a == b) // Equal to
	fmt.Printf("Comparison: %v != %v is %v\n", a, b, a != b) // Not equal to
	fmt.Printf("Comparison: %v > %v is %v\n", a, b, a > b)   // Greater than
	fmt.Printf("Comparison: %v < %v is %v\n", a, b, a < b)   // Less than

	// 3. Logical Operators
	// Used to combine boolean conditions.
	isAdult := true
	hasTicket := false
	fmt.Printf("Logical: %v && %v is %v\n", isAdult, hasTicket, isAdult && hasTicket) // AND
	fmt.Printf("Logical: %v || %v is %v\n", isAdult, hasTicket, isAdult || hasTicket) // OR
	fmt.Printf("Logical: !%v is %v\n", isAdult, !isAdult)                             // NOT

	// 4. Assignment Operators
	// Used to assign or update values of variables.
	x := 5 // Simple assignment
	x += 2 // x = x + 2
	fmt.Printf("Assignment: x after += 2 is %v\n", x)
	x *= 3 // x = x * 3
	fmt.Printf("Assignment: x after *= 3 is %v\n", x)
}
