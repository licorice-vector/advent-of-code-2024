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

func concat(a int, b int) int {
	x, err := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	if err != nil {
		panic("unexpected error in conversion")
	}
	return x
}

func valid(equation []int, x int, i int, shouldConcat bool) bool {
	if i >= len(equation) {
		return equation[0] == x
	}
	if valid(equation, x + equation[i], i + 1, shouldConcat) {
		return true
	}
	if valid(equation, x * equation[i], i + 1, shouldConcat) {
		return true
	}
	if shouldConcat && valid(equation, concat(x, equation[i]), i + 1, shouldConcat) {
		return true
	}
	return false
}

func solvePart1(A [][]int) int {
	var result = 0
	for _, equation := range A {
		if valid(equation, equation[1], 2, false) {
			result += equation[0]
		}
	}
	return result
}

func solvePart2(A [][]int) int {
	var result = 0
	for _, equation := range A {
		if valid(equation, equation[1], 2, true) {
			result += equation[0]
		}
	}
	return result
}

func readInput(fileName string) ([][]int, error) {
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

	var result [][]int

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line format: %s", line)
		}

		firstValue, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return nil, fmt.Errorf("invalid number before colon in line: %s", line)
		}

		numStrings := strings.Fields(parts[1])
		var nums []int
		for _, numStr := range numStrings {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, fmt.Errorf("invalid number in line: %s", line)
			}
			nums = append(nums, num)
		}

		result = append(result, append([]int{firstValue}, nums...))
	}

	return result, nil
}
