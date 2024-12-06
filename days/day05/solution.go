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

type Set map[interface{}]struct{}

func NewSet() Set {
    return make(Set)
}

func (s Set) Add(element interface{}) {
    s[element] = struct{}{}
}

func (s Set) Remove(element interface{}) {
    delete(s, element)
}

func (s Set) Contains(element interface{}) bool {
    _, exists := s[element]
    return exists
}

func (s Set) Size() int {
    return len(s)
}

func topologicalSort(adj map[int][]int) ([]int, error) {
	visited := make(map[int]bool)
	recursionStack := make(map[int]bool)
	result := []int{}
	var dfs func(node int) error

	dfs = func(node int) error {
		if recursionStack[node] {
			return fmt.Errorf("cycle detected")
		}
		if visited[node] {
			return nil
		}

		recursionStack[node] = true

		for _, neighbor := range adj[node] {
			if err := dfs(neighbor); err != nil {
				return err
			}
		}

		recursionStack[node] = false
		visited[node] = true

		result = append(result, node)
		return nil
	}

	for node := range adj {
		if !visited[node] {
			if err := dfs(node); err != nil {
				return nil, err
			}
		}
	}

	reversedResult := reverse(result)
	return reversedResult, nil
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func fix(adj map[int][]int, update []int) error {
	for _, node := range update {
		if _, exists := adj[node]; !exists {
			adj[node] = []int{}
		}
	}

	sortedNodes, err := topologicalSort(adj)
	if err != nil {
		return err
	}

	sortedIndex := make(map[int]int)
	for i, node := range sortedNodes {
		sortedIndex[node] = i
	}

	for i := 0; i < len(update); i++ {
		for j := i + 1; j < len(update); j++ {
			if sortedIndex[update[i]] > sortedIndex[update[j]] {
				update[i], update[j] = update[j], update[i]
			}
		}
	}

	return nil
}

func solvePart1(constraints Set, updates [][]int) int {
	var result int = 0

	for _, update := range updates {
		var valid bool = true

		for i := 0; i < len(update); i++ {
			for j := i + 1; j < len(update) && valid; j++ {
				if constraints.Contains(fmt.Sprintf("%d|%d", update[j], update[i])) {
					valid = false
					break
				}
			}
		}
		
		if valid {
			result += update[len(update) / 2]
		}
	}

	return result
}

func solvePart2(constraints Set, updates [][]int) (int, error) {
	var result int = 0

	for _, update := range updates {
		var valid bool = true
		var adj map[int][]int = make(map[int][]int)

		for i := 0; i < len(update); i++ {
			for j := i + 1; j < len(update); j++ {
				a, b := update[i], update[j]

				edge := fmt.Sprintf("%d|%d", a, b)
				edge_t := fmt.Sprintf("%d|%d", b, a)
				
				if constraints.Contains(edge_t) {
					valid = false
					adj[b] = append(adj[b], a)
				}

				if constraints.Contains(edge) {
					adj[a] = append(adj[a], b)
				}
			}
		}
		
		if !valid {
			err := fix(adj, update)
			if err != nil {
				return 0, err
			}
			
			result += update[len(update) / 2]
		}
	}

	return result, nil
}

func readInput(fileName string) (Set, [][]int, error) {
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

	lines := strings.Split(string(data), "\n")

	var constraints Set = NewSet()

	var i int = 0
	for i < len(lines) && lines[i] != "" {
		constraints.Add(lines[i])
		i++
	}

	var updates [][]int = [][]int{}
	for i < len(lines) {
		if lines[i] == "" {
			i++
			continue
		}

		parts := strings.Split(lines[i], ",")

		var update []int = []int{}
		for _, part := range parts {
			
			x, err := strconv.Atoi(part)
			if err != nil {
				return nil, nil, err
			}

			update = append(update, x)
		}

		updates = append(updates, update)

		i++
	}

	return constraints, updates, nil
}
