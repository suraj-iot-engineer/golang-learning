# 04. Testing & Debugging

Writing reliable software requires rigorous testing. Go has a built-in testing framework that is simple yet powerful.

## 📚 Topics Covered

1.  **Unit Testing**: Writing `xx_test.go` files, `t.Run`, and assertions.
2.  **Table-Driven Tests**: The idiomatic way to structure Go tests.
3.  **Benchmarking**: Measuring performance with `testing.B`.
4.  **Fuzzing**: Automated random input testing (Go 1.18+).
5.  **Integration Testing**: Testing components together.

## 🛠️ Running Tests

```bash
# Run all tests in the current directory
go test -v

# Run with coverage
go test -cover

# Run benchmarks
go test -bench=.
```

## 🚀 Example
The example includes a calculator package and a comprehensive test suite using table-driven tests.
