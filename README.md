# Go Tutorial

Go tutorial with code examples and a booking application.

## Content

- Go Introduction: Go history and use cases
- Setup: Installation and environment configuration
- Fundamentals: Variables, constants, and data types
- Operators: Arithmetic, comparison, and logical operators
- Data Structures: Arrays, slices, maps, and structs
- Control Flow: Loops, conditionals, and input validation
- Functions: Parameters, return values, and packages
- Concurrency: Goroutines, Channels, and WaitGroups

## Project Structure

```text
.
├── README.md                   # Project overview
├── docs/                       # Documentation guides
│   ├── deep_dive_guide.md      # Technical guide
│   ├── chapter_guide.md        # Sequential instructions
│   ├── operators_guide.md      # Operators reference
│   ├── cheatsheet.md           # Syntax reference
│   ├── syntax_reference.md     # Language details
│   ├── exercises_and_projects.md # Exercises
│   └── common_pitfalls.md      # Troubleshooting
│
├── main.go                     # Booking App
├── helper.go                   # Logic helpers
├── go-mod.txt                  # Module instructions
│
└── chapters/                   # Code examples
    ├── 01-basic.go             # Hello World
    ├── 01.5-operators.go       # Operators
    ├── 02-variables.go         # Variables
    ├── 03-arrays.go            # Slices
    ├── 04-loops.go             # Iteration
    ├── 05-if-else.go           # Conditionals
    ├── 06-functions.go         # Functions
    ├── 07-maps.go              # Maps
    ├── 08-structs.go           # Structs
    ├── 09-concurrency.go       # Concurrency Primitives
    └── packages/               # Package organization
```

## Setup

### 1. Install Go

- macOS: `brew install go`
- Linux: `sudo apt install golang-go`
- Windows: Download from golang.org/dl

Verify: `go version`

### 2. Run code

```bash
go run chapters/01-basic.go
```

### 3. Booking App

```bash
go mod init booking-app
go run main.go helper.go
```

## Plan

| Week   | Focus              | Files                           |
| :----- | :----------------- | :------------------------------ |
| Week 1 | Foundations        | `chapters/01` to `03`           |
| Week 2 | Control Flow       | `chapters/04` to `06`           |
| Week 3 | Data & Concurrency | `chapters/07` to `09`           |
| Week 4 | Projects           | `chapters/packages/`, `main.go` |

## Test

- Verify documentation links in `README.md`.
- Run `go run main.go` to ensure functionality.
