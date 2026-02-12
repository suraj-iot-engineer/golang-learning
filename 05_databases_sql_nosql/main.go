package main

import (
	"database/sql"
	"fmt"
	"log"
	// In a real app, you would import the driver
	// _ "github.com/lib/pq"
)

// User represents a database record
type User struct {
	ID    int
	Name  string
	Email string
}

// UserRepository defines the interface for data access
type UserRepository interface {
	GetByID(id int) (*User, error)
	Create(user *User) error
}

// MemoryRepository simulates a database
type MemoryRepository struct {
	store map[int]*User
}

func (r *MemoryRepository) GetByID(id int) (*User, error) {
	if user, ok := r.store[id]; ok {
		return user, nil
	}
	return nil, sql.ErrNoRows
}

func (r *MemoryRepository) Create(user *User) error {
	if _, ok := r.store[user.ID]; ok {
		return fmt.Errorf("user already exists")
	}
	r.store[user.ID] = user
	return nil
}

func main() {
	fmt.Println("=== 05 Databases ===")

	// Simulate DB connection
	repo := &MemoryRepository{
		store: make(map[int]*User),
	}

	// Create User
	newUser := &User{ID: 1, Name: "Alice", Email: "alice@example.com"}
	if err := repo.Create(newUser); err != nil {
		log.Fatal(err)
	}
	fmt.Println("✅ Created User:", newUser.Name)

	// Fetch User
	u, err := repo.GetByID(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("🔍 Found User: %v\n", u)

	// Simulate "Not Found"
	_, err = repo.GetByID(99)
	if err == sql.ErrNoRows {
		fmt.Println("✅ Correctly handled non-existent user.")
	}
}
