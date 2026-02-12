package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func main() {
	// Using an external dependency
	color.Cyan("=== 02 Modules & Workspace ===")

	c := color.New(color.FgHiGreen, color.Bold)
	c.Println("Successfully imported github.com/fatih/color!")

	fmt.Println("This module demonstrates how to manage dependencies in Go.")
	fmt.Printf("Current Time: %s\n", time.Now().Format(time.RFC1123))

	color.Red("Ensure you run 'go mod tidy' if dependencies are missing.")
}
