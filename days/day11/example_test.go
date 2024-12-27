package main

import (
	"testing"
)

func TestExamplePart1(t *testing.T) {
	A, err := readInput("example.txt")

	if err != nil {
		t.Fatalf("Error when reading input: %v\n", err)
	}

	result := solvePart1(A)

	if result != 55312 {
		t.Fatalf("Expected result == 55312 but got result == %d\n", result)
	}
}
