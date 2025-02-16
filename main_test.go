package main

import (
	"testing"
	"github.com/google/uuid"
)

// Test cases for sum function
func Test_sum(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"Both positive", 5, 10, 15},
		{"Both negative", -5, -10, -15},
		{"Positive and negative", 5, -10, -5},
		{"Zero and positive", 0, 10, 10},
		{"Zero and negative", 0, -10, -10},
		{"Both zero", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sum(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("sum(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

// Test UUID generation
func Test_uuidGeneration(t *testing.T) {
	id := uuid.New()
	if id.String() == "" {
		t.Error("Generated UUID is empty")
	}
}
