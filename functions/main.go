package main

import "fmt"

// same as add(a int, b int)
func add(a, b int) int {
	return a + b
}

// the way that a function in go returns multiple values
func addAndMultiply(a, b int) (int, int) {
	return a + b, a * b
}

func main() {
	addResult := add(2, 3)
	fmt.Printf("add: %v\n", addResult)

	// why didn't go complain about below add variable when there is add function?
	add, multiple := addAndMultiply(2, 3)

	fmt.Printf("add: %v, multiply: %v\n", add, multiple)
}
