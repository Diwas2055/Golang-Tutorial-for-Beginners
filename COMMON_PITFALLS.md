# Common Pitfalls & Troubleshooting Guide

A guide to common mistakes, errors, and how to fix them when learning Go.

---

## 🚨 Common Beginner Mistakes

### 1. Unused Variables and Imports

**Problem**:
```go
package main

import "fmt"
import "strings"  // unused

func main() {
    name := "John"  // declared but not used
    fmt.Println("Hello")
}
```

**Error**:
```
imported and not used: "strings"
name declared and not used
```

**Solution**:
```go
package main

import "fmt"

func main() {
    name := "John"
    fmt.Println("Hello", name)
}
```

**Why**: Go enforces clean code by not allowing unused variables or imports.

**Tip**: Use `_` to explicitly ignore values:
```go
_, err := someFunction()  // Ignore first return value
```

---

### 2. Short Declaration vs Assignment

**Problem**:
```go
func main() {
    name := "John"
    name := "Jane"  // Error: no new variables on left side
}
```

**Error**:
```
no new variables on left side of :=
```

**Solution**:
```go
func main() {
    name := "John"  // Declaration
    name = "Jane"   // Assignment
}
```

**Rule**:
- `:=` declares NEW variables
- `=` assigns to EXISTING variables
- `:=` can only be used inside functions

---

### 3. Modifying Loop Variables

**Problem**:
```go
numbers := []int{1, 2, 3}
for _, num := range numbers {
    num = num * 2  // Doesn't modify original
}
fmt.Println(numbers)  // Still [1, 2, 3]
```

**Solution**:
```go
numbers := []int{1, 2, 3}
for i := range numbers {
    numbers[i] = numbers[i] * 2
}
fmt.Println(numbers)  // [2, 4, 6]
```

**Why**: `range` creates a copy of each element.

---

### 4. Nil Slice vs Empty Slice

**Problem**:
```go
var slice1 []int        // nil slice
slice2 := []int{}       // empty slice

fmt.Println(slice1 == nil)  // true
fmt.Println(slice2 == nil)  // false
```

**Confusion**: Both have length 0, but they're different!

**Best Practice**:
```go
// For function returns, prefer nil
func getItems() []int {
    return nil  // Not []int{}
}

// For initialization, either works
items := []int{}
items := make([]int, 0)
```

---

### 5. Slice Append Gotcha

**Problem**:
```go
slice1 := []int{1, 2, 3}
slice2 := slice1
slice2 = append(slice2, 4)

// slice1 might be modified!
```

**Why**: Slices share underlying arrays.

**Solution**:
```go
// Create a true copy
slice1 := []int{1, 2, 3}
slice2 := make([]int, len(slice1))
copy(slice2, slice1)
slice2 = append(slice2, 4)
```

---

### 6. Map Concurrent Access

**Problem**:
```go
m := make(map[string]int)

go func() {
    m["key"] = 1  // Concurrent write
}()

go func() {
    m["key"] = 2  // Concurrent write
}()
// Fatal error: concurrent map writes
```

**Solution**:
```go
import "sync"

var (
    m  = make(map[string]int)
    mu sync.Mutex
)

go func() {
    mu.Lock()
    m["key"] = 1
    mu.Unlock()
}()

go func() {
    mu.Lock()
    m["key"] = 2
    mu.Unlock()
}()
```

**Or use sync.Map**:
```go
var m sync.Map

m.Store("key", 1)
value, ok := m.Load("key")
```

---

### 7. Goroutine Loop Variable Capture

**Problem**:
```go
for i := 0; i < 5; i++ {
    go func() {
        fmt.Println(i)  // Prints 5, 5, 5, 5, 5
    }()
}
time.Sleep(time.Second)
```

**Why**: All goroutines share the same `i` variable.

**Solution 1**: Pass as parameter
```go
for i := 0; i < 5; i++ {
    go func(n int) {
        fmt.Println(n)
    }(i)
}
```

**Solution 2**: Create local copy
```go
for i := 0; i < 5; i++ {
    i := i  // Create new variable
    go func() {
        fmt.Println(i)
    }()
}
```

---

