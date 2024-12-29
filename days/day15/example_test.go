package main

import (
	"testing"
)

func TestExamplePart1(t *testing.T) {
	A, s, err := readInput("example.txt")

	if err != nil {
		t.Fatalf("Error when reading input: %v\n", err)
	}

	result := solvePart1(A, s)

	if result != 10092 {
		t.Fatalf("Expected result == 10092 but got result == %d\n", result)
	}
}

func TestExamplePart2(t *testing.T) {
	A, s, err := readInput("example.txt")

	if err != nil {
		t.Fatalf("Error when reading input: %v\n", err)
	}

	result := solvePart2(A, s)

	if result != 9021 {
		t.Fatalf("Expected result == 9021 but got result == %d\n", result)
	}
}
