package main

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func bfs(A []string, startY, startX int) [][]int {
	n, m := len(A), len(A[0])
	const inf = int(1e9)
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, m)
		for j := range dist[i] {
			dist[i][j] = inf
		}
	}
	queue := [][2]int{{startY, startX}}
	dist[startY][startX] = 0
	
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // up, down, left, right
	
	for len(queue) > 0 {
		y, x := queue[0][0], queue[0][1]
		queue = queue[1:]
		
		for _, d := range dirs {
			ny, nx := y + d[0], x + d[1]
			if ny >= 0 && ny < n && nx >= 0 && nx < m && A[ny][nx] != '#' && dist[ny][nx] == inf {
				dist[ny][nx] = dist[y][x] + 1
				queue = append(queue, [2]int{ny, nx})
			}
		}
	}
	return dist
}

func solvePart1(A []string, threshold int) int {
	n, m := len(A), len(A[0])
	var srcX, srcY, dstX, dstY int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if A[i][j] == 'S' {
				srcY, srcX = i, j
			} else if A[i][j] == 'E' {
				dstY, dstX = i, j
			}
		}
	}
	
	distFromSrc := bfs(A, srcY, srcX)
	distToDst := bfs(A, dstY, dstX)
	dist := distFromSrc[dstY][dstX]
	count := 0

	for i := 1; i + 1 < n; i++ {
		for j := 1; j + 1 < m; j++ {
			if A[i][j] != '#' {
				continue
			}
			if A[i - 1][j] != '#' && A[i + 1][j] != '#' {
				if dist - (distFromSrc[i - 1][j] + 1 + distToDst[i + 1][j]) >= threshold {
					count++
				}
				if dist - (distFromSrc[i + 1][j] + 1 + distToDst[i - 1][j]) >= threshold {
					count++
				}
			}
			if A[i][j - 1] != '#' && A[i][j + 1] != '#' {
				if dist - (distFromSrc[i][j - 1] + 1 + distToDst[i][j + 1]) >= threshold {
					count++
				}
				if dist - (distFromSrc[i][j + 1] + 1 + distToDst[i][j - 1]) >= threshold {
					count++
				}
			}
		}
	}

	return count
}

func solvePart2(A []string, threshold int) int {
	n, m := len(A), len(A[0])
	var srcX, srcY, dstX, dstY int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if A[i][j] == 'S' {
				srcY, srcX = i, j
			} else if A[i][j] == 'E' {
				dstY, dstX = i, j
			}
		}
	}
	
	distFromSrc := bfs(A, srcY, srcX)
	distToDst := bfs(A, dstY, dstX)
	dist := distFromSrc[dstY][dstX]
	count := 0

	for y1 := 1; y1 + 1 < n; y1++ {
		for x1 := 1; x1 + 1 < m; x1++ {
			for y2 := 1; y2 + 1 < n; y2++ {
				for x2 := 1; x2 + 1 < m; x2++ {
					time := abs(x1 - x2) + abs(y1 - y2)
					if time <= 20 {
						if dist - (distFromSrc[y1][x1] + time + distToDst[y2][x2]) >= threshold {
							count++
						}
					}
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

	var result []string
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			result = append(result, line)
		}
	}

	return result, nil
}
