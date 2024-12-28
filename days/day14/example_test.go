package main

import (
	"testing"
)

func TestExamplePart1(t *testing.T) {
	m, n, A, err := readInput("example.txt")

	if err != nil {
		t.Fatalf("Error when reading input: %v\n", err)
	}

	result := solvePart1(m, n, A)

	if result != 12 {
		t.Fatalf("Expected result == 12 but got result == %d\n", result)
	}
}
