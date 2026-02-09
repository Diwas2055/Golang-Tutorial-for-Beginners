# Syntax Reference

Go language specification.

## Setup

- CLI: `go version`, `go env`.
- Module: `go mod init`.
- Execute: `go run`, `go build`.

## Variables

- Declaration: `var name string`.
- Short hand: `name := "value"`.
- Pointers: `&variable`, `*pointer`.

## Data Types

- bool, string, int, float64, uint.
- array, slice, map, struct.

## Control Flow

- if / else
- for loops
- switch statements

## Functions

- Declaration: `func name()`.
- Return values: `func name() int`.
- Multi-return: `func name() (int, error)`.

## Packages

- Import: `import "name"`.
- Export: Capitalize first letter.

## Concurrency

- Goroutines: `go function()`.
- Channels: `chan type`, `<-ch`, `ch <-`.
- Select: `select { case ... }`.

## Plan

1. Review CLI commands.
2. Complete variable exercises.
3. Implement data structure examples.
4. Integrate concurrency.

## Test

- Run all scripts in `chapters/`.
- Validate logic in `main.go`.
