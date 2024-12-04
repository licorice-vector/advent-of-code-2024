package main

import (
	"testing"
)

func TestExamplePart1(t *testing.T) {
	A, err := readInput("example.txt")

	if err != nil {
		t.Fatalf("Error when reading input: %v\n", err)
	}

	count := solvePart1(A)

	if count != 18 {
		t.Fatalf("Expected count == 18 but got count == %d\n", count)
	}
}

func TestExamplePart2(t *testing.T) {
	A, err := readInput("example.txt")

	if err != nil {
		t.Fatalf("Error when reading input: %v\n", err)
	}

	count := solvePart2(A)

	if count != 9 {
		t.Fatalf("Expected count == 9 but got count == %d\n", count)
	}
}