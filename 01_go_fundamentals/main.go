package main

import (
	"fmt"
	"log"
)

// User struct demonstrates composite types
type User struct {
	ID       int
	Username string
	Active   bool
}

// Interface demonstration
type Greeter interface {
	Greet() string
}

func (u User) Greet() string {
	return fmt.Sprintf("Hi, I'm %s (ID: %d)", u.Username, u.ID)
}

func main() {
	fmt.Println("=== 01 Go Fundamentals ===")

	// 1. Variables & Zero Values
	var count int
	fmt.Printf("Zero Value of int: %d\n", count)

	// 2. Short Variable Declaration
	message := "Hello, Go!"
	fmt.Println(message)

	// 3. Pointers
	ptr := &message
	fmt.Printf("Pointer address: %v, Value: %s\n", ptr, *ptr)

	// 4. Structs & Methods
	admin := User{
		ID:       1,
		Username: "admin_user",
		Active:   true,
	}

	// 5. Interface Usage
	sayHello(admin)

	// 6. Error Handling
	res, err := divide(10, 0)
	if err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %d\n", res)
	}
}

func sayHello(g Greeter) {
	fmt.Println(g.Greet())
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
