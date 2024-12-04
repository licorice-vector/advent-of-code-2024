package main

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func valid(s string) bool {
	return s == "XMAS" || s == "SAMX"
}

func validX(a string, b string) bool {
	return (a == "MAS" || a == "SAM") && (b == "MAS" || b == "SAM")
}

func solvePart1(A []string) int {
	count := 0

	n := len(A)
	m := len(A[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i >= 3 && valid(string([]byte{A[i - 3][j], A[i - 2][j], A[i - 1][j], A[i][j]})) {
				count++
			}
			if j >= 3 && valid(string([]byte{A[i][j - 3], A[i][j - 2], A[i][j - 1], A[i][j]})) {
				count++
			}
			if i >= 3 && j >= 3 && valid(string([]byte{A[i - 3][j - 3], A[i - 2][j - 2], A[i - 1][j - 1], A[i][j]})) {
				count++
			}
			if i >= 3 && j >= 3 && valid(string([]byte{A[i - 3][j], A[i - 2][j - 1], A[i - 1][j - 2], A[i][j - 3]})) {
				count++
			}
		}
	}

	return count
}

func solvePart2(A []string) int {
	count := 0

	n := len(A)
	m := len(A[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i >= 2 && j >= 2 {
				a := string([]byte{A[i - 2][j - 2], A[i - 1][j - 1], A[i][j]})
				b := string([]byte{A[i - 2][j], A[i - 1][j - 1], A[i][j - 2]})
				if validX(a, b) {
					count++
				}
			}
		}
	}

	return count
}

func readInput(fileName string) ([]string, error) {
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

	return lines, nil
}
