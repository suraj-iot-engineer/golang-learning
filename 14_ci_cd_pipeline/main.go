package main

import "fmt"

func main() {
	fmt.Println("=== 14 CI/CD Pipeline ===")
	fmt.Println("This code is tested automatically by GitHub Actions!")
}

func Add(a, b int) int {
	return a + b
}
