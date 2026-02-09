# Troubleshooting

Common errors and solutions in Go.

## Errors

### 1. Unused Variables and Imports

Go does not allow unused variables or imports.

Error:

```
imported and not used: "strings"
name declared and not used
```

Solution:
Remove the variable or import, or use it.

### 2. Short Declaration vs Assignment

- `:=` declares new variables.
- `=` assigns to existing variables.

Error:

```
no new variables on left side of :=
```

### 3. Modifying Loop Variables

`range` creates a copy of each element. To modify the original, use the index.

### 4. Nil Slice vs Empty Slice

- `var slice []int` is nil.
- `slice := []int{}` is not nil.

### 5. Concurrent Map Access

Maps are not thread-safe. Use `sync.Mutex` or `sync.Map`.

### 6. WaitGroup Deadlock

Ensure `wg.Done()` is called for every `wg.Add(1)`. Use `defer wg.Done()`.

## Plan

1. Identify the error type (Compilation or Runtime).
2. Check variable scope and initialization.
3. Apply standard Go patterns.

## Test

- Use `go vet ./...` to detect common issues.
- Use `go run -race main.go` to detect data races.
