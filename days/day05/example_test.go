package main

import (
	"fmt"
	"testing"
)

func TestExamplePart1(t *testing.T) {
	constraints, updates, err := readInput("example.txt")

	if err != nil {
		t.Fatalf("Error when reading input: %v\n", err)
	}

	result := solvePart1(constraints, updates)

	if result != 143 {
		t.Fatalf("Expected result == 143 but got result == %d\n", result)
	}
}

func TestExamplePart2(t *testing.T) {
	constraints, updates, err := readInput("example.txt")

	if err != nil {
		t.Fatalf("Error when reading input: %v\n", err)
	}

	result, err := solvePart2(constraints, updates)

	if err != nil {
		fmt.Printf("Error in part 2: %v\n", err)
		return
	}

	if result != 123 {
		t.Fatalf("Expected result == 123 but got result == %d\n", result)
	}
}