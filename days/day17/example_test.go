package main

import (
	"slices"
	"testing"
)

func TestExamplePart1(t *testing.T) {
	A, err := readInput("example.txt")

	if err != nil {
		t.Fatalf("Error when reading input: %v\n", err)
	}

	result := solvePart1(A)

	if !slices.Equal(result, []int{5, 7, 3, 0}) {
		t.Fatalf("Expected result == [5 7 3 0] but got result == %v\n", result)
	}
}
