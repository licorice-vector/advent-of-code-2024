package main

import (
	"testing"
)

func TestExamplePart1(t *testing.T) {
	A, err := readInput("example.txt")

	if err != nil {
		t.Fatalf("Error when reading input: %v\n", err)
	}

	result := solvePart1(A, 20)

	if result != 5 {
		t.Fatalf("Expected result == 5 but got result == %d\n", result)
	}
}

func TestExamplePart2(t *testing.T) {
	A, err := readInput("example.txt")

	if err != nil {
		t.Fatalf("Error when reading input: %v\n", err)
	}

	result := solvePart2(A, 50)

	if result != 285 {
		t.Fatalf("Expected result == 285 but got result == %d\n", result)
	}
}