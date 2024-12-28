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

func solvePart1(m int, n int, A [][]int) int {
	for i := 0; i < 100; i++ {
		for _, robot := range A {
			x, y, vX, vY := robot[0], robot[1], robot[2], robot[3]
			x += vX
			y += vY
			x = (x + m) % m
			y = (y + n) % n
			robot[0], robot[1] = x, y
		}
	}
	var Q = make([]int, 4)
	for _, robot := range A {
		x, y := robot[0], robot[1]
		if x < m / 2 && y < n / 2 { // Q1
			Q[0]++
		} else if x > m / 2 && y < n / 2 { // Q2
			Q[1]++
		} else if x < m / 2 && y > n / 2 { // Q3
			Q[2]++
		} else if x > m / 2 && y > n / 2 { // Q4
			Q[3]++
		}
	}
	var result = 1
	for _, count := range Q {
		result *= count
	}
	return result
}

func makeGrid(m int, n int, A [][]int) [][]string {
	grid := make([][]string, n)
	for i := range grid {
		grid[i] = make([]string, m)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}
	for _, robot := range A {
		x, y := robot[0], robot[1]
		grid[y][x] = "R"
	}
	return grid
}

func printRobots(m int, n int, A [][]int) {
	grid := makeGrid(m, n, A)
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < m; j++ {
			fmt.Print(grid[i][j])
		}
		fmt.Println()
	}
}

func is5x5(m int, n int, A [][]int) bool {
	size := 5
	grid := makeGrid(m, n, A)
	for i := 0; i < n - size + 1; i++ {
		for j := 0; j < m - size + 1; j++ {
			allOccupied := true
			for dx := 0; dx < size; dx++ {
				for dy := 0; dy < size; dy++ {
					if grid[i+dx][j+dy] != "R" {
						allOccupied = false
						break
					}
				}
				if !allOccupied {
					break
				}
			}
			if allOccupied {
				return true
			}
		}
	}
	return false
}

func getRobotState(A [][]int) string {
	var state []string
	for _, robot := range A {
		state = append(state, fmt.Sprintf("%d,%d", robot[0], robot[1]))
	}
	return strings.Join(state, ";")
}

func solvePart2(m int, n int, A [][]int) {
	seen := make(map[string]int)

	var i = 0
	for {
		i++
		state := getRobotState(A)
		if prevIteration, exists := seen[state]; exists {
			fmt.Printf("Pattern repeats after %d iterations\n", i - prevIteration)
			return
		}

		seen[state] = i

		for _, robot := range A {
			x, y, vX, vY := robot[0], robot[1], robot[2], robot[3]
			x += vX
			y += vY
			x = (x + m) % m
			y = (y + n) % n
			robot[0], robot[1] = x, y
		}
		if is5x5(m, n, A) {
			printRobots(m, n, A)
			fmt.Printf("Elapsed seconds: %d\n", i)
		}
	}
}

func readInput(fileName string) (int, int, [][]int, error) {
	_, currentFile, _, ok := runtime.Caller(0)
	
	if !ok {
		return 0, 0, nil, errors.New("could not determine the current file")
	}
	
	dir := filepath.Dir(currentFile)
	filePath := filepath.Join(dir, fileName)

	data, err := os.ReadFile(filePath)

	if err != nil {
		return 0, 0, nil, err
	}

	var m, n int
	var positions [][]int

	lines := strings.Split(string(data), "\n")

	if len(lines) > 0 {
		firstLine := strings.Fields(lines[0])
		if len(firstLine) == 2 {
			m, _ = strconv.Atoi(firstLine[0])
			n, _ = strconv.Atoi(firstLine[1])
		}
	}

	for _, line := range lines[1:] {
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) >= 2 {
			p := strings.Split(parts[0][2:], ",")
			v := strings.Split(parts[1][2:], ",")

			if len(p) == 2 && len(v) == 2 {
				px, _ := strconv.Atoi(p[0])
				py, _ := strconv.Atoi(p[1])
				vx, _ := strconv.Atoi(v[0])
				vy, _ := strconv.Atoi(v[1])

				positions = append(positions, []int{px, py, vx, vy})
			}
		}
	}

	return m, n, positions, nil
}
