package main

import "fmt"

func main() {
	A, B, err := readInput("input.txt")

	if err != nil {
		fmt.Printf("Error when reading input: %v\n", err)
		return
	}

	dist, err := solvePart1(A, B)

	if err != nil {
		fmt.Printf("Error when solving part 1: %v\n", err)
		return
	}

	fmt.Printf("Distance: %d\n", dist)

	similarity := solvePart2(A, B)

	fmt.Printf("Similarity: %d\n", similarity)
}
