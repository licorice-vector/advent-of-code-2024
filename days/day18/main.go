package main

import "fmt"

func main() {
	board, err := readInput("input.txt")

	if err != nil {
		fmt.Printf("Error when reading input: %v\n", err)
		return
	}

	result := solvePart1(board, 1024, 71)

	fmt.Printf("Result for part 1: %d\n", result)

	coordinate := solvePart2(board, 71)

	fmt.Printf("Coordinate for part 2: %v\n", coordinate)
}
