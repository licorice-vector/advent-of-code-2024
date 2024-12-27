package main

import (
	"testing"
)

func TestExamplePart1(t *testing.T) {
	A, B, err := readInput("example.txt")

	if err != nil {
		t.Fatalf("Error when reading input: %v\n", err)
	}

	dist, err := solvePart1(A, B)

	if err != nil {
		t.Fatalf("Error when solving problem: %v\n", err)
	}

	if dist != 11 {
		t.Fatalf("Expected dist == 11 but got dist == %d\n", dist)
	}
}

func TestExamplePart2(t *testing.T) {
	A, B, err := readInput("example.txt")

	if err != nil {
		t.Fatalf("Error when reading input: %v\n", err)
	}

	similarity := solvePart2(A, B)

	if similarity != 31 {
		t.Fatalf("Expected similarity == 31 but got similarity == %d\n", similarity)
	}
}