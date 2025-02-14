package main

import "fmt"

func main() {
	c := sum(5, 10)
	fmt.Println(c)
}

func sum(a, b int) int {
	return a + b
}
