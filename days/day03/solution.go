package main

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
)

func solvePart1(input string) int {
	pattern := `mul\((\d+),(\d+)\)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(input, -1)

	result := 0

	for _, match := range matches {
		if len(match) == 3 {
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			result += x * y
		}
	}

	return result
}

func solvePart2(input string) int {
	pattern := `mul\((\d+),(\d+)\)|do\(\)|don't\(\)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(input, -1)

	result := 0
	do := true

	for _, match := range matches {
		if len(match) == 0 {
			continue
		}
		if match[0] == "do()" {
			do = true
		} else if match[0] == "don't()" {
			do = false
		} else if len(match) == 3 && do {
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			result += x * y
		}
	}

	return result
}

func readInput(fileName string) (string, error) {
	_, currentFile, _, ok := runtime.Caller(0)
	
	if !ok {
		return "", errors.New("could not determine the current file")
	}
	
	dir := filepath.Dir(currentFile)
	filePath := filepath.Join(dir, fileName)

	data, err := os.ReadFile(filePath)

	if err != nil {
		return "", err
	}

	return string(data), nil
}
