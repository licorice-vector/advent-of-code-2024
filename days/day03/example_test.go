package main

import (
	"testing"
)

func TestExamplePart1(t *testing.T) {
	input, err := readInput("example.txt")

	if err != nil {
		t.Fatalf("Error when reading input: %v\n", err)
	}

	result := solvePart1(input)

	if result != 161 {
		t.Fatalf("Expected result == 161 but got result == %d\n", result)
	}
}

func TestExamplePart2(t *testing.T) {
	input, err := readInput("example.txt")

	if err != nil {
		t.Fatalf("Error when reading input: %v\n", err)
	}

	result := solvePart2(input)

	if result != 48 {
		t.Fatalf("Expected result == 48 but got result == %d\n", result)
	}
}