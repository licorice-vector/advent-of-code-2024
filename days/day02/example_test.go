package main

import (
	"testing"
)

func TestExamplePart1(t *testing.T) {
	A, err := readInput("example.txt")

	if err != nil {
		t.Fatalf("Error when reading input: %v\n", err)
	}

	safe := solvePart1(A)

	if safe != 2 {
		t.Fatalf("Expected safe == 2 but got safe == %d\n", safe)
	}
}

func TestExamplePart2(t *testing.T) {
	A, err := readInput("example.txt")

	if err != nil {
		t.Fatalf("Error when reading input: %v\n", err)
	}

	safe := solvePart2(A)

	if safe != 4 {
		t.Fatalf("Expected safe == 4 but got safe == %d\n", safe)
	}
}