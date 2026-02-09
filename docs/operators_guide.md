# Operators Guide

Go operators for arithmetic, comparison, logic, and memory access.

## Arithmetic Operators

Used for mathematical calculations.

| Operator | Description    | Example       |
| :------- | :------------- | :------------ |
| `+`      | Addition       | `10 + 5 = 15` |
| `-`      | Subtraction    | `10 - 5 = 5`  |
| `*`      | Multiplication | `10 * 5 = 50` |
| `/`      | Division       | `10 / 5 = 2`  |
| `%`      | Modulus        | `10 % 3 = 1`  |
| `++`     | Increment      | `i++`         |
| `--`     | Decrement      | `i--`         |

## Comparison Operators

Used to compare two values. Returns a boolean (`true` or `false`).

| Operator | Description      | Example         |
| :------- | :--------------- | :-------------- |
| `==`     | Equal            | `5 == 5` (true) |
| `!=`     | Not equal        | `5 != 3` (true) |
| `>`      | Greater than     | `5 > 3` (true)  |
| `<`      | Less than        | `3 < 5` (true)  |
| `>=`     | Greater or equal | `5 >= 5` (true) |
| `<=`     | Less or equal    | `3 <= 5` (true) |

## Logical Operators

Used to combine multiple conditions.

| Operator | Description | Example                  |
| :------- | :---------- | :----------------------- |
| `&&`     | Logical AND | `true && false` (false)  |
| `\|\|`   | Logical OR  | `true \|\| false` (true) |
| `!`      | Logical NOT | `!true` (false)          |

## Assignment Operators

Used to assign or declare variables.

| Operator | Description         | Example                             |
| :------- | :------------------ | :---------------------------------- |
| `=`      | Assignment          | `x = 5`                             |
| `:=`     | Short Declaration   | `name := "Go"` (Declare and assign) |
| `+=`     | Add and assign      | `x += 3` (x = x + 3)                |
| `-=`     | Subtract and assign | `x -= 2`                            |
| `*=`     | Multiply and assign | `x *= 4`                            |
| `/=`     | Divide and assign   | `x /= 2`                            |

## Address and Pointer Operators

Used for memory access and references.

| Operator | Description         | Example                            |
| :------- | :------------------ | :--------------------------------- |
| `&`      | Address-of          | `&variable` (Gets memory address)  |
| `*`      | Pointer/Dereference | `*pointer` (Gets value at address) |

## Selector Operator

Used to access members of a struct or package.

| Operator | Description | Example                           |
| :------- | :---------- | :-------------------------------- |
| `.`      | Selector    | `user.firstName` or `fmt.Println` |

## Plan

1. Apply arithmetic and assignment operators.
2. Use comparison and logical operators in control flow.
3. Access memory addresses using `&` and `*`.
4. Access struct fields using the selector `.`.

## Test

- Run `go run chapters/01.5-operators.go`.
- Verify pointer values match variable values after dereferencing.
