package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strconv"
	"strings"
)

func combo(A, B, C int, x int) int {
	if x <= 3 {
		return x
	}
	if x == 4 {
		return A
	}
	if x == 5 {
		return B
	}
	if x == 6 {
		return C
	}
	fmt.Println("not supposed to happen")
	return -1
}

func run(A, B, C int, program []int) []int {
	result := make([]int, 0)
	ip := 0
    for ip < len(program) {
        opcode := program[ip]
        operand := program[ip + 1]
		comboOperand := combo(A, B, C, operand)
        switch opcode {
			case 0: // adv
				A = A / (1 << comboOperand)
			case 1: // bxl
				B = B ^ operand
			case 2: // bst
				B = comboOperand % 8
			case 3: // jnz
				if A != 0 {
					ip = operand
					continue
				}
			case 4: // bxc
				B = B ^ C
			case 5: // out
				result = append(result, comboOperand % 8)
			case 6: // bdv
				B = A / (1 << comboOperand)
			case 7: // cdv
				C = A / (1 << comboOperand)
        }
        ip += 2
    }
	return result
}

func solvePart1(file []int) []int {
	A, B, C := file[0], file[1], file[2]
	program := file[3:]
    return run(A, B, C, program)
}

func solvePart2(file []int) {
	B, C := file[1], file[2]
	program := file[3:]
	/*
		2,4 B = A % 8 (3 first bits of A)
		1,1 B ^= 1
		7,5 C = A >> B (A shifted by at most 7)
		1,5 B ^= 5
		0,3 A >>= 3 (A shifted by 3)
		4,3 B ^= C (B will be a function of A's MSBs)
		5,5 print(B % 8) (3 first bits of B)
		3,0 if A != 0 then ip = 0

		A gets shifted by at most 1024 each iteration
	*/
	A := 0
	for i := len(program) - 1; i >= 0; i-- {
		A <<= 3
		for !slices.Equal(program[i:], run(A, B, C, program)) {
			A++
		}
	}
	fmt.Printf("A: %d\n", A)
	fmt.Printf("program: %v, output: %v\n", program, run(A, B, C, program))
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

	var numbers []int

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "Register") {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				value := strings.TrimSpace(parts[1])
				num, err := strconv.Atoi(value)
				if err != nil {
					return nil, fmt.Errorf("invalid number in Register line: %v", err)
				}
				numbers = append(numbers, num)
			}
		} else if strings.HasPrefix(line, "Program:") {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				values := strings.Split(parts[1], ",")
				for _, value := range values {
					value = strings.TrimSpace(value)
					num, err := strconv.Atoi(value)
					if err != nil {
						return nil, fmt.Errorf("invalid number in Program line: %v", err)
					}
					numbers = append(numbers, num)
				}
			}
		}
	}

	return numbers, nil
}
