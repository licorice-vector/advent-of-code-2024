package main

import (
	"testing"
)

func TestExamplePart1(t *testing.T) {
	A, B, err := readInput("example.txt")

	if err != nil {
		t.Fatalf("Error when reading input: %v\n", err)
	}

	result := solvePart1(A, B)

	if result != 6 {
		t.Fatalf("Expected result == 6 but got result == %d\n", result)
	}
}

func TestExamplePart2(t *testing.T) {
	A, B, err := readInput("example.txt")

	if err != nil {
		t.Fatalf("Error when reading input: %v\n", err)
	}

	result := solvePart2(A, B)

	if result != 16 {
		t.Fatalf("Expected result == 16 but got result == %d\n", result)
	}
}