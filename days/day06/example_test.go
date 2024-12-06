package main

import (
	"testing"
)

func TestExamplePart1(t *testing.T) {
	board, err := readInput("example.txt")

	if err != nil {
		t.Fatalf("Error when reading input: %v\n", err)
	}

	count := solvePart1(board)

	if err != nil {
		t.Fatalf("Error when solving problem: %v\n", err)
	}

	if count != 41 {
		t.Fatalf("Expected count == 41 but got count == %d\n", count)
	}
}

func TestExamplePart2(t *testing.T) {
	board, err := readInput("example.txt")

	if err != nil {
		t.Fatalf("Error when reading input: %v\n", err)
	}

	count := solvePart2(board)

	if count != 6 {
		t.Fatalf("Expected count == 6 but got count == %d\n", count)
	}
}