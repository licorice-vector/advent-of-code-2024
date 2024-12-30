package main

import (
	"testing"
)

func TestExamplePart1(t *testing.T) {
	board, err := readInput("example.txt")

	if err != nil {
		t.Fatalf("Error when reading input: %v\n", err)
	}

	result := solvePart1(board)

	if err != nil {
		t.Fatalf("Error when solving problem: %v\n", err)
	}

	if result != 7036 {
		t.Fatalf("Expected result == 7036 but got result == %d\n", result)
	}
}

func TestExamplePart2(t *testing.T) {
	board, err := readInput("example.txt")

	if err != nil {
		t.Fatalf("Error when reading input: %v\n", err)
	}

	result := solvePart2(board)

	if result != 45 {
		t.Fatalf("Expected result == 45 but got result == %d\n", result)
	}
}