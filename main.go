package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	c := sum(5, 10)
	fmt.Println(c, "gh-events & gh-filters demo")

	// Use an external package
	id := uuid.New()
	fmt.Println("Generated UUID:", id)
}

func sum(a, b int) int {
	return a + b
}
