package main

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func solvePart1(A [][]int, k int, n int) int {
	board := make([][]int, n)
	for i := 0; i < n; i++ {
		board[i] = make([]int, n)
	}

	for i := 0; i < k; i++ {
		x, y := A[i][0], A[i][1]
		board[y][x] = 1
	}

	srcX, srcY := 0, 0
	dstX, dstY := n - 1, n - 1

	directions := [][]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	type Point struct {
		x, y, dist int
	}
	queue := []Point{{srcX, srcY, 0}}
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, n)
	}
	visited[srcY][srcX] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.x == dstX && current.y == dstY {
			return current.dist
		}

		for _, dir := range directions {
			newX, newY := current.x+dir[0], current.y+dir[1]

			if newX >= 0 && newX < n && newY >= 0 && newY < n &&
				!visited[newY][newX] && board[newY][newX] == 0 {

				visited[newY][newX] = true
				queue = append(queue, Point{newX, newY, current.dist + 1})
			}
		}
	}

	return -1
}

func solvePart2(A [][]int, n int) []int {
	l, r := 1, len(A)
	for l < r {
		m := l + (r - l) / 2
		if solvePart1(A, m, n) == -1 {
			r = m
		} else {
			l = m + 1
		}
	}
	if l == len(A) {
		return []int{}
	}
	return A[l - 1]
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

	var pairs [][]int
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			return nil, errors.New("invalid line format, expected pairs of numbers")
		}
		
		num1, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
		num2, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err1 != nil || err2 != nil {
			return nil, errors.New("invalid number format")
		}

		pairs = append(pairs, []int{num1, num2})
	}

	return pairs, nil
}
