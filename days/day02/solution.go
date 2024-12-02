package main

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func tolerate(increasing bool, a int, b int) bool {
	if a == b {
		return false
	}
	
	if increasing && b < a {
		return false
	}

	if !increasing && a < b {
		return false
	}

	difference := a - b
	if increasing {
		difference = -difference
	}

	if difference > 3 {
		return false
	}

	return true
}

func removeAt(A []int, index int) []int {
	var filtered []int

	for i := 0; i < len(A); i++ {
		if (i == index) {
			continue
		}
		filtered = append(filtered, A[i])
	}

	return filtered
}

func isAlmostSafe(report []int) bool {

	n := len(report)

	if n < 3 {
		return true
	}

	for _, increasing := range []bool{true, false} {
		failedAt := -1

		for i := 1; i < n; i++ {
			a, b := report[i - 1], report[i]

			if !tolerate(increasing, a, b) {
				failedAt = i
				break
			}
		}

		if failedAt == -1 {
			return true
		}

		left := removeAt(report, failedAt - 1)
		right := removeAt(report, failedAt)

		if isSafe(left) || isSafe(right) {
			return true
		}
	}

	return false
}

func isSafe(report []int) bool {

	if len(report) < 2 {
		return true
	}

	increasing := report[0] < report[1]

	for i := 1; i < len(report); i++ {
		a, b := report[i - 1], report[i]

		if !tolerate(increasing, a, b) {
			return false
		}
	}

	return true
}

func solvePart1(A [][]int) int {

	safe := 0

	for _, report := range A {
		if isSafe(report) {
			safe++
		}
	}

	return safe
}

func solvePart2(A [][]int) int {
	safe := 0

	for _, report := range A {
		if isAlmostSafe(report) {
			safe++
		}
	}

	return safe
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

	var A [][]int

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Fields(line)

		var slice []int

		for _, part := range parts {
			x, _ := strconv.Atoi(part)

			slice = append(slice, x)
		}

		A = append(A, slice)
	}

	return A, nil
}
