package main

import "fmt"

func main() {
	board, err := readInput("input.txt")

	if err != nil {
		fmt.Printf("Error when reading input: %v\n", err)
		return
	}

	count := solvePart1(board)

	fmt.Printf("Count for part 1: %d\n", count)

	count = solvePart2(board)

	fmt.Printf("Count for part 2: %d\n", count)
}
