# Cheat Sheet

Go syntax reference.

## Basic Syntax

- package main
- import "fmt"
- func main() { }

## Variables & Constants

- var name string = "John"
- name := "John"
- const Pi = 3.14

## Data Types

- bool, string
- int, uint, float64
- array [size]type
- slice []type
- map[key]value
- struct

## Control Flow

- if / else
- switch
- for loop
- for range

## Functions

- func name(params) returnType { }
- Multiple return values: func name() (int, error)

## Concurrency

- go function() (goroutine)
- chan type (channels)
- sync.WaitGroup

## Plan

| Topic    | Focus                            |
| :------- | :------------------------------- |
| Basics   | Syntax, Variables, Types         |
| Control  | If, Switch, Loops                |
| Advanced | Structs, Interfaces, Concurrency |

## Test

- Verify syntax with `go fmt`.
- Check for common issues with `go vet`.
