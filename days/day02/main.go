package main

import "fmt"

func main() {
	A, err := readInput("input.txt")

	if err != nil {
		fmt.Printf("Error when reading input: %v\n", err)
		return
	}

	safe := solvePart1(A)

	fmt.Printf("Number of safe reports: %d\n", safe)

	almostSafe := solvePart2(A)

	fmt.Printf("Number of almost safe reports: %d\n", almostSafe)
}
