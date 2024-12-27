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

type QueueItem struct {
	y, x int
}

func bfs(A [][]int, startY, startX int) int {
	n := len(A)
	m := len(A[0])

	//            U  R  D   L
	dirX := []int{0, 1, 0, -1}
	dirY := []int{-1, 0, 1, 0}

	B := make([][]int, n)
	for i := range B {
		B[i] = make([]int, m)
	}

	queue := []QueueItem{{startY, startX}}
	B[startY][startX] = 1

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		y, x := item.y, item.x

		for dir := 0; dir < 4; dir++ {
			x_next := x + dirX[dir]
			y_next := y + dirY[dir]

			if x_next < 0 || x_next >= m || y_next < 0 || y_next >= n {
				continue
			}

			if A[y_next][x_next] == A[y][x] + 1 {
				B[y_next][x_next] += B[y][x]
				if B[y_next][x_next] == B[y][x] {
					queue = append(queue, QueueItem{y_next, x_next})
				}
			}
		}
	}

	count := 0
	for i := range B {
		for j := range B[i] {
			if A[i][j] == 9 {
				count += B[i][j]
			}
		}
	}
	return count
}

func dfs(A [][]int, B [][]int, visited [][]int, y int, x int) {
	n := len(A)
	m := len(A[0])

	if visited[y][x] == 1 {
		return
	}
	visited[y][x] = 1

	if A[y][x] == 9 {
		B[y][x] = 1
	}
	
	//            U  R  D   L
	dirX := []int{0, 1, 0, -1}
	dirY := []int{-1, 0, 1, 0}

	for dir := 0; dir < 4; dir++ {
		x_next := x + dirX[dir]
		y_next := y + dirY[dir]

		if x_next < 0 || x_next >= m || y_next < 0 || y_next >= n {
			continue
		}

		if A[y_next][x_next] == A[y][x] + 1 {
			dfs(A, B, visited, y_next, x_next)
		}
	}
}

func solvePart1(A [][]int) int {
	n := len(A)
	m := len(A[0])
	var result = 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if A[i][j] != 0 {
				continue
			}
			B := make([][]int, n)
			for i := range B {
				B[i] = make([]int, m)
			}
			visited := make([][]int, n)
			for i := range visited {
				visited[i] = make([]int, m)
			}
			dfs(A, B, visited, i, j)
			var count = 0
			for i := range B {
				for j := range B[0] {
					count += B[i][j]
					B[i][j] = 0
				}
			}
			result += count
		}
	}
	return result
}

func solvePart2(A [][]int) int {
	n := len(A)
	m := len(A[0])
	var result = 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if A[i][j] == 0 {
				result += bfs(A, i, j)
			}
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

	lines := strings.Split(string(data), "\n")

	var result [][]int

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		var currentLine []int

		for _, char := range line {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				return nil, fmt.Errorf("invalid character in input: %v", err)
			}

			currentLine = append(currentLine, num)
		}

		result = append(result, currentLine)
	}

	return result, nil
}
