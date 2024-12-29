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

func solvePart1(A []string, s string) int {
	x, y := -1, -1

	for i := range A {
		for j := range A[i] {
			if A[i][j] == '@' {
				x, y = j, i
			}
		}
	}

	n, m := len(A), len(A[0])

	B := make([][]int, n)
	for i := range B {
		B[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if A[i][j] == '#' {
				B[i][j] = 1
			} else if A[i][j] == 'O' {
				B[i][j] = 2
			}
		}
	}

	for _, dir := range s {
		dX, dY := 0, 0
		if dir == '^' { // up
			dX, dY = 0, -1
		} else if dir == 'v' { // down
			dX, dY = 0, 1
		} else if dir == '>' { // right
			dX, dY = 1, 0
		} else { // left
			dX, dY = -1, 0
		}
		nx, ny := x + dX, y + dY
		kx, ky := nx, ny
		for B[ky][kx] == 2 {
			kx, ky = kx + dX, ky + dY
		}
		if B[ky][kx] == 0 {
			for i := 1; i < max(abs(nx - kx), abs(ny - ky)) + 1; i++ {
				B[ny + i * dY][nx + i * dX] = 2
			}
			B[ny][nx] = 0
		}
		if B[ny][nx] == 0 {
			x, y = nx, ny
		}
	}

	result := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if B[i][j] == 2 {
				result += 100 * i + j
			}
		}
	}
	return result
}

func solvePart2(A []string, s string) int {
	n, m := len(A), len(A[0])

	AA := make([]string, n)
	for i := range AA {
		for j := 0; j < m; j++ {
			if A[i][j] == '#' {
				AA[i] += "##"
			} else if A[i][j] == 'O' {
				AA[i] += "[]"
			} else if A[i][j] == '.' {
				AA[i] += ".."
			} else {
				AA[i] += "@."
			}
		}
	}

	n, m = len(AA), len(AA[0])
	x, y := -1, -1

	for i := range AA {
		for j := range AA[i] {
			if AA[i][j] == '@' {
				x, y = j, i
			}
		}
	}

	B := make([][]int, n)
	for i := range B {
		B[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if AA[i][j] == '#' {
				B[i][j] = 1
			} else if AA[i][j] == '[' {
				B[i][j] = 2
			} else if AA[i][j] == ']' {
				B[i][j] = 3
			}
		}
	}

	for _, dir := range s {
		dX, dY := 0, 0
		if dir == '^' { // up
			dX, dY = 0, -1
		} else if dir == 'v' { // down
			dX, dY = 0, 1
		} else if dir == '>' { // right
			dX, dY = 1, 0
		} else { // left
			dX, dY = -1, 0
		}
		if dX != 0 {
			nx, ny := x + dX, y + dY
			kx, ky := nx, ny
			for B[ky][kx] == 2 || B[ky][kx] == 3 {
				kx, ky = kx + dX, ky + dY
			}
			if B[ky][kx] == 0 {
				for i := 1; i < abs(nx - kx) + 1; i += 2 {
					a, b := 0, 0
					if dX == 1 {
						a, b = 2, 3
					} else {
						a, b = 3, 2
					}
					B[ny + i * dY][nx + i * dX] = a
					B[ny + i * dY + dY][nx + i * dX + dX] = b
				}
				B[ny][nx] = 0
			}
			if B[ny][nx] == 0 {
				x, y = nx, ny
			}
		} else {
			C := make([][]int, n)
			for i := range C {
				C[i] = make([]int, m)
			}
			C[y][x] = 1
			ny := y + dY
			for ny > 0 && ny + 1 < n {
				py := ny - dY
				for j := 0; j < m; j++ {
					if B[ny][j] == 1 || B[ny][j] == 0 || C[py][j] == 0 {
						continue
					}
					C[ny][j] = 1
					if B[ny][j] == 2 {
						C[ny][j + 1] = 1
					} else {
						C[ny][j - 1] = 1
					}
				}
				ny += dY
			}
			move := true
			if dY == 1 {
				for i := n - 2; i >= 0; i-- {
					for j := 0; j < m; j++ {
						if C[i][j] == 1 && B[i + 1][j] == 1 {
							move = false
						}
					}
				}
				if move {
					for i := n - 2; i >= 0; i-- {
						for j := 0; j < m; j++ {
							if C[i][j] == 1 {
								B[i + 1][j] = B[i][j]
								B[i][j] = 0
							}
						}
					}
				}
			} else {
				for i := 1; i < n; i++ {
					for j := 0; j < m; j++ {
						if C[i][j] == 1 && B[i - 1][j] == 1 {
							move = false
						}
					}
				}
				if move {
					for i := 1; i < n; i++ {
						for j := 0; j < m; j++ {
							if C[i][j] == 1 {
								B[i - 1][j] = B[i][j]
								B[i][j] = 0
							}
						}
					}
				}
			}
			if move {
				x, y = x + dX, y + dY
			}
		}
	}

	result := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if B[i][j] == 2 {
				result += 100 * i + j
			}
		}
	}
	return result
}

func readInput(fileName string) ([]string, string, error) {
	_, currentFile, _, ok := runtime.Caller(0)

	if !ok {
		return nil, "", errors.New("could not determine the current file")
	}

	dir := filepath.Dir(currentFile)
	filePath := filepath.Join(dir, fileName)

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, "", err
	}

	lines := strings.Split(string(data), "\n")

	var A []string
	var sBuilder strings.Builder
	parsingGrid := true

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" {
			continue
		}

		if parsingGrid {
			if !strings.HasPrefix(trimmedLine, "#") && !strings.ContainsAny(trimmedLine, ".O") {
				parsingGrid = false
			} else {
				A = append(A, trimmedLine)
				continue
			}
		}

		sBuilder.WriteString(trimmedLine)
	}

	s := sBuilder.String()

	return A, s, nil
}
