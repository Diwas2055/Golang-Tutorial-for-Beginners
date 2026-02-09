# Learning Path

Instructions for Go tutorial chapters.

## Chapter 01: Hello World (chapters/01-basic.go)

- Package declaration
- Import statements
- Main function
- Printing to console

Run:

```bash
go run chapters/01-basic.go
```

## Chapter 01.5: Operators (chapters/01.5-operators.go)

- Arithmetic operators (+, -, \*, /, %)
- Comparison operators (==, !=, <, >)
- Logical operators (&&, ||, !)
- Assignment and Short Declaration (=, :=, +=)
- Pointer operators (&, \*)
- Selector operator (.)

Run:

```bash
go run chapters/01.5-operators.go
```

## Chapter 02: Variables (chapters/02-variables.go)

- Variable declaration (var, :=)
- Constants (const)
- Data types (string, int, uint)
- User input (fmt.Scan)
- Formatted output (fmt.Printf)

Run:

```bash
go run chapters/02-variables.go
```

## Chapter 03: Arrays & Slices (chapters/03-arrays.go)

- Arrays and slices
- Slices (make, append)
- Slice operations

Run:

```bash
go run chapters/03-arrays.go
```

## Chapter 04: Loops (chapters/04-loops.go)

- For loop
- Infinite loop
- Range loop
- String manipulation (strings.Fields)

Run:

```bash
go run chapters/04-loops.go
```

## Chapter 05: Conditionals (chapters/05-if-else.go)

- if-else
- Boolean operators (&&, ||, !)
- Input validation
- break and continue
- switch

Run:

```bash
go run chapters/05-if-else.go
```

## Chapter 06: Functions (chapters/06-functions.go)

- Function declaration
- Parameters and return values
- Multiple return values
- Package-level variables

Run:

```bash
go run chapters/06-functions.go
```

## Chapter 07: Maps (chapters/07-maps.go)

- Map initialization (make)
- Map elements
- Type conversion (strconv)

Run:

```bash
go run chapters/07-maps.go
```

## Chapter 08: Structs (chapters/08-structs.go)

- Defining structs
- Struct instances
- Struct fields
- Slices of structs

Run:

```bash
go run chapters/08-structs.go
```

## Chapter 09: Concurrency Primitives (chapters/09-concurrency.go)

- Goroutines (go keyword)
- WaitGroups (sync.WaitGroup)
- Channels (chan)
- Buffered channels
- Closing channels

Run:

```bash
go run chapters/09-concurrency.go
```

## Chapter 10: Packages (chapters/packages/)

- Custom packages
- Exported functions
- Local packages
- Module initialization (go mod init)

Run:

```bash
cd chapters/packages
go mod init booking-app
go run main.go
```

## Chapter 11: Capstone Application (main.go)

- Integrated concurrency
- Error handling
- Input validation logic
- Modular design

Run:

```bash
go run main.go helper.go
```

## Plan

| Week   | Focus          |
| :----- | :------------- |
| Week 1 | Chapters 01-03 |
| Week 2 | Chapters 04-06 |
| Week 3 | Chapters 07-09 |
| Week 4 | Chapters 10-11 |

## Test

- Run each chapter file.
- Verify output matches expectations.
- Run main.go and test user input.
