package main

import "fmt"

func main() {
	constraints, updates, err := readInput("input.txt")

	if err != nil {
		fmt.Printf("Error when reading input: %v\n", err)
		return
	}

	resultPart1 := solvePart1(constraints, updates)

	fmt.Printf("Result for part 1: %d\n", resultPart1)

	resultPart2, err := solvePart2(constraints, updates)

	if err != nil {
		fmt.Printf("Error in part 2: %v\n", err)
		return
	}

	fmt.Printf("Result for part 2: %d\n", resultPart2)
}
