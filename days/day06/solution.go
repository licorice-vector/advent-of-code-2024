package main

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func walk(board []string, x int, y int, dir int, visited [][]int) {
	visited[y][x] = 1
	//            U  R  D   L
	dirX := []int{0, 1, 0, -1}
	dirY := []int{-1, 0, 1, 0}

	x_next := x + dirX[dir]
	y_next := y + dirY[dir]

	if x_next < 0 || x_next >= len(board[0]) {
		return
	}

	if y_next < 0 || y_next >= len(board) {
		return
	}

	if board[y_next][x_next] == '#' {
		walk(board, x, y, (dir + 1) % 4, visited)
		return
	}

	walk(board, x_next, y_next, dir, visited)
}

func solvePart1(board []string) int {
	x, y := -1, -1

	for i := range board {
		for j := range board[i] {
			if board[i][j] == '^' {
				x, y = j, i
			}
		}
	}

	visited := make([][]int, len(board))
	for i := range visited {
		visited[i] = make([]int, len(board[0]))
	}

	walk(board, x, y, 0, visited)

	count := 0
	for i := range visited {
		for j := range visited[i] {
			count += visited[i][j]
		}
	}

	return count
}

func solvePart2(board []string) int {
	x, y := -1, -1

	for i := range board {
		for j := range board[i] {
			if board[i][j] == '^' {
				x, y = j, i
			}
		}
	}

	visited := make([][]int, len(board))
	for i := range visited {
		visited[i] = make([]int, len(board[0]))
	}

	walk(board, x, y, 0, visited)

	count := 0
	for i := range visited {
		for j := range visited[i] {
			count += visited[i][j]
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

	var board []string = []string{}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		board = append(board, line)
	}

	return board, nil
}
