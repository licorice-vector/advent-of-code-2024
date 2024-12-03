package main

import "fmt"

func main() {
	input, err := readInput("input.txt")

	if err != nil {
		fmt.Printf("Error when reading input: %v\n", err)
		return
	}

	resultPart1 := solvePart1(input)

	fmt.Printf("Result for part 1: %d\n", resultPart1)

	resultPart2 := solvePart2(input)

	fmt.Printf("Result for part 2: %d\n", resultPart2)
}
