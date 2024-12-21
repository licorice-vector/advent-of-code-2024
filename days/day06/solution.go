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

func stuck(board [][]int, x int, y int, dir int, visited [][][]int) bool {
	
	if visited[y][x][dir] == 1 {
		return true
	}
	visited[y][x][dir] = 1
	
	//            U  R  D   L
	dirX := []int{0, 1, 0, -1}
	dirY := []int{-1, 0, 1, 0}

	x_next := x + dirX[dir]
	y_next := y + dirY[dir]

	if x_next < 0 || x_next >= len(board[0]) {
		return false
	}

	if y_next < 0 || y_next >= len(board) {
		return false
	}

	if board[y_next][x_next] == 1 {
		return stuck(board, x, y, (dir + 1) % 4, visited)
	}

	return stuck(board, x_next, y_next, dir, visited)
}

func solvePart2(boardStr []string) int {
	x, y := -1, -1

	for i := range boardStr {
		for j := range boardStr[i] {
			if boardStr[i][j] == '^' {
				x, y = j, i
			}
		}
	}

	var board [][]int
    for _, row := range boardStr {
        var intRow []int
        for _, char := range row {
            if char == '#' {
                intRow = append(intRow, 1)
            } else {
                intRow = append(intRow, 0)
            }
        }
        board = append(board, intRow)
    }

	count := 0
	for i := range board {
		for j := range board[i] {
			visited := make([][][]int, len(board))
			for i := range visited {
				visited[i] = make([][]int, len(board[0]))
				for j := range visited[i] {
					visited[i][j] = make([]int, 4)
				}
			}
			if board[i][j] == 0 && (i != y || j != x) {
				board[i][j] = 1
				if stuck(board, x, y, 0, visited) {
					count += 1
				}
				board[i][j] = 0
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
