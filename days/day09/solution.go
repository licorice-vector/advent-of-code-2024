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

func solvePart1(A []int) int {
	var B = []int{}
	for i := 0; i < len(A); i++ {
		for j := 0; j < A[i]; j++ {
			if i % 2 == 0 {
				B = append(B, i / 2)
				continue
			}
			B = append(B, -1)
		}
	}
	var i = 0
	var j = len(B) - 1
	for i < j {
		for i < j && B[i] != -1 {
			i++
		}
		if i == j {
			continue
		}
		for i < j && B[j] == -1 {
			j--
		}
		if i == j {
			continue
		}
		B[i] = B[j]
		B[j] = -1
	}
	i = 0
	var result = 0
	for i < len(B) {
		if B[i] == -1 {
			break
		}
		result += i * B[i]
		i++
	}
	for i < len(B) {
		if B[i] != -1 {
			return -1
		}
		i++
	}
	return result
}

func solvePart2(A []int) int {
	var B = []int{}
	for i := 0; i < len(A); i++ {
		for j := 0; j < A[i]; j++ {
			if i % 2 == 0 {
				B = append(B, i / 2)
				continue
			}
			B = append(B, -1)
		}
	}
	var j = len(B) - 1
	for j >= 0 {
		for j >= 0 && B[j] == -1 {
			j--
		}
		if j == -1 {
			break
		}
		var r = j
		for r - 1 >= 0 && B[r - 1] == B[j] {
			r--
		}
		for i := 0; i < j; i++ {
			if B[i] != -1 {
				continue
			}
			var l = i
			for l + 1 < j && B[l + 1] == -1 {
				l++
			}
			if l - i + 1 >= j - r + 1 {
				for k := 0; k < j - r + 1; k++ {
					B[i + k] = B[r + k]
					B[r + k] = -1
				}
				break
			}
		}
		j = r - 1
	}
	var result = 0
	for i := 0; i < len(B); i++ {
		if B[i] != -1 {
			result += i * B[i]
		}
	}
	return result
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

	var A []int

	line := strings.TrimSpace(string(data))

	for _, char := range line {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			return nil, fmt.Errorf("invalid digit: %v", err)
		}
		
		A = append(A, num)
	}

	return A, nil
}
