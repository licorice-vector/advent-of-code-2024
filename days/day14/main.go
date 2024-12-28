package main

import "fmt"

func deepCopy(A [][]int) [][]int {
	copyA := make([][]int, len(A))
	for i := range A {
		copyA[i] = make([]int, len(A[i]))
		copy(copyA[i], A[i])
	}
	return copyA
}

func main() {
	m, n, A, err := readInput("input.txt")

	if err != nil {
		fmt.Printf("Error when reading input: %v\n", err)
		return
	}

	result := solvePart1(m, n, deepCopy(A))

	fmt.Printf("Result: %d\n", result)

	solvePart2(m, n, A)
}
