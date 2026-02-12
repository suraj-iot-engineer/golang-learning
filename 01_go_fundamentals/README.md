# 01. Go Fundamentals

This module covers the core building blocks of the Go programming language. We focus on writing idiomatic Go code from day one.

## 📚 Topics Covered

1.  **Variables & Zero Values**: Understanding strict typing and default values.
2.  **Control Structures**: `if`, `switch`, `for` loops (there is no `while`!).
3.  **Functions**: Multiple returns, named returns, and closures.
4.  **Pointers**: Understanding memory addresses and passing by reference.
5.  **Structs & Interfaces**: Go's approach to OOP (Composition over Inheritance).
6.  **Error Handling**: The `if err != nil` pattern.

## 🚀 Running the Examples

Each sub-topic is self-contained. You can run the main entry point to see all examples in action, or navigate to specific folders.

```bash
go run main.go
```

## 💻 Code Examples

### 1. Zero Values
Go gives variables a default value if not initialized.

```go
var i int     // 0
var s string  // ""
var b bool    // false
```

### 2. Error Handling
Go treats errors as values, not exceptions.

```go
file, err := os.Open("config.json")
if err != nil {
    log.Fatal(err)
}
```
