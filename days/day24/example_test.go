package main

import (
	"testing"
)

func TestExamplePart1(t *testing.T) {
	A, B, C, err := readInput("example.txt")

	if err != nil {
		t.Fatalf("Error when reading input: %v\n", err)
	}

	result := solvePart1(A, B, C)

	if result != 2024 {
		t.Fatalf("Expected result == 2024 but got result == %d\n", result)
	}
}
