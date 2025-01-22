package main

import "fmt"

func main() {
	A, B, C, err := readInput("input.txt")

	if err != nil {
		fmt.Printf("Error when reading input: %v\n", err)
		return
	}

	result := solvePart1(A, B, C)

	fmt.Printf("Result: %d\n", result)

	result2 := solvePart2(A, B, C)

	fmt.Printf("Result: %v\n", result2)
}
