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

func evolve(x int) int {
	y := x * 64
	x = x ^ y
	x = x % 16777216
	y = x / 32
	x = x ^ y
	x = x % 16777216
	y = x * 2048
	x = x ^ y
	x = x % 16777216
	return x
}

func solvePart1(A []int) int {
	result := 0
	n := 2000
	for _, x := range A {
		for i := 0; i < n; i++ {
			x = evolve(x)
		}
		result += x
	}
	return result
}

func solvePart2(A []int) int {
	prices := [][]int{}
	changes := [][]int{}
	for _, x := range A {
		price := []int{}
		change := []int{}
		for i := 0; i < 2000; i++ {
			a := x % 10
			x = evolve(x)
			b := x % 10
			price = append(price, b)
			change = append(change, b - a + 9)
		}
		prices = append(prices, price)
		changes = append(changes, change)
	}
	var cost [19][19][19][19]int
	for i, change := range changes {
		var seen [19][19][19][19]int
		for j := 0; j + 3 < len(change); j++ {
			a, b, c, d := change[j], change[j + 1], change[j + 2], change[j + 3]
			if seen[a][b][c][d] == 1 {
				continue
			}
			seen[a][b][c][d] = 1
			cost[a][b][c][d] += prices[i][j + 3]
		}
	}
	max := cost[0][0][0][0]
    for i := 0; i < 19; i++ {
        for j := 0; j < 19; j++ {
            for k := 0; k < 19; k++ {
                for l := 0; l < 19; l++ {
                    if max < cost[i][j][k][l] {
                        max = cost[i][j][k][l]
                    }
                }
            }
        }
    }
	return max
}

func readInput(fileName string) ([]int, error) {
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

	var list []int
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("failed to parse integer from line: %s", line)
		}
		list = append(list, num)
	}

	return list, nil
}
