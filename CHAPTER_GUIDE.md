# Chapter-by-Chapter Learning Guide

This guide provides a structured learning path through the Go tutorial with clear objectives and exercises for each chapter.

---

## Chapter 01: Hello World (`chapters/01-basic.go`)

**Objective**: Understand the basic structure of a Go program

**What You'll Learn**:
- Package declaration
- Import statements
- The `main` function
- Printing to console

**Key Concepts**:
```go
package main        // Declares this as an executable program
import "fmt"        // Imports the format package
func main() { }     // Entry point of the program
```

**Exercise**:
1. Modify the program to print your name
2. Add multiple print statements
3. Try using `fmt.Println()` vs `fmt.Print()`

**Run**:
```bash
go run chapters/01-basic.go
```

---

## Chapter 02: Variables (`chapters/02-variables.go`)

**Objective**: Learn variable declaration and user input

**What You'll Learn**:
- Variable declaration with `var` and `:=`
- Constants with `const`
- Different data types (string, int, uint)
- Getting user input with `fmt.Scan()`
- String formatting with `fmt.Printf()`

**Key Concepts**:
```go
const conferenceTickets int = 50    // Constant
var remainingTickets uint = 50      // Variable with type
conferenceName := "Go Conference"   // Short declaration

fmt.Scanln(&firstName)              // Get user input
fmt.Printf("Welcome %v\n", name)    // Formatted output
```

**Exercise**:
1. Add a new field for phone number
2. Change the total number of tickets
3. Add validation to check if name is empty

**Run**:
```bash
go run chapters/02-variables.go
```

---

## Chapter 03: Arrays & Slices (`chapters/03-arrays.go`)

**Objective**: Store multiple bookings in a dynamic list

**What You'll Learn**:
- Difference between arrays and slices
- Creating slices with `[]type{}`
- Appending to slices with `append()`
- Slice operations

**Key Concepts**:
```go
bookings := []string{}                      // Empty slice
bookings = append(bookings, "John Doe")     // Add element
fmt.Printf("Bookings: %v\n", bookings)      // Print slice
```

**Exercise**:
1. Create a slice to store email addresses
2. Print the number of bookings
3. Try creating a fixed-size array and see the difference

**Run**:
```bash
go run chapters/03-arrays.go
```

---

## Chapter 04: Loops (`chapters/04-loops.go`)

**Objective**: Allow multiple bookings with loops

**What You'll Learn**:
- `for` loop syntax
- Infinite loops with `for { }`
- Range loops with `for _, item := range slice`
- String manipulation with `strings.Fields()`

**Key Concepts**:
```go
for {                                    // Infinite loop
    // Get user input and book tickets
}

for _, booking := range bookings {       // Iterate over slice
    names := strings.Fields(booking)     // Split string
}
```

**Exercise**:
1. Add a condition to break the loop when tickets run out
2. Print booking number (1, 2, 3...) for each booking
3. Extract and print only last names

**Run**:
```bash
go run chapters/04-loops.go
```

---

## Chapter 05: Conditionals (`chapters/05-if-else.go`)

**Objective**: Validate user input before booking

**What You'll Learn**:
- `if-else` statements
- Boolean operators (`&&`, `||`, `!`)
- Input validation
- `break` and `continue` statements
- `switch` statements

**Key Concepts**:
```go
isValidName := len(firstName) >= 2 && len(lastName) >= 2
isValidEmail := strings.Contains(email, "@")

if isValidName && isValidEmail {
    // Book ticket
} else {
    // Show error
    continue  // Skip to next iteration
}

switch city {
case "London", "Berlin":
    // Handle
default:
    // Default case
}
```

**Exercise**:
1. Add validation for minimum ticket purchase (at least 1)
2. Add validation for maximum tickets per person (max 5)
3. Create a switch statement for different ticket types

**Run**:
```bash
go run chapters/05-if-else.go
```

---

## Chapter 06: Functions (`chapters/06-functions.go`)

**Objective**: Organize code into reusable functions

**What You'll Learn**:
- Function declaration
- Parameters and return values
- Multiple return values
- Package-level variables

**Key Concepts**:
```go
func greetUsers() {
    fmt.Println("Welcome")
}

func getUserInput() (string, string, string, uint) {
    // Get input
    return firstName, lastName, email, userTickets
}

func validateUserInput(name string, email string) (bool, bool) {
    isValidName := len(name) >= 2
    isValidEmail := strings.Contains(email, "@")
    return isValidName, isValidEmail
}
```

**Exercise**:
1. Create a function to calculate total revenue
2. Create a function to find a booking by email
3. Add a function to cancel a booking

**Run**:
```bash
go run chapters/06-functions.go
```

---

## Chapter 07: Maps (`chapters/07-maps.go`)

**Objective**: Store structured booking data with maps

**What You'll Learn**:
- Creating maps with `make(map[keyType]valueType)`
- Adding and accessing map elements
- Map of maps for complex data
- Converting types with `strconv`

