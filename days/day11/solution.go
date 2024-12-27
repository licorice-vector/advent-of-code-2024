package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func updateStones(A map[int]int) map[int]int {
	B := make(map[int]int)
	for stone, count := range A {
		if stone == 0 {
			B[1] += count
			continue
		}
		s := strconv.Itoa(stone)
		if len(s) % 2 == 0 {
			mid := len(s) / 2
			l, _ := strconv.Atoi(s[:mid])
			r, _ := strconv.Atoi(s[mid:])
			B[l] += count
			B[r] += count
		} else {
			B[stone * 2024] += count
		}
	}
	return B
}

func createFrequencyMap(stones []int) map[int]int {
	freq := make(map[int]int)
	for _, stone := range stones {
		freq[stone]++
	}
	return freq
}

func countStones(freq map[int]int) int {
	total := 0
	for _, count := range freq {
		total += count
	}
	return total
}

func solvePart1(A []int) int {
	B := createFrequencyMap(A)
	for i := 0; i < 25; i++ {
		B = updateStones(B)
	}
	return countStones(B)
}

func solvePart2(A []int) int {
	B := createFrequencyMap(A)
	for i := 0; i < 75; i++ {
		B = updateStones(B)
	}
	return countStones(B)
}

func readInput(fileName string) ([]int, error) {
	_, currentFile, _, ok := runtime.Caller(0)

	if !ok {
		return nil, errors.New("could not determine the current file")
	}

	dir := filepath.Dir(currentFile)
	filePath := filepath.Join(dir, fileName)

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")

	var result []int

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		parts := strings.Fields(line)
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("invalid number in input: %v", err)
			}
			result = append(result, num)
		}
	}

	return result, nil
}
