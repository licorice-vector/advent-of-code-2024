package main

import "fmt"

func main() {
	A, err := readInput("input.txt")

	if err != nil {
		fmt.Printf("Error when reading input: %v\n", err)
		return
	}

	result := solvePart1(A)

	fmt.Printf("Result: %d\n", result)

	result = solvePart2(A)

	fmt.Printf("Result: %d\n", result)
}