### 8. Forgetting WaitGroup.Done()

**Problem**:
```go
var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(n int) {
        fmt.Println(n)
        // Forgot wg.Done()
    }(i)
}

wg.Wait()  // Deadlock!
```

**Solution**:
```go
var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(n int) {
        defer wg.Done()  // Always use defer
        fmt.Println(n)
    }(i)
}

wg.Wait()
```

---

### 9. Pointer vs Value Receivers

**Problem**:
```go
type Counter struct {
    count int
}

func (c Counter) Increment() {
    c.count++  // Doesn't modify original
}

func main() {
    c := Counter{}
    c.Increment()
    fmt.Println(c.count)  // Still 0
}
```

**Solution**:
```go
func (c *Counter) Increment() {
    c.count++
}
```

**Rule**:
- Use pointer receiver to modify
- Use value receiver for read-only
- Be consistent within a type

---

### 10. Interface Nil Check

**Problem**:
```go
func returnsError() error {
    var err *MyError = nil
    return err  // Returns non-nil interface!
}

func main() {
    if err := returnsError(); err != nil {
        fmt.Println("Error!")  // This prints!
    }
}
```

**Why**: Interface contains type information even if value is nil.

**Solution**:
```go
func returnsError() error {
    var err *MyError = nil
    if err == nil {
        return nil  // Return nil interface
    }
    return err
}
```

---

## 🔧 Common Compilation Errors

### Error: "undefined: variable"

**Cause**: Variable not declared or out of scope

**Fix**:
```go
// Wrong
if x := 5; x > 0 {
    fmt.Println("Positive")
}
fmt.Println(x)  // Error: x not defined

// Right
x := 5
if x > 0 {
    fmt.Println("Positive")
}
fmt.Println(x)  // OK
```

---

### Error: "cannot use X (type Y) as type Z"

**Cause**: Type mismatch

**Fix**:
```go
// Wrong
var i int = 42
var f float64 = i  // Error

// Right
var i int = 42
var f float64 = float64(i)  // Explicit conversion
```

---

### Error: "missing return"

**Cause**: Not all code paths return a value

**Fix**:
```go
// Wrong
func divide(a, b int) int {
    if b != 0 {
        return a / b
    }
    // Missing return here
}

// Right
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

---

### Error: "too many arguments" or "not enough arguments"

**Cause**: Function call doesn't match signature

**Fix**:
```go
func greet(name string, age int) {
    fmt.Printf("%s is %d years old\n", name, age)
}

// Wrong
greet("John")  // Not enough arguments

// Right
greet("John", 30)
```

---

## 🐛 Runtime Errors

### Panic: "index out of range"

**Cause**: Accessing invalid slice/array index

**Fix**:
```go
// Wrong
slice := []int{1, 2, 3}
fmt.Println(slice[5])  // Panic!

// Right
if len(slice) > 5 {
    fmt.Println(slice[5])
} else {
    fmt.Println("Index out of range")
}
```

---

### Panic: "nil pointer dereference"

**Cause**: Dereferencing nil pointer

**Fix**:
```go
// Wrong
var p *int
fmt.Println(*p)  // Panic!

// Right
var p *int
if p != nil {
    fmt.Println(*p)
} else {
    fmt.Println("Pointer is nil")
}
```

---

### Panic: "send on closed channel"

**Cause**: Sending to a closed channel

**Fix**:
```go
ch := make(chan int)
close(ch)
// ch <- 1  // Panic!

// Use defer to ensure channel is closed only once
defer func() {
    if r := recover(); r != nil {
        fmt.Println("Recovered:", r)
    }
}()
```

---

### Deadlock: "all goroutines are asleep"

**Causes**:
1. Waiting on channel that never receives
2. WaitGroup never reaches zero
3. Mutex never unlocked

**Fix**:
```go
// Wrong
ch := make(chan int)
<-ch  // Deadlock! No one sends

// Right
ch := make(chan int)
go func() {
    ch <- 42
}()
<-ch
```

---

## 💡 Best Practices to Avoid Issues

### 1. Always Handle Errors

```go
// Bad
result, _ := doSomething()

