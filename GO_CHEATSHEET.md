# Go Programming Cheat Sheet

Quick reference for Go syntax and common patterns.

---

## Basic Syntax

### Hello World
```go
package main
import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### Comments
```go
// Single line comment

/*
Multi-line
comment
*/
```

---

## Variables & Constants

### Declaration
```go
// With type
var name string = "John"
var age int = 30

// Type inference
var name = "John"
var age = 30

// Short declaration (inside functions only)
name := "John"
age := 30

// Multiple variables
var x, y int = 10, 20
a, b := "hello", "world"

// Constants
const Pi = 3.14159
const MaxSize int = 100
```

### Zero Values
```go
var i int       // 0
var f float64   // 0.0
var b bool      // false
var s string    // ""
var p *int      // nil
```

---

## Data Types

### Basic Types
```go
bool                    // true, false
string                  // "text"
int, int8, int16, int32, int64
uint, uint8, uint16, uint32, uint64
byte                    // alias for uint8
rune                    // alias for int32 (Unicode)
float32, float64
complex64, complex128
```

### Type Conversion
```go
i := 42
f := float64(i)
u := uint(f)

// String conversion
import "strconv"
s := strconv.Itoa(42)           // int to string
n, err := strconv.Atoi("42")    // string to int
```

---

## Operators

### Arithmetic
```go
+    // Addition
-    // Subtraction
*    // Multiplication
/    // Division
%    // Modulus
++   // Increment
--   // Decrement
```

### Comparison
```go
==   // Equal
!=   // Not equal
<    // Less than
>    // Greater than
<=   // Less than or equal
>=   // Greater than or equal
```

### Logical
```go
&&   // AND
||   // OR
!    // NOT
```

---

## Control Flow

### If-Else
```go
if x > 0 {
    fmt.Println("Positive")
} else if x < 0 {
    fmt.Println("Negative")
} else {
    fmt.Println("Zero")
}

// If with initialization
if err := doSomething(); err != nil {
    fmt.Println("Error:", err)
}
```

### Switch
```go
switch day {
case "Monday":
    fmt.Println("Start of week")
case "Friday":
    fmt.Println("End of week")
default:
    fmt.Println("Midweek")
}

// Switch without expression
switch {
case x < 0:
    fmt.Println("Negative")
case x > 0:
    fmt.Println("Positive")
default:
    fmt.Println("Zero")
}
```

### For Loop
```go
// Traditional for
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// While-style
count := 0
for count < 10 {
    count++
}

// Infinite loop
for {
    // break to exit
}

// Range over slice
for index, value := range slice {
    fmt.Println(index, value)
}

// Range over map
for key, value := range myMap {
    fmt.Println(key, value)
}

// Ignore index/key
for _, value := range slice {
    fmt.Println(value)
}
```

---

## Functions

### Basic Function
```go
func greet(name string) {
    fmt.Println("Hello,", name)
}
```

### With Return Value
```go
func add(a, b int) int {
    return a + b
}
```

### Multiple Return Values
```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Usage
result, err := divide(10, 2)
if err != nil {
    fmt.Println("Error:", err)
}
```

### Named Return Values
```go
func calculate(a, b int) (sum int, product int) {
    sum = a + b
    product = a * b
    return  // Naked return
}
```

### Variadic Functions
```go
func sum(numbers ...int) int {
    total := 0
    for _, n := range numbers {
        total += n
    }
    return total
}

// Usage
sum(1, 2, 3, 4, 5)
```

---

## Data Structures

### Arrays
```go
// Declaration
var arr [5]int
arr[0] = 10

// Initialization
arr := [3]int{1, 2, 3}
arr := [...]int{1, 2, 3}  // Compiler counts

// Length
len(arr)
```

### Slices
```go
// Declaration
var s []int

// Make
s := make([]int, 5)       // length 5
s := make([]int, 5, 10)   // length 5, capacity 10

// Literal
s := []int{1, 2, 3}

// Append
s = append(s, 4)
s = append(s, 5, 6, 7)

// Slicing
s[1:3]   // elements 1 to 2
s[:3]    // elements 0 to 2
s[2:]    // elements 2 to end
s[:]     // all elements

// Length and capacity
len(s)
cap(s)

// Copy
dest := make([]int, len(src))
copy(dest, src)
```

### Maps
```go
// Declaration
var m map[string]int

// Make
m := make(map[string]int)

// Literal
m := map[string]int{
    "one": 1,
    "two": 2,
}

// Add/Update
m["three"] = 3

// Access
value := m["one"]

// Check existence
value, exists := m["one"]
if exists {
    fmt.Println(value)
}

// Delete
delete(m, "one")

// Length
len(m)

// Iterate
for key, value := range m {
    fmt.Println(key, value)
}
```

### Structs
```go
// Define
type Person struct {
    Name string
    Age  int
}

// Create
var p Person
p.Name = "John"
p.Age = 30

// Literal
p := Person{Name: "John", Age: 30}
p := Person{"John", 30}  // Order matters

// Pointer to struct
p := &Person{Name: "John", Age: 30}
p.Age = 31  // Automatic dereferencing
```

---

## Pointers

```go
// Declare
var p *int

