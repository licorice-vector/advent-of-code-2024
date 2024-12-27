package main

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func computeDistance(A []int, B []int) int {
	dist, n := 0, len(A)

	for i := 0; i < n; i++ {
		dist += abs(A[i] - B[i])
	}

	return dist
}

func solvePart1(A []int, B []int) (int, error) {
	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})

	sort.Slice(B, func(i, j int) bool {
		return B[i] < B[j]
	})

	if len(A) != len(B) {
		return 0, errors.New("slices are not of equal length")
	}

	dist := computeDistance(A, B)

	return dist, nil
}

func solvePart2(A []int, B []int) int {
	freq := map[int]int{}

	for _, value := range B {
		freq[value]++
	}

	similarity := 0

	for _, value := range A {
		similarity += freq[value] * value
	}

	return similarity
}

func readInput(fileName string) ([]int, []int, error) {
	_, currentFile, _, ok := runtime.Caller(0)
	
	if !ok {
		return nil, nil, errors.New("could not determine the current file")
	}
	
	dir := filepath.Dir(currentFile)
	filePath := filepath.Join(dir, fileName)

	data, err := os.ReadFile(filePath)

	if err != nil {
		return nil, nil, err
	}

	var A, B []int

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Fields(line)

		if len(parts) == 2 {
			a, _ := strconv.Atoi(parts[0])
			b, _ := strconv.Atoi(parts[1])

			A = append(A, a)
			B = append(B, b)
		}
	}

	return A, B, nil
}
