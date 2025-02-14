package main

import "testing"

func Test_main(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		// TODO: Add test cases.
		{"Both positive", 5, 10, 15},
		{"Both negative", -5, -10, -15},
		{"Positive and negative", 5, -10, -5},
		{"Zero and positive", 0, 10, 10},
		{"Zero and negative", 0, -10, -10},
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
