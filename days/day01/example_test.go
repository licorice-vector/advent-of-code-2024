package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	A, B, err := readInput("example.txt")

	if err != nil {
		t.Fatalf("Error when reading input: %v\n", err)
	}

	dist, err := solve(A, B)

	if err != nil {
		t.Fatalf("Error when solving problem: %v\n", err)
	}

	if dist != 11 {
		t.Fatalf("Expected dist == 11 but got dist == %d\n", dist)
	}
}
