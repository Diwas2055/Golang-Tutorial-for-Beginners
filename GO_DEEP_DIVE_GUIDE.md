# Go Programming Language - Complete Deep Dive Guide

## Table of Contents
1. [Introduction & Setup](#introduction--setup)
2. [Core Fundamentals](#core-fundamentals)
3. [Data Structures](#data-structures)
4. [Control Flow](#control-flow)
5. [Functions & Methods](#functions--methods)
6. [Advanced Concepts](#advanced-concepts)
7. [Concurrency](#concurrency)
8. [Best Practices](#best-practices)

---

## Introduction & Setup

### What is Go?

Go (Golang) is a statically typed, compiled programming language designed at Google in 2007 and open-sourced in 2009. It combines the performance of compiled languages like C++ with the simplicity of dynamically typed languages like Python.

**Key Characteristics:**
- **Fast Compilation**: Compiles to native machine code
- **Concurrent by Design**: Built-in support for concurrent programming
- **Garbage Collected**: Automatic memory management
- **Simple Syntax**: Easy to learn and read
- **Cross-Platform**: Write once, compile for multiple platforms

**Real-World Usage:**
- Docker (containerization platform)
- Kubernetes (container orchestration)
- Terraform (infrastructure as code)
- Prometheus (monitoring system)
- Ethereum (blockchain)

### Installation & Setup

#### Linux Installation
```bash
# Using apt (Ubuntu/Debian)
sudo apt update
sudo apt install golang-go

# Verify installation
go version

# Check Go paths
go env GOROOT  # Go installation directory
go env GOPATH  # Go workspace directory
```

#### macOS Installation
```bash
# Using Homebrew
brew install go

# Or download from https://golang.org/dl/
```

#### Windows Installation
Download the installer from [golang.org/dl](https://golang.org/dl/) and follow the installation wizard.

#### Editor Setup
**Recommended: Visual Studio Code**
- Install the official Go extension
- Provides IntelliSense, debugging, and formatting

---

## Core Fundamentals

### 1. Your First Go Program

Every Go program starts with a package declaration and a main function.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

**Key Components:**
- `package main`: Defines an executable program (not a library)
- `import "fmt"`: Imports the format package for I/O operations
- `func main()`: Entry point of the program

**Running Go Programs:**
```bash
# Run directly
go run main.go

# Build executable
go build main.go
./main

# Initialize a module
go mod init myapp
```

### 2. Variables & Constants

#### Variable Declaration

**Method 1: Full Declaration**
```go
var name string = "John"
var age int = 30
var isActive bool = true
```

**Method 2: Type Inference**
```go
var name = "John"  // Go infers string type
var age = 30       // Go infers int type
```

**Method 3: Short Declaration (Inside Functions Only)**
```go
name := "John"
age := 30
isActive := true
```

**Multiple Variables**
```go
var x, y int = 10, 20
firstName, lastName := "John", "Doe"
```

#### Constants

Constants are immutable values defined at compile time.

```go
const Pi = 3.14159
const MaxUsers = 100
const AppName string = "MyApp"

// Multiple constants
const (
    StatusOK = 200
    StatusNotFound = 404
    StatusError = 500
)
```

#### Zero Values

Variables declared without explicit initialization get zero values:
```go
var i int       // 0
var f float64   // 0.0
var b bool      // false
var s string    // "" (empty string)
var p *int      // nil
```

### 3. Data Types

#### Basic Types

```go
// Boolean
var isValid bool = true

// String
var message string = "Hello"

// Integer types
var age int = 30           // Platform dependent (32 or 64 bit)
var count int8 = 127       // -128 to 127
var population int16       // -32768 to 32767
var bigNumber int32        // -2^31 to 2^31-1
var hugeNumber int64       // -2^63 to 2^63-1

// Unsigned integers
var positive uint = 42     // Only positive numbers
var byte1 uint8 = 255      // 0 to 255 (alias: byte)
var port uint16 = 8080     // 0 to 65535

// Floating point
var price float32 = 19.99
var precise float64 = 3.141592653589793

// Complex numbers
var complex1 complex64 = 1 + 2i
var complex2 complex128 = 2 + 3i

// Byte and Rune
var b byte = 'A'           // alias for uint8
var r rune = '世'          // alias for int32, represents Unicode code point
```

#### Type Conversion

Go requires explicit type conversion:
```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

// String conversion
import "strconv"
str := strconv.Itoa(42)           // int to string
num, err := strconv.Atoi("42")    // string to int
```

### 4. Input/Output Operations

#### Formatted Output

```go
name := "Alice"
age := 25
score := 95.5

// Print (no newline)
fmt.Print("Hello")

// Println (with newline)
fmt.Println("Hello, World!")

// Printf (formatted)
fmt.Printf("Name: %s, Age: %d, Score: %.2f\n", name, age, score)

// Common format specifiers:
// %v  - default format
// %T  - type of value
// %t  - boolean
// %d  - decimal integer
// %f  - floating point
// %s  - string
// %p  - pointer address
```

#### User Input

```go
var name string
var age int

fmt.Print("Enter your name: ")
fmt.Scan(&name)  // & gets the memory address

fmt.Print("Enter your age: ")
fmt.Scan(&age)

// Scanln reads until newline
fmt.Scanln(&name)
```

### 5. Pointers

Pointers store memory addresses of variables.

```go
// Declare a variable
x := 42

// Declare a pointer
var ptr *int = &x  // & gets the address

fmt.Println(ptr)   // Prints memory address
fmt.Println(*ptr)  // Prints value (42) - dereferencing

// Modify value through pointer
*ptr = 100
fmt.Println(x)     // Prints 100

// Nil pointer
var p *int
if p == nil {
    fmt.Println("Pointer is nil")
}
```

**Why Use Pointers?**
- Pass large structs efficiently (avoid copying)
- Modify function arguments
- Share data between functions

**Example: Function with Pointer**
```go
func increment(n *int) {
    *n = *n + 1
}

func main() {
    count := 5
    increment(&count)
    fmt.Println(count)  // Prints 6
}
```

---

## Data Structures

### 1. Arrays

Arrays have fixed size and store elements of the same type.

```go
// Declaration
var numbers [5]int
numbers[0] = 10
numbers[1] = 20

// Declaration with initialization
fruits := [3]string{"Apple", "Banana", "Orange"}

// Let compiler count
colors := [...]string{"Red", "Green", "Blue"}

// Get length
fmt.Println(len(fruits))  // 3

// Iterate
for i := 0; i < len(numbers); i++ {
    fmt.Println(numbers[i])
}

// Multi-dimensional arrays
var matrix [3][3]int
matrix[0][0] = 1
```

**Limitations:**
- Fixed size (cannot grow or shrink)
- Size is part of the type: `[3]int` ≠ `[5]int`

### 2. Slices

Slices are dynamic, flexible views into arrays.

```go
// Create slice
var numbers []int  // nil slice

// Using make
scores := make([]int, 5)      // length 5, capacity 5
buffer := make([]int, 3, 10)  // length 3, capacity 10

// Slice literal
fruits := []string{"Apple", "Banana", "Orange"}

// Append elements
fruits = append(fruits, "Mango")
fruits = append(fruits, "Grape", "Kiwi")

// Length and capacity
fmt.Println(len(fruits))  // number of elements
fmt.Println(cap(fruits))  // capacity

// Slicing
slice1 := fruits[1:3]   // elements 1 to 2
slice2 := fruits[:2]    // elements 0 to 1
slice3 := fruits[2:]    // elements 2 to end
slice4 := fruits[:]     // all elements

// Copy slices
source := []int{1, 2, 3}
dest := make([]int, len(source))
copy(dest, source)
```

**Practical Example:**
```go
// Dynamic list of bookings
bookings := []string{}
bookings = append(bookings, "John Doe")
bookings = append(bookings, "Jane Smith")

for i, booking := range bookings {
    fmt.Printf("%d: %s\n", i, booking)
}
```

### 3. Maps

Maps store key-value pairs (like dictionaries or hash tables).

```go
// Declaration
var person map[string]string

// Initialize with make
person = make(map[string]string)

// Literal initialization
user := map[string]string{
    "name":  "John",
    "email": "john@example.com",
    "city":  "New York",
}

// Add/Update
user["age"] = "30"
user["name"] = "John Doe"  // Update

// Access
name := user["name"]

// Check if key exists
email, exists := user["email"]
if exists {
    fmt.Println("Email:", email)
}

// Delete
delete(user, "city")

// Length
fmt.Println(len(user))

// Iterate
for key, value := range user {
    fmt.Printf("%s: %s\n", key, value)
}

// Map of slices
bookings := make(map[string][]string)
bookings["Monday"] = []string{"John", "Jane"}
bookings["Tuesday"] = []string{"Bob"}
```

**Complex Example:**
```go
// Store user data
type UserData struct {
    firstName string
    tickets   int
}

users := make(map[string]UserData)
users["user1"] = UserData{firstName: "John", tickets: 2}
users["user2"] = UserData{firstName: "Jane", tickets: 3}

// Access
if user, found := users["user1"]; found {
    fmt.Printf("%s has %d tickets\n", user.firstName, user.tickets)
}
```

### 4. Structs

Structs are custom data types that group related fields.

```go
// Define struct
type Person struct {
    FirstName string
    LastName  string
    Age       int
    Email     string
}

// Create instance
var p1 Person
p1.FirstName = "John"
p1.LastName = "Doe"
p1.Age = 30

// Struct literal
p2 := Person{
    FirstName: "Jane",
    LastName:  "Smith",
    Age:       25,
    Email:     "jane@example.com",
}

// Short form (order matters)
p3 := Person{"Bob", "Johnson", 35, "bob@example.com"}

// Access fields
fmt.Println(p2.FirstName)
```

**Struct with Methods:**
```go
type Rectangle struct {
    Width  float64
    Height float64
}

// Method with receiver
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Method with pointer receiver (can modify)
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    fmt.Println(rect.Area())  // 50
    
    rect.Scale(2)
    fmt.Println(rect.Area())  // 200
}
```

**Embedded Structs (Composition):**
```go
type Address struct {
    Street string
    City   string
}

type Employee struct {
    Name    string
    Address // Embedded struct
}

func main() {
    emp := Employee{
        Name: "John",
        Address: Address{
            Street: "123 Main St",
            City:   "Boston",
        },
    }
    
    fmt.Println(emp.City)  // Direct access to embedded field
}
```

---

## Control Flow

### 1. If-Else Statements

```go
age := 18

// Basic if
if age >= 18 {
    fmt.Println("Adult")
}

// If-else
if age >= 18 {
    fmt.Println("Adult")
} else {
    fmt.Println("Minor")
}

// If-else if-else
if age < 13 {
    fmt.Println("Child")
} else if age < 18 {
    fmt.Println("Teenager")
} else {
    fmt.Println("Adult")
}

// If with initialization
if score := getScore(); score > 90 {
    fmt.Println("Excellent")
}

// Multiple conditions
isValid := true
hasPermission := true

if isValid && hasPermission {
    fmt.Println("Access granted")
}
```

**Validation Example:**
```go
func validateInput(name string, email string, tickets int) (bool, string) {
    isValidName := len(name) >= 2
    isValidEmail := strings.Contains(email, "@")
    isValidTickets := tickets > 0 && tickets <= 10
    
    if !isValidName {
        return false, "Name too short"
    }
    if !isValidEmail {
        return false, "Invalid email"
    }
    if !isValidTickets {
        return false, "Invalid ticket count"
    }
    
    return true, "Valid"
}
```

### 2. Switch Statements

```go
// Basic switch
day := "Monday"

switch day {
case "Monday":
    fmt.Println("Start of week")
case "Friday":
    fmt.Println("End of week")
case "Saturday", "Sunday":
    fmt.Println("Weekend")
default:
    fmt.Println("Midweek")
}

// Switch with expressions
score := 85

switch {
case score >= 90:
    fmt.Println("A")
case score >= 80:
    fmt.Println("B")
case score >= 70:
    fmt.Println("C")
default:
    fmt.Println("F")
}

// Switch with initialization
switch hour := time.Now().Hour(); {
case hour < 12:
    fmt.Println("Morning")
case hour < 18:
    fmt.Println("Afternoon")
default:
    fmt.Println("Evening")
}

// Type switch
func describe(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %s\n", v)
    case bool:
        fmt.Printf("Boolean: %t\n", v)
    default:
        fmt.Printf("Unknown type\n")
    }
}
```

### 3. Loops

Go has only one loop keyword: `for`

**Basic For Loop:**
```go
// Traditional for loop
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// While-style loop
count := 0
for count < 5 {
    fmt.Println(count)
    count++
}

// Infinite loop
for {
    fmt.Println("Forever")
    break  // Exit with break
}
```

**Range Loop:**
```go
// Iterate over slice
fruits := []string{"Apple", "Banana", "Orange"}
for index, fruit := range fruits {
    fmt.Printf("%d: %s\n", index, fruit)
}

// Ignore index
for _, fruit := range fruits {
    fmt.Println(fruit)
}

// Only index
for i := range fruits {
    fmt.Println(i)
}

// Iterate over map
user := map[string]string{"name": "John", "email": "john@example.com"}
for key, value := range user {
    fmt.Printf("%s: %s\n", key, value)
}

// Iterate over string (runes)
for index, char := range "Hello" {
    fmt.Printf("%d: %c\n", index, char)
}
```

**Loop Control:**
```go
// Break - exit loop
for i := 0; i < 10; i++ {
    if i == 5 {
        break
    }
    fmt.Println(i)
}

// Continue - skip iteration
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue
    }
    fmt.Println(i)  // Only odd numbers
}

// Nested loops with labels
outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if i == 1 && j == 1 {
            break outer  // Break outer loop
        }
        fmt.Printf("(%d,%d) ", i, j)
    }
}
```

---

## Functions & Methods

### 1. Function Basics

```go
// Simple function
func greet() {
    fmt.Println("Hello!")
}

// Function with parameters
func greetPerson(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

// Function with return value
func add(a int, b int) int {
    return a + b
}

// Shortened parameter syntax
func multiply(a, b int) int {
    return a * b
}

// Multiple return values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// Named return values
func calculate(a, b int) (sum int, product int) {
    sum = a + b
    product = a * b
    return  // Naked return
}

// Using functions
func main() {
    greet()
    greetPerson("Alice")
    
    result := add(5, 3)
    fmt.Println(result)
    
    quotient, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", quotient)
    }
    
    s, p := calculate(4, 5)
    fmt.Printf("Sum: %d, Product: %d\n", s, p)
}
```

### 2. Variadic Functions

Functions that accept variable number of arguments.

```go
// Variadic function
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2, 3))           // 6
    fmt.Println(sum(1, 2, 3, 4, 5))     // 15
    
    // Pass slice
    nums := []int{10, 20, 30}
    fmt.Println(sum(nums...))           // 60
}
```

**Mixed Parameters:**
```go
func printInfo(prefix string, values ...int) {
    fmt.Print(prefix, ": ")
    for _, v := range values {
        fmt.Print(v, " ")
    }
    fmt.Println()
}

func main() {
    printInfo("Numbers", 1, 2, 3, 4, 5)
}
```

### 3. Anonymous Functions & Closures

```go
// Anonymous function
func main() {
    // Assign to variable
    add := func(a, b int) int {
        return a + b
    }
    fmt.Println(add(3, 4))
    
    // Immediate execution
    func(msg string) {
        fmt.Println(msg)
    }("Hello from anonymous function")
}

// Closure - function that captures variables
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    increment := counter()
    fmt.Println(increment())  // 1
    fmt.Println(increment())  // 2
    fmt.Println(increment())  // 3
}
```

### 4. Methods

Methods are functions with a receiver.

```go
type Circle struct {
    Radius float64
}

// Value receiver (read-only)
func (c Circle) Area() float64 {
    return 3.14159 * c.Radius * c.Radius
}

// Pointer receiver (can modify)
func (c *Circle) Scale(factor float64) {
    c.Radius *= factor
}

func main() {
    circle := Circle{Radius: 5}
    fmt.Println(circle.Area())  // 78.53975
    
    circle.Scale(2)
    fmt.Println(circle.Radius)  // 10
}
```

**When to use pointer receivers:**
- Need to modify the receiver
- Struct is large (avoid copying)
- Consistency (if some methods use pointers, all should)

---

## Advanced Concepts

### 1. Packages & Modules

**Package Structure:**
```
myapp/
├── go.mod
├── main.go
└── helper/
    └── helper.go
```

**helper/helper.go:**
```go
package helper

import "strings"

// Exported function (starts with capital letter)
func ValidateEmail(email string) bool {
    return strings.Contains(email, "@")
}

// Unexported function (starts with lowercase)
func internalHelper() {
    // Only accessible within helper package
}
```

**main.go:**
```go
package main

import (
    "fmt"
    "myapp/helper"
)

func main() {
    email := "test@example.com"
    if helper.ValidateEmail(email) {
        fmt.Println("Valid email")
    }
}
```

**Initialize Module:**
```bash
go mod init myapp
go run main.go

# Or run all files
go run *.go
```

**Scope Rules:**
- **Local Scope**: Variables inside functions
- **Package Scope**: Variables outside functions (accessible in same package)
- **Exported**: Capitalized names (accessible from other packages)

### 2. Interfaces

Interfaces define behavior (method sets).

```go
// Define interface
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Rectangle implements Shape
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// Circle implements Shape
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * 3.14159 * c.Radius
}

// Function that accepts any Shape
func printShapeInfo(s Shape) {
    fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    circle := Circle{Radius: 7}
    
    printShapeInfo(rect)
    printShapeInfo(circle)
}
```

**Empty Interface:**
```go
// interface{} accepts any type
func describe(i interface{}) {
    fmt.Printf("Type: %T, Value: %v\n", i, i)
}

func main() {
    describe(42)
    describe("hello")
    describe(true)
    describe([]int{1, 2, 3})
}
```

### 3. Error Handling

Go uses explicit error handling (no exceptions).

```go
import (
    "errors"
    "fmt"
)

// Return error
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

// Custom error
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func validateAge(age int) error {
    if age < 0 {
        return &ValidationError{
            Field:   "age",
            Message: "must be positive",
        }
    }
    return nil
}

func main() {
    // Handle error
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result)
    
    // Check error
    if err := validateAge(-5); err != nil {
        fmt.Println("Validation error:", err)
    }
}
```

**Panic and Recover:**
```go
func riskyOperation() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from:", r)
        }
    }()
    
    panic("something went wrong")
}

func main() {
    riskyOperation()
    fmt.Println("Program continues")
}
```

### 4. Defer Statement

Defer postpones function execution until surrounding function returns.

```go
func main() {
    defer fmt.Println("World")
    fmt.Println("Hello")
    // Output: Hello World
}

// Multiple defers (LIFO - Last In First Out)
func example() {
    defer fmt.Println("1")
    defer fmt.Println("2")
    defer fmt.Println("3")
    // Output: 3 2 1
}

// Practical use: cleanup
func readFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()  // Ensures file is closed
    
    // Read file...
    return nil
}
```

---

## Concurrency

Go's killer feature: built-in concurrency with goroutines and channels.

### 1. Goroutines

Lightweight threads managed by Go runtime.

```go
import (
    "fmt"
    "time"
)

func sayHello() {
    fmt.Println("Hello from goroutine")
}

func main() {
    // Start goroutine
    go sayHello()
    
    // Main continues immediately
    fmt.Println("Hello from main")
    
    time.Sleep(100 * time.Millisecond)  // Wait for goroutine
}

// Multiple goroutines
func printNumbers() {
    for i := 1; i <= 5; i++ {
        fmt.Printf("Number: %d\n", i)
        time.Sleep(100 * time.Millisecond)
    }
}

func printLetters() {
    for i := 'a'; i <= 'e'; i++ {
        fmt.Printf("Letter: %c\n", i)
        time.Sleep(150 * time.Millisecond)
    }
}

func main() {
    go printNumbers()
    go printLetters()
    
    time.Sleep(1 * time.Second)
}
```

### 2. WaitGroups

Proper way to wait for goroutines.

```go
import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()  // Decrement counter when done
    
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup
    
    for i := 1; i <= 5; i++ {
        wg.Add(1)  // Increment counter
        go worker(i, &wg)
    }
    
    wg.Wait()  // Block until counter is 0
    fmt.Println("All workers done")
}
```

### 3. Channels

Channels allow goroutines to communicate safely.

```go
// Create channel
ch := make(chan int)

// Send to channel
ch <- 42

// Receive from channel
value := <-ch

// Buffered channel
buffered := make(chan int, 3)
buffered <- 1
buffered <- 2
buffered <- 3
```

**Example: Producer-Consumer**
```go
func producer(ch chan int) {
    for i := 1; i <= 5; i++ {
        fmt.Printf("Producing: %d\n", i)
        ch <- i
        time.Sleep(500 * time.Millisecond)
    }
    close(ch)  // Close channel when done
}

func consumer(ch chan int) {
    for value := range ch {  // Receive until channel closed
        fmt.Printf("Consuming: %d\n", value)
    }
}

func main() {
    ch := make(chan int)
    
    go producer(ch)
    consumer(ch)
}
```

**Buffered Channels:**
```go
func main() {
    ch := make(chan string, 2)  // Buffer size 2
    
    ch <- "first"   // Doesn't block
    ch <- "second"  // Doesn't block
    // ch <- "third"  // Would block (buffer full)
    
    fmt.Println(<-ch)  // first
    fmt.Println(<-ch)  // second
}
```

### 4. Select Statement

Handle multiple channel operations.

```go
func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "from channel 1"
    }()
    
    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "from channel 2"
    }()
    
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println(msg1)
        case msg2 := <-ch2:
            fmt.Println(msg2)
        }
    }
}
```

**Select with Timeout:**
```go
func main() {
    ch := make(chan string)
    
    go func() {
        time.Sleep(2 * time.Second)
        ch <- "result"
    }()
    
    select {
    case result := <-ch:
        fmt.Println(result)
    case <-time.After(1 * time.Second):
        fmt.Println("timeout")
    }
}
```

**Select with Default:**
```go
func main() {
    ch := make(chan int, 1)
    
    select {
    case ch <- 42:
        fmt.Println("Sent")
    default:
        fmt.Println("Channel full")
    }
}
```

### 5. Practical Concurrency Example

Ticket booking with goroutines:

```go
import (
    "fmt"
    "sync"
    "time"
)

type Booking struct {
    firstName string
    lastName  string
    email     string
    tickets   uint
}

var wg sync.WaitGroup

func sendTicket(booking Booking) {
    defer wg.Done()
    
    time.Sleep(10 * time.Second)  // Simulate email sending
    
    ticket := fmt.Sprintf("%d tickets for %s %s", 
        booking.tickets, booking.firstName, booking.lastName)
    
    fmt.Println("#####################")
    fmt.Printf("Sending ticket:\n%s\nto email: %s\n", ticket, booking.email)
    fmt.Println("#####################")
}

func main() {
    bookings := []Booking{
        {"John", "Doe", "john@example.com", 2},
        {"Jane", "Smith", "jane@example.com", 3},
        {"Bob", "Johnson", "bob@example.com", 1},
    }
    
    for _, booking := range bookings {
        wg.Add(1)
        go sendTicket(booking)
    }
    
    wg.Wait()
    fmt.Println("All tickets sent")
}
```

---

## Best Practices

### 1. Code Organization

```
project/
├── go.mod
├── go.sum
├── main.go
├── cmd/
│   └── app/
│       └── main.go
├── internal/
│   ├── handlers/
│   ├── models/
│   └── services/
├── pkg/
│   └── utils/
├── configs/
├── tests/
└── README.md
```

**Guidelines:**
- `cmd/`: Application entry points
- `internal/`: Private application code
- `pkg/`: Public library code
- Keep packages focused and cohesive
- Use meaningful package names

### 2. Naming Conventions

```go
// Good
var userCount int
var isActive bool
func getUserByID(id int) User
type UserService struct{}

// Bad
var user_count int
var active bool
func get_user_by_id(id int) User
type User_Service struct{}
```

**Rules:**
- Use camelCase for variables and functions
- Use PascalCase for exported names
- Short names for short scopes (`i` for loop index)
- Descriptive names for package scope
- Acronyms should be all caps: `userID`, `HTTPServer`

### 3. Error Handling

```go
// Good - handle errors immediately
result, err := doSomething()
if err != nil {
    return fmt.Errorf("failed to do something: %w", err)
}

// Bad - ignoring errors
result, _ := doSomething()

// Good - wrap errors with context
if err := validateInput(data); err != nil {
    return fmt.Errorf("validation failed: %w", err)
}
```

### 4. Comments

```go
// Package comment describes the package
package calculator

// Add returns the sum of two integers.
// It handles overflow by returning an error.
func Add(a, b int) (int, error) {
    // Check for overflow
    if a > 0 && b > math.MaxInt-a {
        return 0, errors.New("integer overflow")
    }
    return a + b, nil
}
```

### 5. Testing

```go
// calculator.go
package calculator

func Add(a, b int) int {
    return a + b
}

// calculator_test.go
package calculator

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}

// Table-driven tests
func TestAddTable(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive", 2, 3, 5},
        {"negative", -2, -3, -5},
        {"zero", 0, 5, 5},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("got %d, want %d", result, tt.expected)
            }
        })
    }
}
```

**Run tests:**
```bash
go test
go test -v          # Verbose
go test -cover      # Coverage
go test ./...       # All packages
```

### 6. Performance Tips

```go
// Use pointers for large structs
func ProcessLargeStruct(data *LargeStruct) {
    // Avoids copying
}

// Preallocate slices when size is known
items := make([]Item, 0, 100)

// Use string builder for concatenation
var builder strings.Builder
for i := 0; i < 100; i++ {
    builder.WriteString("item")
}
result := builder.String()

// Avoid unnecessary allocations
// Bad
for i := 0; i < n; i++ {
    temp := make([]int, size)
    // use temp
}

// Good
temp := make([]int, size)
for i := 0; i < n; i++ {
    // reuse temp
}
```

### 7. Common Patterns

**Singleton:**
```go
import "sync"

type Database struct {
    // fields
}

var (
    instance *Database
    once     sync.Once
)

func GetInstance() *Database {
    once.Do(func() {
        instance = &Database{}
    })
    return instance
}
```

**Options Pattern:**
```go
type Server struct {
    host string
    port int
}

type Option func(*Server)

func WithHost(host string) Option {
    return func(s *Server) {
        s.host = host
    }
}

func WithPort(port int) Option {
    return func(s *Server) {
        s.port = port
    }
}

func NewServer(opts ...Option) *Server {
    s := &Server{
        host: "localhost",
        port: 8080,
    }
    
    for _, opt := range opts {
        opt(s)
    }
    
    return s
}

// Usage
server := NewServer(
    WithHost("0.0.0.0"),
    WithPort(3000),
)
```

---

## Quick Reference

### Common Commands

```bash
# Initialize module
go mod init module-name

# Run program
go run main.go
go run *.go

# Build executable
go build
go build -o myapp

# Install dependencies
go get package-name
go mod tidy

# Format code
go fmt ./...
gofmt -w .

# Run tests
go test
go test -v
go test -cover

# Vet code
go vet ./...

# Download dependencies
go mod download

# View documentation
go doc package-name
```

### Useful Packages

```go
// Standard library
import (
    "fmt"           // Formatting
    "strings"       // String manipulation
    "strconv"       // String conversion
    "time"          // Time operations
    "os"            // OS operations
    "io"            // I/O operations
    "encoding/json" // JSON encoding/decoding
    "net/http"      // HTTP client/server
    "sync"          // Synchronization primitives
    "context"       // Context for cancellation
)
```

### Memory Management

```go
// new() - allocates zeroed memory, returns pointer
p := new(int)
*p = 42

// make() - initializes slices, maps, channels
slice := make([]int, 10)
m := make(map[string]int)
ch := make(chan int)
```

---

## Complete Example: Booking Application

```go
package main

import (
    "fmt"
    "strings"
    "sync"
    "time"
)

const conferenceTickets = 50

var (
    conferenceName   = "Go Conference"
    remainingTickets uint = 50
    bookings         = make([]UserData, 0)
    wg               = sync.WaitGroup{}
)

type UserData struct {
    firstName string
    lastName  string
    email     string
    tickets   uint
}

func main() {
    greetUsers()
    
    for remainingTickets > 0 {
        firstName, lastName, email, userTickets := getUserInput()
        
        isValidName, isValidEmail, isValidTickets := validateInput(
            firstName, lastName, email, userTickets,
        )
        
        if !isValidName || !isValidEmail || !isValidTickets {
            if !isValidName {
                fmt.Println("First or last name is too short")
            }
            if !isValidEmail {
                fmt.Println("Email doesn't contain @ sign")
            }
            if !isValidTickets {
                fmt.Println("Invalid number of tickets")
            }
            continue
        }
        
        bookTicket(firstName, lastName, email, userTickets)
        
        wg.Add(1)
        go sendTicket(UserData{firstName, lastName, email, userTickets})
        
        fmt.Printf("First names of bookings: %v\n", getFirstNames())
    }
    
    wg.Wait()
}

func greetUsers() {
    fmt.Printf("Welcome to %v booking application\n", conferenceName)
    fmt.Printf("We have %v tickets and %v are available\n",
        conferenceTickets, remainingTickets)
}

func getUserInput() (string, string, string, uint) {
    var firstName, lastName, email string
    var userTickets uint
    
    fmt.Print("Enter your first name: ")
    fmt.Scan(&firstName)
    
    fmt.Print("Enter your last name: ")
    fmt.Scan(&lastName)
    
    fmt.Print("Enter your email: ")
    fmt.Scan(&email)
    
    fmt.Print("Enter number of tickets: ")
    fmt.Scan(&userTickets)
    
    return firstName, lastName, email, userTickets
}

func validateInput(firstName, lastName, email string, tickets uint) (bool, bool, bool) {
    isValidName := len(firstName) >= 2 && len(lastName) >= 2
    isValidEmail := strings.Contains(email, "@")
    isValidTickets := tickets > 0 && tickets <= remainingTickets
    
    return isValidName, isValidEmail, isValidTickets
}

func bookTicket(firstName, lastName, email string, tickets uint) {
    remainingTickets -= tickets
    
    user := UserData{
        firstName: firstName,
        lastName:  lastName,
        email:     email,
        tickets:   tickets,
    }
    
    bookings = append(bookings, user)
    
    fmt.Printf("Thank you %v %v for booking %v tickets\n",
        firstName, lastName, tickets)
    fmt.Printf("%v tickets remaining\n", remainingTickets)
}

func getFirstNames() []string {
    firstNames := []string{}
    for _, booking := range bookings {
        firstNames = append(firstNames, booking.firstName)
    }
    return firstNames
}

func sendTicket(user UserData) {
    defer wg.Done()
    
    time.Sleep(10 * time.Second)
    
    ticket := fmt.Sprintf("%v tickets for %v %v",
        user.tickets, user.firstName, user.lastName)
    
    fmt.Println("\n#####################")
    fmt.Printf("Sending ticket:\n%v\nto: %v\n", ticket, user.email)
    fmt.Println("#####################\n")
}
```

---

## Next Steps

1. **Build Projects**: Create real applications
2. **Read Standard Library**: Explore `pkg.go.dev`
3. **Learn Web Development**: `net/http`, Gin, Echo
4. **Database Integration**: `database/sql`, GORM
5. **Testing**: Write comprehensive tests
6. **Concurrency Patterns**: Master goroutines and channels
7. **Performance**: Learn profiling with `pprof`
8. **Deployment**: Docker, Kubernetes

## Resources

- Official Documentation: [golang.org/doc](https://golang.org/doc)
- Go by Example: [gobyexample.com](https://gobyexample.com)
- Effective Go: [golang.org/doc/effective_go](https://golang.org/doc/effective_go)
- Go Playground: [play.golang.org](https://play.golang.org)
- Awesome Go: [awesome-go.com](https://awesome-go.com)

---

**Happy Coding! 🚀**
