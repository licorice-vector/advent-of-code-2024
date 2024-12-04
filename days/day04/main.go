package main

import "fmt"

func main() {
	A, err := readInput("input.txt")

	if err != nil {
		fmt.Printf("Error when reading input: %v\n", err)
		return
	}

	countPart1 := solvePart1(A)

	fmt.Printf("Count for part 1: %d\n", countPart1)

	countPart2 := solvePart2(A)

	fmt.Printf("Count for part 2: %d\n", countPart2)
}
