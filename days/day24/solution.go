package main

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func solve(A []string, s string) []int {
	n := len(s)
	dp := make([]int, n)

	for _, t := range A {
		m := len(t)
		if m <= n && s[0:m] == t {
			dp[m - 1] = 1
		}
	}

	for i := 0; i < n; i++ {
		if dp[i] == 0 {
			continue
		}
		for _, t := range A {
			m := len(t)
			if i + m + 1 <= n && s[i + 1:i + m + 1] == t {
				dp[i + m] += dp[i]
			}
		}
	}
	
	return dp
}

func solvePart1(A []string, B []string) int {
	count := 0

	for _, s := range B {
		dp := solve(A, s)
		if dp[len(s) - 1] != 0 {
			count++
		}
	}

	return count
}

func solvePart2(A []string, B []string) int {
	count := 0

	for _, s := range B {
		dp := solve(A, s)
		count += dp[len(s) - 1]
	}

	return count
}

func readInput(fileName string) ([]string, []string, error) {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		return nil, nil, errors.New("could not determine the current file")
	}
	
	dir := filepath.Dir(currentFile)
	filePath := filepath.Join(dir, fileName)

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, nil, err
	}

	var list1, list2 []string
	lines := strings.Split(string(data), "\n")

	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if i == 0 {
			list1 = strings.Split(line, ", ")
		} else {
			list2 = append(list2, line)
		}
	}

	return list1, list2, nil
}
