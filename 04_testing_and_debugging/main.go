package main

import (
	"errors"
	"fmt"
)

// Calculator performs basic arithmetic
type Calculator struct{}

// Add adds two integers
func (c Calculator) Add(a, b int) int {
	return a + b
}

// Divide divides a by b, returning error on division by zero
func (c Calculator) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("=== 04 Testing & Debugging ===")
	fmt.Println("Run 'go test -v' to see the tests in action!")

	calc := Calculator{}
	sum := calc.Add(5, 3)
	fmt.Printf("5 + 3 = %d\n", sum)
}
