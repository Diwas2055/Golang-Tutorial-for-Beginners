# Go Programming Tutorial for Beginners

A comprehensive, hands-on tutorial to learn Go (Golang) from scratch with practical examples and a real-world booking application.

## 📚 What You'll Learn

### Part 1: Introduction & Fundamentals
- **Introduction to Go** - What is Go? Why Go? Real-world use cases
- **Setup & Installation** - Install Go and configure your development environment
- **First Program** - Write and run your first Go program
- **Variables & Constants** - Declaration, types, and best practices
- **Data Types** - Strings, integers, floats, booleans, and more
- **Input/Output** - User input and formatted output with `fmt`
- **Pointers** - Understanding memory addresses and references

### Part 2: Data Structures
- **Arrays** - Fixed-size collections
- **Slices** - Dynamic arrays and slice operations
- **Maps** - Key-value pairs and hash tables
- **Structs** - Custom data types and object-like structures

### Part 3: Control Flow
- **Loops** - For loops, range, break, and continue
- **Conditionals** - If/else statements and boolean logic
- **Switch Statements** - Multi-way branching
- **Input Validation** - Practical validation techniques

### Part 4: Functions & Organization
- **Functions** - Declaration, parameters, return values
- **Multiple Return Values** - Error handling patterns
- **Variadic Functions** - Variable number of arguments
- **Packages** - Code organization and modularity
- **Scope Rules** - Local, package, and exported scope

### Part 5: Advanced Concepts
- **Methods** - Attaching functions to types
- **Interfaces** - Polymorphism and abstraction
- **Error Handling** - Go's error handling philosophy
- **Defer, Panic, Recover** - Resource cleanup and error recovery

### Part 6: Concurrency
- **Goroutines** - Lightweight concurrent execution
- **Channels** - Communication between goroutines
- **WaitGroups** - Synchronization primitives
- **Select Statement** - Multiplexing channel operations

## 🗂️ Repository Structure

```
.
├── README.md                          # This file
├── Introduction.md                    # Go introduction and overview
├── go_syntax_concepts.md             # Detailed syntax reference
├── GO_DEEP_DIVE_GUIDE.md            # Comprehensive deep dive guide
├── main.go                           # Complete booking app with concurrency
├── helper.go                         # Helper functions
├── go-mod.txt                        # Module setup instructions
└── chapters/                         # Progressive learning chapters
    ├── 01-basic.go                   # Hello World
    ├── 02-variables.go               # Variables and user input
    ├── 03-arrays.go                  # Arrays and slices
    ├── 04-loops.go                   # Loops and iteration
    ├── 05-if-else.go                 # Conditionals and validation
    ├── 06-functions.go               # Functions and code organization
    ├── 07-maps.go                    # Maps and key-value storage
    ├── 08-structs.go                 # Structs and custom types
    └── packages/                     # Package organization example
        ├── main.go                   # Main package
        ├── go-mod.txt               # Module configuration
        └── helper/
            └── helper.go             # Helper package
```

## 🚀 Getting Started

### Prerequisites
- Basic programming knowledge
- Terminal/command line familiarity

### Installation

**Linux (Ubuntu/Debian):**
```bash
sudo apt update
sudo apt install golang-go
go version
```

**macOS:**
```bash
brew install go
go version
```

**Windows:**
Download from [golang.org/dl](https://golang.org/dl/)

### Running the Examples

1. **Clone or download this repository**

2. **Navigate to the project directory**
```bash
cd golang-tutorial
```

3. **Run individual chapters**
```bash
# Run a specific chapter
go run chapters/01-basic.go
go run chapters/02-variables.go
# ... and so on
```

4. **Run the complete application**
```bash
# Initialize module (first time only)
go mod init booking-app

# Run the main application
go run main.go helper.go
```

5. **Run the package example**
```bash
cd chapters/packages
go mod init booking-app
go run main.go
```

## 📖 Learning Path

### Beginner Track (Start Here)
1. Read `Introduction.md` - Understand what Go is and why it matters
2. Follow `GO_DEEP_DIVE_GUIDE.md` - Comprehensive guide with examples
3. Work through chapters sequentially:
   - `01-basic.go` → `02-variables.go` → ... → `08-structs.go`
4. Study the package organization in `chapters/packages/`
5. Examine the complete application in `main.go`

### Reference Materials
- `go_syntax_concepts.md` - Quick syntax reference
- `GO_DEEP_DIVE_GUIDE.md` - In-depth explanations with examples

## 🎯 Project: Conference Booking Application

Throughout this tutorial, you'll build a real conference ticket booking system that demonstrates:

- ✅ User input and validation
- ✅ Data storage with slices and structs
- ✅ Business logic with functions
- ✅ Code organization with packages
- ✅ Concurrent ticket processing with goroutines
- ✅ Error handling and edge cases

**Features:**
- Book tickets with user information
- Validate email and ticket availability
- Track remaining tickets
- Send confirmation emails (simulated with goroutines)
- Display all bookings

## 💡 Key Concepts Covered

| Concept | Chapter | Description |
|---------|---------|-------------|
| Hello World | 01 | Basic program structure |
| Variables | 02 | Variable declaration and types |
| Slices | 03 | Dynamic arrays |
| Loops | 04 | Iteration and range |
| Conditionals | 05 | If/else and validation |
| Functions | 06 | Code organization |
| Maps | 07 | Key-value storage |
| Structs | 08 | Custom data types |
| Packages | packages/ | Multi-file organization |
| Concurrency | main.go | Goroutines and WaitGroups |

## 🛠️ Best Practices Demonstrated

- ✨ Clean, readable code
- 📦 Proper package organization
- 🔍 Input validation
- ⚡ Efficient data structures
- 🔄 Concurrent programming
- 📝 Meaningful variable names
- 🎯 Single responsibility functions

## 📚 Additional Resources

- **Official Documentation**: [golang.org/doc](https://golang.org/doc)
- **Go by Example**: [gobyexample.com](https://gobyexample.com)
- **Effective Go**: [golang.org/doc/effective_go](https://golang.org/doc/effective_go)
- **Go Playground**: [play.golang.org](https://play.golang.org)
- **Standard Library**: [pkg.go.dev/std](https://pkg.go.dev/std)

## 🤝 Contributing

Feel free to submit issues or pull requests to improve this tutorial!

## 📄 License

This tutorial is open source and available for educational purposes.

---

**Happy Learning! Start with `Introduction.md` or dive into `GO_DEEP_DIVE_GUIDE.md` for a comprehensive guide! 🚀**