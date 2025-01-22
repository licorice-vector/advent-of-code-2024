package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

func dfs(value, op, left, right map[int]int, node int) {
	if value[node] != -1 {
		return
	}
	l, r, o := left[node], right[node], op[node]
	dfs(value, op, left, right, l)
	dfs(value, op, left, right, r)
	switch o {
		case 1:
			value[node] = value[l] ^ value[r]
		case 2:
			value[node] = value[l] & value[r]
		case 3:
			value[node] = value[l] | value[r]
	}
}

func solvePart1(A []int, B [][]int, C map[string]int) int {
	n := len(A) / 2
	value := map[int]int{}
	op := map[int]int{}
	left := map[int]int{}
	right := map[int]int{}
	for _, b := range B {
		l, r, node, o := b[0], b[1], b[2], b[3]
		value[node] = -1
		op[node] = o
		left[node] = l
		right[node] = r
	}
	for i := 0; i < n; i++ {
		value[C[fmt.Sprintf("x%02d", i)]] = A[i]
		value[C[fmt.Sprintf("y%02d", i)]] = A[n + i]
	}
	result := 0
	pow2 := 1
	for i := 0; i < 100; i++ {
		if id, exists := C[fmt.Sprintf("z%02d", i)]; exists {
			dfs(value, op, left, right, id)
			result += value[id] * pow2
			pow2 *= 2
		}
	}
	return result
}

func solvePart2(A []int, B [][]int, C map[string]int) []string {
	edges := [][]string{}
	D := make(map[int]string)
	for key, value := range C {
		D[value] = key
	}
	swaps := [][]int{
		{C["z37"], C["rrn"]},
		{C["z16"], C["fkb"]},
		{C["z31"], C["rdn"]},
		{C["nnr"], C["rqf"]},
	}
	result := []string{}
	for _, swap := range swaps {
		result = append(result, D[swap[0]])
		result = append(result, D[swap[1]])
		rows := []int{}
		for i := 0; i < 2; i++ {
			for j, b := range B {
				if b[2] == swap[i] {
					rows = append(rows, j)
					break
				}
			}
		}
		B[rows[0]][2], B[rows[1]][2] = B[rows[1]][2], B[rows[0]][2]
	}
	for _, b := range B {
		l, r, node, o := b[0], b[1], b[2], b[3]
		op := ""
		switch o {
			case 1:
				op = "XOR"
			case 2:
				op = "AND"
			case 3:
				op = "OR"
		}
		op = fmt.Sprintf("%s %s%s", op, D[l], D[r])
		edges = append(edges, []string{op, D[l]})
		edges = append(edges, []string{op, D[r]})
		edges = append(edges, []string{D[node], op})
	}
	var sb strings.Builder
	sb.WriteString("digraph G {\n")
	for _, edge := range edges {
		sb.WriteString(fmt.Sprintf("  \"%s\" -> \"%s\";\n", edge[0], edge[1]))
	}
	sb.WriteString("}\n")
	// fmt.Println(sb.String())
	// use Graphviz to visualize the graph
	sort.Strings(result)
	return result
}

func readInput(fileName string) ([]int, [][]int, map[string]int, error) {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		return nil, nil, nil, errors.New("could not determine the current file")
	}

	dir := filepath.Dir(currentFile)
	filePath := filepath.Join(dir, fileName)

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, nil, nil, err
	}

	var values []int
	var operations [][]int
	idMap := make(map[string]int)
	idCounter := 0

	lines := strings.Split(string(data), "\n")
	i := 0

	for i < len(lines) {
		line := lines[i]

		line = strings.TrimSpace(line)
		if line == "" {
			i++
			break
		}

		pair := strings.Split(line, ", ")[0]
		parts := strings.Split(pair, ": ")
		key := strings.TrimSpace(parts[0])
		value, _ := strconv.Atoi(strings.TrimSpace(parts[1]))

		if _, exists := idMap[key]; !exists {
			idMap[key] = idCounter
			idCounter++
		}

		values = append(values, value)

		i++
	}

	for i < len(lines) {
		line := lines[i]

		tokens := strings.Fields(line)
		if len(tokens) < 5 {
			continue
		}

		op1, op, op2, _, result := tokens[0], tokens[1], tokens[2], tokens[3], tokens[4]

		for _, key := range []string{op1, op2, result} {
			if _, exists := idMap[key]; !exists {
				idMap[key] = idCounter
				idCounter++
			}
		}

		var operatorID int
		switch op {
		case "XOR":
			operatorID = 1
		case "AND":
			operatorID = 2
		case "OR":
			operatorID = 3
		default:
			return nil, nil, nil, fmt.Errorf("unknown operator: %s", op)
		}

		operations = append(operations, []int{idMap[op1], idMap[op2], idMap[result], operatorID})

		i++
	}

	return values, operations, idMap, nil
}
