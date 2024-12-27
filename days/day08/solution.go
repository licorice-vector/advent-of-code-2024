package main

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func solvePart1(A []string) int {
	n := len(A)
	m := len(A[0])
	B := make([][]int, n)
	for i := range B {
		B[i] = make([]int, m)
	}
	for i1 := 0; i1 < n; i1++ {
		for j1 := 0; j1 < m; j1++ {
			for i2 := i1; i2 < n; i2++ {
				for j2 := 0; j2 < m; j2++ {
					if i1 == i2 && j2 <= j1 {
						continue
					}
					a, b := A[i1][j1], A[i2][j2]
					if a == '.' || a != b {
						continue
					}
					dx, dy := j1 - j2, i1 - i2
					i3, j3 := i1 + dy, j1 + dx
					if !(i3 < 0 || i3 >= n || j3 < 0 || j3 >= m) {
						B[i3][j3] = 1
					}
					i3, j3 = i2 - dy, j2 - dx
					if !(i3 < 0 || i3 >= n || j3 < 0 || j3 >= m) {
						B[i3][j3] = 1
					}
				}
			}
		}
	}
	var count = 0
	for i := range B {
		for j := range B[0] {
			count += B[i][j]
		}
	}
	return count
}

func solvePart2(A []string) int {
	n := len(A)
	m := len(A[0])
	B := make([][]int, n)
	for i := range B {
		B[i] = make([]int, m)
	}
	for i1 := 0; i1 < n; i1++ {
		for j1 := 0; j1 < m; j1++ {
			for i2 := i1; i2 < n; i2++ {
				for j2 := 0; j2 < m; j2++ {
					if i1 == i2 && j2 <= j1 {
						continue
					}
					a, b := A[i1][j1], A[i2][j2]
					if a == '.' || a != b {
						continue
					}
					B[i1][j1] = 1
					B[i2][j2] = 1
					dx, dy := j1 - j2, i1 - i2
					i3, j3 := i1 + dy, j1 + dx
					for !(i3 < 0 || i3 >= n || j3 < 0 || j3 >= m) {
						B[i3][j3] = 1
						i3, j3 = i3 + dy, j3 + dx
					}
					i3, j3 = i2 - dy, j2 - dx
					for !(i3 < 0 || i3 >= n || j3 < 0 || j3 >= m) {
						B[i3][j3] = 1
						i3, j3 = i3 - dy, j3 - dx
					}
				}
			}
		}
	}
	var count = 0
	for i := range B {
		for j := range B[0] {
			count += B[i][j]
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

	var result []string
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			result = append(result, line)
		}
	}

	return result, nil
}
