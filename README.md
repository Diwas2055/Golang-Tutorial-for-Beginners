# Go Tutorial

Go tutorial with code examples and a booking application.

## Content

- Go Introduction: History and use cases
- Setup: Installation and configuration
- Fundamentals: Variables and types
- Operators: Math, Pointers, Symbols
- Data Structures: Slices, maps, and structs
- Control Flow: Loops and conditionals
- Functions: Parameters and organization
- Concurrency: Goroutines and Channels

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
    ├── 02-operators.go         # Operators & Symbols
    ├── 03-variables.go         # Variables & Input
    ├── 04-arrays.go            # Slices & Arrays
    ├── 05-loops.go             # Iteration
    ├── 06-if-else.go           # Conditionals
    ├── 07-functions.go         # Functions
    ├── 08-maps.go              # Maps
    ├── 09-structs.go           # Structs
    ├── 10-concurrency.go       # Concurrency Primitives
    └── packages/               # Package organization
```

## Setup

1. Install Go: golang.org/dl
2. Verify: `go version`

## Run

```bash
# Run a chapter
go run chapters/01-basic.go

# Run the app
go mod init booking-app
go run main.go helper.go
```

## Plan

| Week   | Focus            | Files                    |
| :----- | :--------------- | :----------------------- |
| Week 1 | Foundations      | `chapters/01` to `03`    |
| Week 2 | Logic            | `chapters/04` to `06`    |
| Week 3 | Data & Functions | `chapters/07` to `09`    |
| Week 4 | Concurrency      | `chapters/10`, `main.go` |

## Test

- Run all chapter files sequentially.
- Verify operators in `02-operators.go`.
