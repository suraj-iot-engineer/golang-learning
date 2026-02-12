package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	calc := Calculator{}
	result := calc.Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

func TestDivide(t *testing.T) {
	calc := Calculator{}

	// Table Driven Tests
	tests := []struct {
		name        string
		a, b        int
		want        int
		expectError bool
	}{
		{"Normal division", 10, 2, 5, false},
		{"Divide by zero", 10, 0, 0, true},
		{"Negative division", -10, 2, -5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calc.Divide(tt.a, tt.b)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if got != tt.want {
					t.Errorf("got %d, want %d", got, tt.want)
				}
			}
		})
	}
}
