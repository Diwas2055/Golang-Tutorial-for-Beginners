# Technical Guide

Detailed explanation of Go syntax and runtime behavior.

## Setup

- Binaries: Download from golang.org/dl.
- Path: Ensure `go/bin` is in the system PATH.
- Verify: `go version` confirms installation.

## Fundamentals

- **Packages**: Every Go file belongs to a package. `package main` defines an executable.
- **Functions**: `main()` is the program entry point.
- **Variables**: `var` for explicit type or package-level; `:=` for short declaration in functions.
- **Constants**: Immutable values defined with `const`.
- **Zero Values**: Uninitialized variables receive default values (0, "", false, nil).

## Data Structures

- **Arrays**: Fixed-length sequences of a single type. Size is part of the type.
- **Slices**: Dynamic windows into arrays. Created via `[]type` or `make()`. Use `append()` to grow.
- **Maps**: Unordered key-value pairs created with `make(map[K]V)`.
- **Structs**: Typed collections of fields used to represent objects.

## Concurrency primitives

Concurrency in Go is the execution of independent tasks. The runtime manages scheduling across OS threads.

### 1. Goroutines

A goroutine is a function executing concurrently with other code in the same address space.

- **Execution**: Prefix a function call with `go`.
- **Cost**: Starts with a small stack (approx. 2KB) that grows/shrinks dynamically.
- **Behavior**: The `main` function does not wait for goroutines to finish before exiting.

### 2. WaitGroups

The `sync.WaitGroup` type synchronizes the completion of multiple goroutines.

- **Add(n)**: Increments the counter by `n`. Call this before starting goroutines.
- **Done()**: Decrements the counter by 1. Usually called via `defer` inside the goroutine.
- **Wait()**: Blocks execution until the counter reaches 0.

### 3. Channels

Channels provide synchronized communication between goroutines. They ensure that values are passed without explicit locks or condition variables.

- **Initialization**: `ch := make(chan type)`.
- **Send**: `ch <- value` (blocks until a receiver is ready).
- **Receive**: `variable := <-ch` (blocks until a sender is ready).
- **Buffered Channels**: `make(chan type, capacity)` allows sending up to `capacity` values without blocking.
- **Closing**: `close(ch)` signals that no more values will be sent. Receiving from a closed channel returns the type's zero value.

### 4. Select Statement

The `select` statement lets a goroutine wait on multiple communication operations.

- It blocks until one of its cases can run, then executes that case.
- If multiple cases are ready, it chooses one pseudo-randomly.
- A `default` case can be used for non-blocking communication.

## Plan

1. Initialize environment using `go mod init`.
2. Apply basic syntax for input/output logic.
3. Organize data using structs and maps.
4. Implement asynchronous tasks using goroutines and channels.

## Test

- **Static Analysis**: Run `go vet` to check for suspicious constructs.
- **Race Detection**: Run `go run -race main.go` to find data races.
- **Execution**: Verify output consistency across multiple runs.
