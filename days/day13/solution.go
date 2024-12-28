package main

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func solvePart1(A [][]int) int {
	var cost = 0
	for i := range A {
		aX, aY, bX, bY, X, Y := A[i][0], A[i][1], A[i][2], A[i][3], A[i][4], A[i][5]
		for a := 0; a <= 100; a++ {
			x, y := aX * a, aY * a
			xLeft, yLeft := X - x, Y - y
			if xLeft % bX == 0 && yLeft % bY == 0 {
				b := xLeft / bX
				if yLeft / bY != b {
					continue
				}
				cost += a * 3 + b
			}
		}
	}
	return cost
}

func solveEquations(aX, aY, bX, bY, X, Y int) (int, int, bool) {
	det := aX * bY - aY * bX
    
    if det == 0 {
        return 0, 0, false
    }
    
    a := bY * X - bX * Y
    b := aX * Y - aY * X

	if a % det != 0 || b % det != 0 {
		return 0, 0, false
	}

	a /= det
	b /= det

	if a < 0 || b < 0 {
		return 0, 0, false
	}
    
    return a, b, true
}

func solvePart2(A [][]int) int {
	var cost = 0
	for i := range A {
		aX, aY, bX, bY, X, Y := A[i][0], A[i][1], A[i][2], A[i][3], A[i][4], A[i][5]
		X += 10000000000000
		Y += 10000000000000
		a, b, found := solveEquations(aX, aY, bX, bY, X, Y)
		if found {
			cost += a * 3 + b
		}
	}
	return cost
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

	var results [][]int
	lines := strings.Split(string(data), "\n")
	var block []int

	re := regexp.MustCompile(`[+-]?\d+`)

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" {
			if len(block) > 0 {
				results = append(results, block)
				block = nil
			}
			continue
		}

		matches := re.FindAllString(line, -1)
		for _, match := range matches {
			num, err := strconv.Atoi(match)
			if err == nil {
				block = append(block, num)
			}
		}
	}

	if len(block) > 0 {
		results = append(results, block)
	}

	return results, nil
}
