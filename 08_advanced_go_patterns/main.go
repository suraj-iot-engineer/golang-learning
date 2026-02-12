package main

import (
	"fmt"
	"time"
)

// --- 1. Generics ---

// Stack is a generic LIFO data structure
type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(val T) {
	s.elements = append(s.elements, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	index := len(s.elements) - 1
	val := s.elements[index]
	s.elements = s.elements[:index]
	return val, true
}

// --- 2. Functional Options Pattern ---

type Server struct {
	Host    string
	Port    int
	Timeout time.Duration
}

type ServerOption func(*Server)

func NewServer(opts ...ServerOption) *Server {
	// Default values
	s := &Server{
		Host:    "localhost",
		Port:    8080,
		Timeout: 30 * time.Second,
	}

	for _, opt := range opts {
		opt(s)
	}
	return s
}

func WithPort(port int) ServerOption {
	return func(s *Server) {
		s.Port = port
	}
}

func WithTimeout(d time.Duration) ServerOption {
	return func(s *Server) {
		s.Timeout = d
	}
}

func main() {
	fmt.Println("=== 08 Advanced Patterns ===")

	// Generics Usage
	intStack := Stack[int]{}
	intStack.Push(10)
	intStack.Push(20)
	val, _ := intStack.Pop()
	fmt.Printf("Generic Stack Pop: %d\n", val)

	stringStack := Stack[string]{}
	stringStack.Push("Go")
	stringStack.Push("Generics")
	strVal, _ := stringStack.Pop()
	fmt.Printf("Generic Stack Pop: %s\n", strVal)

	// Functional Options Usage
	srv := NewServer(
		WithPort(9000),
		WithTimeout(1*time.Minute),
	)
	fmt.Printf("Server Config: %+v\n", srv)
}