// Get address
x := 42
p = &x

// Dereference
fmt.Println(*p)  // 42

// Modify through pointer
*p = 100
fmt.Println(x)   // 100

// New
p := new(int)
*p = 42
```

---

## Methods

```go
type Rectangle struct {
    Width, Height float64
}

// Value receiver
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Pointer receiver
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

// Usage
rect := Rectangle{Width: 10, Height: 5}
area := rect.Area()
rect.Scale(2)
```

---

## Interfaces

```go
// Define
type Shape interface {
    Area() float64
}

// Implement (implicit)
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14159 * c.Radius * c.Radius
}

// Use
func printArea(s Shape) {
    fmt.Println(s.Area())
}

// Empty interface
var i interface{}
i = 42
i = "hello"
i = true

// Type assertion
s := i.(string)

// Type switch
switch v := i.(type) {
case int:
    fmt.Println("int:", v)
case string:
    fmt.Println("string:", v)
}
```

---

## Error Handling

```go
// Create error
import "errors"
err := errors.New("something went wrong")

// Format error
err := fmt.Errorf("failed: %w", originalErr)

// Check error
result, err := doSomething()
if err != nil {
    return err
}

// Custom error
type MyError struct {
    Msg string
}

func (e *MyError) Error() string {
    return e.Msg
}
```

---

## Concurrency

### Goroutines
```go
// Start goroutine
go myFunction()

// Anonymous goroutine
go func() {
    fmt.Println("Hello from goroutine")
}()
```

### Channels
```go
// Create
ch := make(chan int)
buffered := make(chan int, 10)

// Send
ch <- 42

// Receive
value := <-ch

// Close
close(ch)

// Range over channel
for value := range ch {
    fmt.Println(value)
}

// Check if closed
value, ok := <-ch
if !ok {
    fmt.Println("Channel closed")
}
```

### WaitGroup
```go
import "sync"

var wg sync.WaitGroup

func worker(id int) {
    defer wg.Done()
    fmt.Println("Worker", id)
}

func main() {
    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go worker(i)
    }
    wg.Wait()
}
```

### Select
```go
select {
case msg := <-ch1:
    fmt.Println(msg)
case msg := <-ch2:
    fmt.Println(msg)
case <-time.After(1 * time.Second):
    fmt.Println("timeout")
default:
    fmt.Println("no message")
}
```

### Mutex
```go
import "sync"

var (
    counter int
    mutex   sync.Mutex
)

func increment() {
    mutex.Lock()
    counter++
    mutex.Unlock()
}
```

---

## Packages

### Import
```go
import "fmt"
import "strings"

// Multiple imports
import (
    "fmt"
    "strings"
)

// Alias
import f "fmt"

// Import for side effects
import _ "image/png"
```

### Create Package
```go
// mypackage/mypackage.go
package mypackage

// Exported (public)
func PublicFunction() {}

// Unexported (private)
func privateFunction() {}
```

---

## Common Standard Library

### fmt
```go
fmt.Print("text")
fmt.Println("text")
fmt.Printf("Name: %s, Age: %d\n", name, age)
fmt.Sprintf("formatted %s", str)
fmt.Scan(&variable)
fmt.Scanln(&variable)
```

### strings
```go
strings.Contains(s, substr)
strings.HasPrefix(s, prefix)
strings.HasSuffix(s, suffix)
strings.Split(s, sep)
strings.Join(slice, sep)
strings.ToLower(s)
strings.ToUpper(s)
strings.TrimSpace(s)
strings.Replace(s, old, new, n)
```

### strconv
```go
strconv.Itoa(42)              // int to string
strconv.Atoi("42")            // string to int
strconv.ParseFloat("3.14", 64)
strconv.FormatFloat(3.14, 'f', 2, 64)
```

### time
```go
time.Now()
time.Sleep(1 * time.Second)
time.After(1 * time.Second)

// Format
t.Format("2006-01-02 15:04:05")

// Parse
time.Parse("2006-01-02", "2024-01-15")
```

### os
```go
os.Open(filename)
os.Create(filename)
os.Remove(filename)
os.Getenv("PATH")
os.Exit(1)
```

---

## Testing

```go
// mypackage_test.go
package mypackage

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}

// Table-driven test
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

---

## Common Commands

```bash
# Run
go run main.go

# Build
go build
go build -o myapp

# Test
go test
go test -v
go test -cover

# Format
go fmt ./...

# Vet
go vet ./...

# Module
go mod init module-name
go mod tidy
go get package-name

# Documentation
go doc package-name
```

---

## Format Specifiers

```go
%v   // Default format
%+v  // Struct with field names
%#v  // Go-syntax representation
%T   // Type
%t   // Boolean
%d   // Decimal integer
%b   // Binary
%o   // Octal
%x   // Hexadecimal
%f   // Floating point
%.2f // Floating point with 2 decimals
%s   // String
%q   // Quoted string
%p   // Pointer address
%%   // Literal %
```

---

**Quick Tip**: Keep this cheat sheet handy while coding! 🚀