// Good
result, err := doSomething()
if err != nil {
    return fmt.Errorf("failed to do something: %w", err)
}
```

### 2. Use defer for Cleanup

```go
file, err := os.Open("file.txt")
if err != nil {
    return err
}
defer file.Close()  // Always closes, even on error

// Use file...
```

### 3. Initialize Maps Before Use

```go
// Bad
var m map[string]int
m["key"] = 1  // Panic!

// Good
m := make(map[string]int)
m["key"] = 1
```

### 4. Check Slice Length Before Access

```go
if len(slice) > index {
    value := slice[index]
}
```

### 5. Use Buffered Channels Carefully

```go
// Unbuffered - blocks until received
ch := make(chan int)

// Buffered - blocks only when full
ch := make(chan int, 10)
```

### 6. Don't Copy Mutexes

```go
// Bad
type Counter struct {
    mu    sync.Mutex
    count int
}

func (c Counter) Increment() {  // Copies mutex!
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++
}

// Good
func (c *Counter) Increment() {  // Pointer receiver
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++
}
```

---

## 🔍 Debugging Tips

### 1. Print Debugging

```go
fmt.Printf("DEBUG: variable = %+v\n", variable)
fmt.Printf("DEBUG: type = %T\n", variable)
```

### 2. Use go vet

```bash
go vet ./...
```

Catches common mistakes like:
- Printf format errors
- Unreachable code
- Suspicious constructs

### 3. Use go fmt

```bash
go fmt ./...
```

Automatically formats code correctly.

### 4. Enable Race Detector

```bash
go run -race main.go
go test -race ./...
```

Detects data races in concurrent code.

### 5. Use Delve Debugger

```bash
go install github.com/go-delve/delve/cmd/dlv@latest
dlv debug main.go
```

---

## 📋 Checklist Before Running

- [ ] All variables used
- [ ] All imports used
- [ ] Error handling in place
- [ ] Maps initialized
- [ ] Channels properly closed
- [ ] WaitGroups balanced (Add/Done)
- [ ] Mutexes unlocked
- [ ] Defer statements for cleanup
- [ ] Slice bounds checked
- [ ] Nil checks for pointers

---

## 🆘 When You're Stuck

### 1. Read the Error Message
Go's error messages are usually clear. Read them carefully!

### 2. Check the Documentation
```bash
go doc package.Function
```

### 3. Use the Playground
Test small snippets at [play.golang.org](https://play.golang.org)

### 4. Search for the Error
Copy the error message and search online.

### 5. Ask for Help
- Stack Overflow
- Reddit r/golang
- Go Forum
- Discord servers

---

## 📚 Common Error Patterns

### Pattern 1: "cannot assign to X"

**Means**: Trying to modify immutable value

**Example**:
```go
// Wrong
for _, item := range items {
    item.field = value  // Can't modify copy
}

// Right
for i := range items {
    items[i].field = value
}
```

---

### Pattern 2: "invalid operation"

**Means**: Operation not supported for type

**Example**:
```go
// Wrong
var x interface{} = 5
y := x + 1  // Error

// Right
var x interface{} = 5
if num, ok := x.(int); ok {
    y := num + 1
}
```

---

### Pattern 3: "mismatched types"

**Means**: Types don't match in operation

**Example**:
```go
// Wrong
var a int = 5
var b int64 = 10
c := a + b  // Error

// Right
var a int = 5
var b int64 = 10
c := int64(a) + b
```

---

## 🎯 Quick Fixes

| Problem | Quick Fix |
|---------|-----------|
| Unused variable | Use it or remove it |
| Unused import | Remove it or use `_` |
| Type mismatch | Add type conversion |
| Nil pointer | Add nil check |
| Index out of range | Check length first |
| Deadlock | Check channel/WaitGroup usage |
| Race condition | Add mutex or use channels |
| Map panic | Initialize with `make()` |

---

## 🚀 Moving Forward

Remember:
- Errors are learning opportunities
- Read error messages carefully
- Test small pieces of code
- Use the tools Go provides
- Ask for help when stuck

**Most importantly**: Don't get discouraged! Every Go developer has made these mistakes. 💪

---

**Happy Debugging! 🐛🔨**
