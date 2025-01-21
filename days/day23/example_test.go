package main

import (
	"slices"
	"testing"
)

func TestExamplePart1(t *testing.T) {
	A, err := readInput("example.txt")

	if err != nil {
		t.Fatalf("Error when reading input: %v\n", err)
	}

	result := solvePart1(A)

	if result != 7 {
		t.Fatalf("Expected result == 7 but got result == %d\n", result)
	}
}

func TestExamplePart2(t *testing.T) {
	A, err := readInput("example.txt")

	if err != nil {
		t.Fatalf("Error when reading input: %v\n", err)
	}

	result := solvePart2(A)

	if !slices.Equal(result, []string{"co", "de", "ka", "ta"}) {
		t.Fatalf("Expected result == [\"co\", \"de\", \"ka\", \"ta\"] but got result == %v\n", result)
	}
}
