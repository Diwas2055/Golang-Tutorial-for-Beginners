// Chapter 02: Operators & Symbols
// This chapter demonstrates Go's operators and special symbols:
// 1. Math & Comparison
// 2. Short Declaration (:=)
// 3. Pointers (&, *)
// 4. Access (.) and Ignore (_)
// 5. Variadic (...) and Slices ([])
package main

import "fmt"

func main() {
	// 1. Math & Comparison
	a, b := 10, 3
	fmt.Printf("Add: %v, Mod: %v, Equal: %v\n", a+b, a%b, a == b)

	// 2. Short Declaration (:=)
	// Combines declaration and assignment into one step.
	x := 25

	// 3. Pointers (&, *)
	// '&' gets the memory address. '*' defines a pointer type or dereferences.
	var ptr *int = &x
	fmt.Printf("Address of x: %v, Value via ptr: %v\n", &x, *ptr)

	// 4. Blank Identifier (_)
	// Used to ignore a return value that you don't need.
	val, _ := dummyFunction()
	fmt.Printf("Value is %v, but we ignored the second return.\n", val)

	// 5. Variadic (...)
	// Functions can take any number of arguments.
	printAll("Colors:", "Red", "Green", "Blue")

	// 6. Selector/Access (.)
	// Used to access fields in a struct or functions in a package.
	type User struct{ name string }
	u := User{name: "Alice"}
	fmt.Printf("Access name via dot: %s\n", u.name)
}

func dummyFunction() (int, int) {
	return 100, 200
}

func printAll(prefix string, items ...string) {
	// 7. Index/Slice ([])
	// Use [] to access elements or define slice types.
	fmt.Printf("%s First item is %s\n", prefix, items[0])
}
