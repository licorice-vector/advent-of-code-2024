package main

import (
	"errors"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func permute(s string) []string {
    if len(s) == 0 {
        return []string{""}
    }
    var result []string
    for i, c := range s {
        rest := s[:i] + s[i + 1:]
        for _, p := range permute(rest) {
            result = append(result, string(c) + p)
        }
    }
    return result
}

func valid(p string, pos, nogo []int) bool {
	y, x := pos[0], pos[1]
	for _, move := range p {
		switch move {
			case '^':
				y--
			case 'v':
				y++
			case '<':
				x--
			case '>':
				x++
		}
		if y == nogo[0] && x == nogo[1] {
			return false
		}
	}
	return true
}

func getPaths(a, b rune, useKeypad bool) []string {
	numpad := map[rune][]int{
		'7': {0, 0},
		'8': {0, 1},
		'9': {0, 2},
		'4': {1, 0},
		'5': {1, 1},
		'6': {1, 2},
		'1': {2, 0},
		'2': {2, 1},
		'3': {2, 2},
		'0': {3, 1},
		'A': {3, 2},
	}
	keypad := map[rune][]int{
		'^': {0, 1},
		'A': {0, 2},
		'<': {1, 0},
		'v': {1, 1},
		'>': {1, 2},
	}
	var aPos, bPos, nogo []int
	if useKeypad {
		aPos, bPos = keypad[a], keypad[b]
		nogo = []int{0, 0}
	} else {
		aPos, bPos = numpad[a], numpad[b]
		nogo = []int{3, 0}
	}
	dy, dx := bPos[0] - aPos[0], bPos[1] - aPos[1]
	path := ""
	if dy < 0 {
		path += strings.Repeat("^", -dy)
	} else if dy > 0 {
		path += strings.Repeat("v", dy)
	}
	if dx < 0 {
		path += strings.Repeat("<", -dx)
	} else if dx > 0 {
		path += strings.Repeat(">", dx)
	}
	paths := permute(path)
	var validPaths []string
	for _, p := range paths {
		if valid(p, aPos, nogo) {
			validPaths = append(validPaths, p + "A")
		}
	}
	return validPaths
}

func getCost(dp map[rune]map[rune]map[int]int, a, b rune, useKeypad bool, depth int) int {
	/* How much does it cost to move from a to b and press b? */
	if depth == 0 {
		return 1
	}
	if dp[a] == nil {
        dp[a] = make(map[rune]map[int]int)
    }
	if dp[a][b] == nil {
		dp[a][b] = make(map[int]int)
	}
	if cost, exists := dp[a][b][depth]; exists {
		return cost
	}
	paths := getPaths(a, b, useKeypad)
	minCost := math.MaxInt64
	for _, path := range paths {
		path = "A" + path
		cost := 0
		for i := 1; i < len(path); i++ {
			c, d := rune(path[i - 1]), rune(path[i])
			cost += getCost(dp, c, d, true, depth - 1)
		}
		if cost < minCost {
			minCost = cost
		}
	}
	dp[a][b][depth] = minCost
	return minCost
}

func solve(buttons string, depth int) int {
	dp := map[rune]map[rune]map[int]int{}
	buttons = "A" + buttons
	cost := 0
	for i := 1; i < len(buttons); i++ {
		a, b := rune(buttons[i - 1]), rune(buttons[i])
		cost += getCost(dp, a, b, false, depth)
	}
	return cost
}

func numerical(s string) int {
	i := 0
	for s[i] == '0' {
		i++
	}
	value, _ := strconv.Atoi(string(s[i:len(s) - 1]))
	return value
}

func solvePart1(A []string) int {
	result := 0
	for _, buttons := range A {
		length := solve(buttons, 3)
		result += length * numerical(buttons)
	}
	return result
}

func solvePart2(A []string) int {
	result := 0
	for _, buttons := range A {
		length := solve(buttons, 26)
		result += length * numerical(buttons)
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

	var list []string
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		list = append(list, line)
	}

	return list, nil
}
