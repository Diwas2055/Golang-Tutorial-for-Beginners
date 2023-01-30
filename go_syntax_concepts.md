# GO Syntax And Concepts

## Install Go 

- Download Go from [here](https://golang.org/dl/)

> For Guide to install Go Version Manager (GVM) on Linux [click here](https://github.com/moovweb/gvm)

- Install Go on Linux Using Apt command
```bash
sudo apt install golang-go
```
- Uninstall Go on Linux Using Apt command
```bash
sudo apt remove --autoremove go
```
- To check if Go is installed or not
```bash
go version
```
- To check the path of Go
```bash
which go
```
- To check the location of Go
```bash
go env GOROOT
```
- To check the location of Go workspace
```bash
go env GOPATH
```
- To check the location of Go bin
```bash
go env GOBIN
```
- To check the location of Go src
```bash
go env GOSRC
```
- To check the location of Go pkg
```bash
go env GOPKG
```
> For more information about Go environment variables [click here](https://golang.org/doc/install/source#environment)


> Note : Install `Go` extension in VS Code for better experience

## Write our First Program & Structure of a Go File

- Create a folder `booking_app` and open it in VS Code
```bash
mkdir booking_app
cd booking_app
code .
```
- Create a file `main.go` and write the following code

- Initialize a module
```bash
go mod init booking-app // booking-app is the name of the module . It create a go.mod file
```
- To run the program
```bash
`go run main.go`
```

## Variables

- Variables are used to store values
- Like container for values
- Give the variable a name and reference everywhere in the app
```go
var name string = "John"
var age int = 25
var isCool bool = true
```
- Short hand declaration
```go
name := "John"
age := 25
isCool := true
```
- Multiple variables
```go
var name, email string = "John", "
```
- Multiple variables short hand
```go
name, email := "John", "
```
## Formatted Output - printf 
```go
name := "John"
age := 25
isCool := true
fmt.Printf("Some text %s %d %t", name, age, isCool)
```

## Data Types
List of Data Types in Go
- bool
    - For boolean values like (true or false)
- string
    - For text like "Hello World"
- int
    - For integers like 1, 2, 3
- float
    - For floating point numbers like 1.2, 3.4
- uint
    - For unsigned integers like 1, 2, 3 cannot be negative
- maps
    - For key value pairs like {"name": "John"}
- arrays
    - For lists of values like [1, 2, 3]

> Note : Check the type of variable by using `reflect.TypeOf()`
```go
import "reflect"
var name = "John"
fmt.Println(reflect.TypeOf(name))
```

## Getting User Input
```go
var name string
fmt.Print("Enter your name: ")
fmt.Scan(&name) // `&` is used to get the address of the variable
fmt.Println("Hello", name)
```
## What is a Pointer?

- A `pointer` is a reference to a value
- It is a memory address
- It is a variable that stores the memory address of another variable
```go
var name string = "John"
var namePointer *string = &name // * is used to get the address of the variable
fmt.Println(namePointer)
```
- To get the value of the variable using pointer
```go
var name string = "John"
var namePointer *string = &name
fmt.Println(*namePointer) // `*` is used to get the value of the variable
```

## Arrays & Slices
### Arrays
- Arrays are fixed length
- Arrays are used to store multiple values of the same type
```go
var fruits [2]string
fruits[0] = "Apple"
fruits[1] = "Orange"
fmt.Println(fruits)
```
- Declare and assign values
```go
fruits := [2]string{"Apple", "Orange"}
fmt.Println(fruits)
```
- Get length of array
```go
fmt.Println(len(fruits))
```
### Slices
- Slices are like arrays but they are dynamic
- Slice is an abstraction of an array
- Slices are more flexible and powerful:variable-length or get an sub-array of its own
- Slice are also index-based and have a size,but is resizable when needed

```go
var fruits []string
fruits = append(fruits, "Apple")
fruits = append(fruits, "Orange")
fmt.Println(fruits)
```
- Declare and assign values
```go
fruits := []string{"Apple", "Orange"}
fmt.Println(fruits)
```
- Get length of slice
```go
fmt.Println(len(fruits))
```
- Get capacity of slice
```go
fmt.Println(cap(fruits))
```
- Create a slice using make
```go
fruits := make([]string, 2) 
fruits[0] = "Apple"
fruits[1] = "Orange"
fmt.Println(fruits)
```
- Create a slice using make with length and capacity
```go
fruits := make([]string, 2, 3)
fruits[0] = "Apple"
fruits[1] = "Orange"
fmt.Println(fruits)
```
- Create a slice using array
```go
fruitsArray := [2]string{"Apple", "Orange"}
fruits := fruitsArray[:]
fmt.Println(fruits)
```

## Loops & Conditionals
### Loop
- Loop is used to repeat a block of code
- Loop statements allows us to execute a statement or group of statements multiple times
### For Loop
```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```
### Foreach Loop
```go
fruits := []string{"Apple", "Orange"}
for index, fruit := range fruits {
    fmt.Println(index, fruit)
}
```
> Note : If you don't need the index then use `_` instead of `index`
```go
fruits := []string{"Apple", "Orange"}
for _, fruit := range fruits {
    fmt.Println(fruit)
}
```
> `Range` is iterates over elements for different data structures (son not only for arrays and slices)

## Conditionals (if / else) and Boolean Data Type

### if else statement
```go
if age >= 18 {
    fmt.Println("You are old enough to vote")
} else {
    fmt.Println("You are not old enough to vote")
}
```
### boolean data type
```go
var isCool bool = true
```
## Validate User Input
```go
var age int
fmt.Print("Enter your age: ")
fmt.Scan(&age)

if age >= 18 {
    fmt.Println("You are old enough to vote")
} else {
    fmt.Println("You are not old enough to vote")
}
```
## Switch Statement

- Switch statement is used to perform different actions based on different conditions
- Switch statement is like a multi-way if-else statement
- Switch statement is more readable than nested if-else statements
```go
var age int
fmt.Print("Enter your age: ")
fmt.Scan(&age)

switch age {
case age>=18:
    fmt.Println("Go to school")
case age==18:
    fmt.Println("Go to college")
default: // default is optional
    fmt.Println("Go to work")
}
```

## Functions
- Functions are used to perform a specific task
- Functions are reusable
- Functions are used to avoid code duplication
- Functions are used to make code more readable
- Functions are used to group together a set of related statements to perform a specific task

### Function Declaration
```go
func sayHello() {
    fmt.Println("Hello World")
}
```
### Function Call
```go
sayHello()
```
### Function with Parameters
```go
func sayHello(name string) {
    fmt.Println("Hello", name)
}
```
### Function with Return Value
```go
func getSum(num1 int, num2 int) int {
    return num1 + num2
}
```
### Function with Multiple Return Values
```go
func getSumAndDiff(num1 int, num2 int) (int, int) {
    sum := num1 + num2
    diff := num1 - num2
    return sum, diff
}
```
### Function with Named Return Values
```go
func getSumAndDiff(num1 int, num2 int) (sum int, diff int) {
    sum = num1 + num2
    diff = num1 - num2
    return
}
```
### Function with Variadic Parameters
```go
func getSum(nums ...int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    return sum
}
```
> Note : In Go, a function can have a `variadic` parameter, which allows it to accept zero or more arguments of a specified type. The type is indicated by appending an `ellipsis (...)` to the type name.
```go
// Example of a function with a variadic parameter
import "fmt"

func sum(numbers ...int) int {
  result := 0
  for _, n := range numbers {
    result += n
  }
  return result
}

func main() {
  fmt.Println(sum(1, 2, 3, 4, 5))
}
```
## Organize Code with Go Packages

### Package Level Variables
- Package level variables are declared outside of any function

### Package in GO
- GO programs are organized into packages
- A package is a collection of related source files

### Scope: Package Level
- Package level variables are accessible throughout the package
- Variables and function defined outside any function,con be accessed in all other files within the same package

> Note : To run multiple files in the same package, you need to run `go run *.go` instead of `go run main.go` or `go run file1.go file2.go`
```go
go run *.go // run all files in the same package

// or

go run file1.go file2.go // run specific files in the same package
```
### Multiple Packages in GO

- GO programs can have multiple packages
- It helps to organize code into different packages
- Logically group related code into packages

> Note : Exporting a function or variable from a package is done by `capitalizing` the first letter of the function or variable name

## Scope Rules in Go

- Scope is the region of the program where a variable is accessible
- Variables can be declared at different levels in a program

Three types of scopes in Go
- Local Scope
- Package Scope
- Global Scope

### Local Scope
- Variables declared inside a function are accessible only inside that function
- Local variables are created when the function is called and destroyed when the function is completed
- Declaration within block

### Package Scope
- Variables declared outside of any function are accessible throughout the package
- Package level variables are accessible throughout the package

### Global Scope
- Variables declared outside of any function and outside of any package are accessible throughout the program
- Global variables are accessible throughout the program

## Maps in Go

- Maps are used to store key-value pairs
- Maps are unordered
- Maps are reference types

### Map Declaration
```go
var person map[string]string
```
### Map Initialization
```go
person := make(map[string]string)
```
### Map Initialization with Values
```go
person := map[string]string{
    "name": "John",
    "age":  "25",
}
```
### Map Access
```go
person := map[string]string{
    "name": "John",
    "age":  "25",
}

fmt.Println(person["name"])
```
### Map Add
```go
person := map[string]string{
    "name": "John",
    "age":  "25",
}

person["address"] = "New York"
```
### Map Update
```go
person := map[string]string{
    "name": "John",
    "age":  "25",
}

person["age"] = "30"
```
### Map Delete
```go
person := map[string]string{
    "name": "John",
    "age":  "25",
}

delete(person, "age")
```
### Map Length
```go
person := map[string]string{
    "name": "John",
    "age":  "25",
}

fmt.Println(len(person))
```
### Map Iteration
```go
person := map[string]string{
    "name": "John",
    "age":  "25",
}

for key, value := range person {
    fmt.Printf("Key: %s, Value: %s\n", key, value)
}
```
### Map Check Key Exists
```go
person := map[string]string{
    "name": "John",
    "age":  "25",
}

if _, ok := person["name"]; ok {
    fmt.Println("Key exists")
} else {
    fmt.Println("Key does not exist")
}
```

## Structs in Go

- Structs are used to represent a record
- Structs are value types
- Structs are used to group together related data

### Struct Declaration
```go
type Person struct {
    name string
    age  int
}
```
### Struct Initialization
```go
var person Person
person.name = "John"
person.age = 25
```
### Struct Initialization with Values
```go
person := Person{
    name: "John",
    age:  25,
}
```
### Struct Access
```go
person := Person{
    name: "John",
    age:  25,
}

fmt.Println(person.name)
```
### Struct Update
```go
person := Person{
    name: "John",
    age:  25,
}

person.age = 30
```
### Struct Pointer
```go
person := Person{
    name: "John",
    age:  25,
}

personPointer := &person
personPointer.age = 30
```
> Note : Difference between `struct` and `struct pointer` is that, when you use `struct pointer`, you don't need to use `&` to access the struct fields
```go
person := Person{
    name: "John",
    age:  25,
}

personPointer := &person

personPointer.age = 30

// or

person.age = 30
```
### Struct Literals
```go
person := Person{
    name: "John",
    age:  25,
}

person2 := Person{
    name: "Steve",
    age:  30,
}

person3 := Person{
    name: "Maria",
    age:  28,
}
```

## Interfaces in Go

- Interfaces are used to define a set of methods
- Interfaces are implemented implicitly

### Interface Declaration
```go

type Shape interface {
    area() float64
}
```

### Interface Implementation
```go

type Rectangle struct {
    width  float64
    height float64
}

func (r Rectangle) area() float64 {
    return r.width * r.height
}

type Circle struct {
    radius float64
}

func (c Circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
```

### Interface Usage
```go

type Shape interface {
    area() float64
}

func printArea(s Shape) {
    fmt.Println(s.area())
}

func main() {
    r := Rectangle{width: 10, height: 5}
    c := Circle{radius: 10}

    printArea(r)
    printArea(c)
}
```

## Concurrency in Go

- Concurrency is the ability to run multiple tasks at the same time

### Goroutines

- Goroutines are lightweight threads managed by the Go runtime

### Goroutine Declaration
```go

func printHello() {
    fmt.Println("Hello")
}

func main() {
    go printHello()
    time.Sleep(100 * time.Millisecond)
}
```

### Channels

- Channels are used to communicate between goroutines
- Channels are typed
- Channels are unbuffered by default

### Channel Declaration
```go

func printHello(ch chan string) {
    ch <- "Hello"
}

func main() {
    ch := make(chan string)
    go printHello(ch)
    fmt.Println(<-ch)
}
```

### Buffered Channels

- Channels can be buffered
- Buffered channels can hold multiple values
- Buffered channels are created by passing a buffer size to the make function

### Buffered Channel Declaration
```go

func printHello(ch chan string) {
    ch <- "Hello"
}

func main() {
    ch := make(chan string, 2) 
    ch <- "Hello" // <- this will not block
    ch <- "Hello" 
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}
```