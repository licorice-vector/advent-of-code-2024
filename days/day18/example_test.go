package main

import (
	"slices"
	"testing"
)

func TestExamplePart1(t *testing.T) {
	board, err := readInput("example.txt")

	if err != nil {
		t.Fatalf("Error when reading input: %v\n", err)
	}

	result := solvePart1(board, 12, 7)

	if err != nil {
		t.Fatalf("Error when solving problem: %v\n", err)
	}

	if result != 22 {
		t.Fatalf("Expected result == 22 but got result == %d\n", result)
	}
}

func TestExamplePart2(t *testing.T) {
	board, err := readInput("example.txt")

	if err != nil {
		t.Fatalf("Error when reading input: %v\n", err)
	}

	result := solvePart2(board, 7)

	if !slices.Equal(result, []int{6, 1}) {
		t.Fatalf("Expected result == [6 1] but got result == %v\n", result)
	}
}