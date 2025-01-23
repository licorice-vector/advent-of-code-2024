package main

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func solvePart1(A [][]string) int {
	locks := [][]int{}
	keys := [][]int{}
	for _, a := range A {
		lockOrKey := []int{}
		for j := 0; j < 5; j++ {
			sum := 0
			for i := 0; i < 7; i++ {
				if a[i][j] == '#' {
					sum++
				}
			}
			lockOrKey = append(lockOrKey, sum - 1)
		}
		if a[0][0] == '#' {
			locks = append(locks, lockOrKey)
		} else {
			keys = append(keys, lockOrKey)
		}
	}
	cnt := 0
	for _, lock := range locks {
		for _, key := range keys {
			fail := false
			for i := 0; i < 5; i++ {
				if lock[i] + key[i] > 5 {
					fail = true
				}
			}
			if !fail {
				cnt++
			}
		}
	}
	return cnt
}

func readInput(fileName string) ([][]string, error) {
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

	var result [][]string
	maps := strings.Split(string(data), "\n\n")

	for _, m := range maps {
		var mapData []string
		lines := strings.Split(m, "\n")

		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			runes := []rune(line)
			mapData = append(mapData, string(runes))
		}

		result = append(result, mapData)
	}

	return result, nil
}
