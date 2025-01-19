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

	if result != 37327623 {
		t.Fatalf("Expected result == 37327623 but got result == %d\n", result)
	}
}

func TestExamplePart2(t *testing.T) {
	A, err := readInput("example.txt")

	if err != nil {
		t.Fatalf("Error when reading input: %v\n", err)
	}

	result := solvePart2(A)

	if result != 24 {
		t.Fatalf("Expected result == 24 but got result == %d\n", result)
	}
}