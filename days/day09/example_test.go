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

	if result != 1928 {
		t.Fatalf("Expected result == 1928 but got result == %d\n", result)
	}
}

func TestExamplePart2(t *testing.T) {
	A, err := readInput("example.txt")

	if err != nil {
		t.Fatalf("Error when reading input: %v\n", err)
	}

	result := solvePart2(A)

	if result != 2858 {
		t.Fatalf("Expected result == 2858 but got result == %d\n", result)
	}
}
