# Operators & Symbols Guide

Reference for Go operators, punctuation, and special symbols.

## Arithmetic & Assignment

| Symbol              | Purpose          | Example                            |
| :------------------ | :--------------- | :--------------------------------- |
| `+` `-` `*` `/` `%` | Basic Math       | `a + b`, `10 % 3`                  |
| `=`                 | Assignment       | `x = 5`                            |
| `:=`                | Declare + Assign | `name := "Go"` (Short declaration) |
| `+=` `-=` `*=` `/=` | Update + Assign  | `x += 2` (x = x + 2)               |
| `++` `--`           | Inc/Dec          | `i++`, `i--` (Statements only)     |

## Comparison & Logical

| Symbol            | Purpose     | Example                   |
| :---------------- | :---------- | :------------------------ |
| `==`              | Equality    | `a == b` (Compare values) |
| `!=`              | Inequality  | `a != b`                  |
| `>` `<` `>=` `<=` | Relational  | `x > 10`, `y <= 5`        |
| `&&`              | Logical AND | `cond1 && cond2`          |
| `\|\|`            | Logical OR  | `cond1 \|\| cond2`        |
| `!`               | Logical NOT | `!isValid`                |

## Pointers & Memory

| Symbol | Purpose             | Example                              |
| :----- | :------------------ | :----------------------------------- |
| `&`    | Address-of          | `&x` (Get memory address of x)       |
| `*`    | Pointer/Dereference | `*ptr` (Get value stored at address) |

## Structural & Special

| Symbol | Purpose         | Example                                   |
| :----- | :-------------- | :---------------------------------------- |
| `.`    | Access/Selector | `fmt.Println`, `user.name`                |
| `_`    | Ignore          | `val, _ := fn()` (Discard return value)   |
| `...`  | Variadic/Expand | `func fn(args ...int)`, `slice...`        |
| `[]`   | Index/Slice     | `arr[0]`, `[]string{}` (Types & indexing) |

## Plan

1. Use `:=` for internal function declarations.
2. Apply `&` and `*` for memory manipulation.
3. Utilize `_` to keep code clean by ignoring unused returns.
4. Access struct fields or package functions using `.`.

## Test

- Run `go run chapters/02-operators.go`.
- Check for "unused variable" errors to ensure `_` is used correctly.
