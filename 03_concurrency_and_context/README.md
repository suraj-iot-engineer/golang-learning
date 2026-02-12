# 03. Concurrency & Context

Go's superpower is its first-class support for concurrency. This module moves beyond basic goroutines to production-grade concurrent patterns.

## 📚 Topics Covered

1.  **Goroutines**: Lightweight threads managed by the Go runtime.
2.  **Channels**: Typed conduits for synchronization and communication.
    - Unbuffered vs Buffered
    - Closing channels
    - `range` over channels
3.  **Select Statement**: Handling multiple channel operations.
4.  **Sync Package**: `WaitGroup`, `Mutex` for safe data access.
5.  **Context API**: Managing timeouts, deadlines, and cancellation signals.

## 🔑 Key Concepts

### The "Share Memory By Communicating" Philosophy
Don't communicate by sharing memory; share memory by communicating.

### Context
Always pass `context.Context` as the first argument to functions that involve I/O or long-running operations.

## 🚀 Running the Example

The example simulates a concurrent worker pool processing jobs with timeout handling.

```bash
go run main.go
```
