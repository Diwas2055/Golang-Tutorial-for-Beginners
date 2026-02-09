// Chapter 01.5: Operators
// This chapter demonstrates the different types of operators in Go:
// 1. Arithmetic Operators
// 2. Comparison Operators
// 3. Logical Operators
// 4. Assignment & Short Declaration
// 5. Pointer & Selector Operators
package main

import "fmt"

func main() {
	// 1. Arithmetic Operators
	a, b := 10, 3
	fmt.Printf("Arithmetic: %v + %v = %v\n", a, b, a+b)
	fmt.Printf("Arithmetic: %v %% %v = %v (Modulus)\n", a, b, a%b)

	// 2. Comparison Operators
	fmt.Printf("Comparison: %v == %v is %v\n", a, b, a == b)
	fmt.Printf("Comparison: %v != %v is %v\n", a, b, a != b)

	// 3. Logical Operators
	isAdult, hasTicket := true, false
	fmt.Printf("Logical: %v && %v is %v\n", isAdult, hasTicket, isAdult && hasTicket)
	fmt.Printf("Logical: %v || %v is %v\n", isAdult, hasTicket, isAdult || hasTicket)

	// 4. Assignment & Short Declaration
	// ':=' declares and initializes a variable in one step.
	x := 20
	x += 5 // x = x + 5
	fmt.Printf("Assignment: x is now %v\n", x)

	// 5. Pointer Operators
	// '&' gets the memory address of a variable.
	// '*' (when used before a type) declares a pointer.
	// '*' (when used before a variable) dereferences the pointer (gets the value).
	var ptr *int = &x
	fmt.Printf("Address-of: address of x is %v\n", &x)
	fmt.Printf("Pointer: ptr stores address %v\n", ptr)
	fmt.Printf("Dereference: value at ptr is %v\n", *ptr)

	// 6. Selector Operator ('.')
	// The dot operator is used to access package members (like fmt.Println)
	// or fields/methods of a struct.
	type Sample struct {
		id int
	}
	s := Sample{id: 101}
	// Accessing the 'id' field using the selector operator.
	fmt.Printf("Selector: sample ID is %v\n", s.id)
}
