package main

import "fmt"

func main() {
	A, B, err := readInput("input.txt")

	if err != nil {
		fmt.Printf("Error when reading input: %v\n", err)
		return
	}

	dist, err := solve(A, B)

	if err != nil {
		fmt.Printf("Error when solving problem: %v\n", err)
		return
	}

	fmt.Println(dist)
}