**Key Concepts**:
```go
bookings := make([]map[string]string, 0)

user := make(map[string]string)
user["firstName"] = firstName
user["email"] = email

bookings = append(bookings, user)

// Access
name := bookings[0]["firstName"]
```

**Exercise**:
1. Add a booking ID to each booking
2. Create a function to search bookings by email
3. Store booking timestamp

**Run**:
```bash
go run chapters/07-maps.go
```

---

## Chapter 08: Structs (`chapters/08-structs.go`)

**Objective**: Use custom types for better code organization

**What You'll Learn**:
- Defining structs
- Creating struct instances
- Accessing struct fields
- Slices of structs

**Key Concepts**:
```go
type User struct {
    firstName       string
    lastName        string
    email           string
    numberOfTickets uint
}

var user = User{
    firstName:       "John",
    lastName:        "Doe",
    email:           "john@example.com",
    numberOfTickets: 2,
}

bookings := make([]User, 0)
bookings = append(bookings, user)

// Access
fmt.Println(bookings[0].firstName)
```

**Exercise**:
1. Add a `bookingID` field to the struct
2. Add a `bookingDate` field (use `time.Time`)
3. Create a method to print user details

**Run**:
```bash
go run chapters/08-structs.go
```

---

## Chapter 09: Packages (`chapters/packages/`)

**Objective**: Organize code across multiple files and packages

**What You'll Learn**:
- Creating custom packages
- Exporting functions (capitalized names)
- Importing local packages
- Module initialization with `go mod init`

**Structure**:
```
packages/
├── main.go
├── go-mod.txt
└── helper/
    └── helper.go
```

**Key Concepts**:
```go
// helper/helper.go
package helper

func ValidateUserInput(...) (bool, bool, bool) {
    // Exported function (starts with capital letter)
}

// main.go
package main

import "booking-app/helper"

func main() {
    isValid, _, _ := helper.ValidateUserInput(...)
}
```

**Exercise**:
1. Create a new package for email operations
2. Move the booking logic to a separate package
3. Create a config package for constants

**Run**:
```bash
cd chapters/packages
go mod init booking-app
go run main.go
```

---

## Chapter 10: Concurrency (`main.go`)

**Objective**: Send tickets concurrently using goroutines

**What You'll Learn**:
- Goroutines with `go` keyword
- WaitGroups for synchronization
- Concurrent execution
- The `defer` statement

**Key Concepts**:
```go
var wg = sync.WaitGroup{}

func main() {
    wg.Add(1)                    // Increment counter
    go sendTicket(userData)      // Start goroutine
    wg.Wait()                    // Wait for all goroutines
}

func sendTicket(user UserData) {
    defer wg.Done()              // Decrement when done
    time.Sleep(10 * time.Second) // Simulate work
    // Send ticket
}
```

**Exercise**:
1. Add multiple goroutines for different tasks
2. Use channels to communicate between goroutines
3. Implement a timeout for ticket sending

**Run**:
```bash
go run main.go helper.go
```

---

## Progressive Learning Path

### Week 1: Fundamentals
- Day 1-2: Chapters 01-02 (Basics, Variables)
- Day 3-4: Chapters 03-04 (Slices, Loops)
- Day 5-7: Chapter 05 (Conditionals, Validation)

### Week 2: Functions & Data Structures
- Day 1-2: Chapter 06 (Functions)
- Day 3-4: Chapter 07 (Maps)
- Day 5-7: Chapter 08 (Structs)

### Week 3: Advanced Topics
- Day 1-3: Chapter 09 (Packages)
- Day 4-7: Chapter 10 (Concurrency)

### Week 4: Practice & Projects
- Build your own projects
- Explore standard library
- Read `GO_DEEP_DIVE_GUIDE.md`

---

## Common Challenges & Solutions

### Challenge 1: "undefined: variable"
**Solution**: Make sure variables are declared before use. Use `:=` for new variables, `=` for existing ones.

### Challenge 2: "cannot use variable (type X) as type Y"
**Solution**: Go requires explicit type conversion. Use `int(variable)` or similar.

### Challenge 3: "all goroutines are asleep - deadlock"
**Solution**: Make sure to call `wg.Done()` for every `wg.Add(1)`.

### Challenge 4: "package not found"
**Solution**: Run `go mod init module-name` and ensure import paths match module name.

---

## Next Steps After Completing All Chapters

1. **Build Projects**:
   - Todo list application
   - REST API server
   - CLI tool
   - Web scraper

2. **Learn Advanced Topics**:
   - Interfaces and polymorphism
   - Error handling patterns
   - Testing with `testing` package
   - Database integration

3. **Explore Frameworks**:
   - Gin (web framework)
   - GORM (ORM)
   - Cobra (CLI)
   - Echo (web framework)

4. **Read Documentation**:
   - Effective Go
   - Go standard library
   - Go blog

---

**Remember**: Practice is key! Try to modify each example and experiment with the code. Don't just read—write code! 🚀
