# Operators Guide

Go operators for arithmetic, comparison, and logic.

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

| Operator | Description | Example                 |
| :------- | :---------- | :---------------------- | ---------- | ----- | --- | ------------- |
| `&&`     | Logical AND | `true && false` (false) |
| `        |             | `                       | Logical OR | `true |     | false` (true) |
| `!`      | Logical NOT | `!true` (false)         |

## Assignment Operators

Used to assign values to variables.

| Operator | Description         | Example              |
| :------- | :------------------ | :------------------- |
| `=`      | Assign              | `x = 5`              |
| `+=`     | Add and assign      | `x += 3` (x = x + 3) |
| `-=`     | Subtract and assign | `x -= 2`             |
| `*=`     | Multiply and assign | `x *= 4`             |
| `/=`     | Divide and assign   | `x /= 2`             |

## Plan

1. Apply arithmetic operators in basic calculations.
2. Use comparison operators in `if` statements.
3. Combine conditions using logical operators.
4. Verify results using `fmt.Println`.

## Test

- Create a script using each operator type.
- Verify boolean results for logical and comparison operations.
