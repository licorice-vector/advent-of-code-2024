package main

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type QueueItem struct {
	y, x int
}

func bfs(
	A []string, 
	visited [][]int, 
	startY, 
	startX int,
	computePerimeterFunc func([]string, int, int, []int, []int, int, int) int, 
) int {
	n := len(A)
	m := len(A[0])

	//            U  R  D   L
	dirX := []int{0, 1, 0, -1}
	dirY := []int{-1, 0, 1, 0}

	area := make([][]int, n)
	for i := range area {
		area[i] = make([]int, m)
	}

	perimiter := make([][]int, n)
	for i := range perimiter {
		perimiter[i] = make([]int, m)
	}

	queue := []QueueItem{{startY, startX}}
	visited[startY][startX] = 1
	area[startY][startX] = 1
	perimiter[startY][startX] = computePerimeterFunc(A, n, m, dirY, dirX, startY, startX)

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

			if A[y_next][x_next] == A[y][x] {
				area[y_next][x_next] = 1
				perimiter[y_next][x_next] = computePerimeterFunc(A, n, m, dirY, dirX, y_next, x_next)
				if visited[y_next][x_next] == 0 {
					visited[y_next][x_next] = 1
					queue = append(queue, QueueItem{y_next, x_next})
				}
			}
		}
	}

	a, p := 0, 0
	for i := range A {
		for j := range A[i] {
			a += area[i][j]
			p += perimiter[i][j]
		}
	}
	return a * p
}

func solvePart1(A []string) int {
	n := len(A)
	m := len(A[0])
	visited := make([][]int, n)
	for i := range visited {
		visited[i] = make([]int, m)
	}
	var result = 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if visited[i][j] == 1 {
				continue
			}
			result += bfs(
				A, 
				visited, 
				i, 
				j,
				func(A []string, n int, m int, dirY []int, dirX []int, y int, x int) int {
					var perimiter = 0
					for dir := 0; dir < 4; dir++ {
						x_next := x + dirX[dir]
						y_next := y + dirY[dir]
						if x_next < 0 || x_next >= m || y_next < 0 || y_next >= n {
							perimiter++
							continue
						}
						if A[y][x] != A[y_next][x_next] {
							perimiter++
						}
					}
					return perimiter
				},
			)
		}
	}
	return result
}

func corner(x int, adj [8]int) int {
	u, r, d, l := adj[0], adj[1], adj[2], adj[3]
	ul, dr, ur, dl := adj[4], adj[5], adj[6], adj[7]
	var count = 0
	if x != u && x != l {
		count++
	}
	if x != u && x != r {
		count++
	}
	if x != d && x != l {
		count++
	}
	if x != d && x != r {
		count++
	}
	if x == u && x == l && x != ul {
		count++
	}
	if x == u && x == r && x != ur {
		count++
	}
	if x == d && x == l && x != dl {
		count++
	}
	if x == d && x == r && x != dr {
		count++
	}
	return count
}

func solvePart2(A []string) int {
	n := len(A)
	m := len(A[0])
	visited := make([][]int, n)
	for i := range visited {
		visited[i] = make([]int, m)
	}
	var result = 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if visited[i][j] == 1 {
				continue
			}
			result += bfs(
				A, 
				visited, 
				i, 
				j,
				func(A []string, n int, m int, dirY []int, dirX []int, y int, x int) int {
					var perimiter = 0
					adj := [8]int{-1, -1, -1, -1, -1, -1, -1, -1}
					for dir := 0; dir < 4; dir++ {
						ny, nx := y + dirY[dir], x + dirX[dir]
						if ny >= 0 && ny < n && nx >= 0 && nx < m {
							adj[dir] = int(A[ny][nx] - '0')
						}
					}
					//          UL DR  UR DL
					dX := []int{-1, 1, 1, -1}
					dY := []int{-1, 1, -1, 1}
					for dir := 0; dir < 4; dir++ {
						ny, nx := y + dY[dir], x + dX[dir]
						if ny >= 0 && ny < n && nx >= 0 && nx < m {
							adj[4 + dir] = int(A[ny][nx] - '0')
						}
					}
					perimiter += corner(int(A[y][x] - '0'), adj)
					return perimiter
				},
			)
		}
	}
	return result
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