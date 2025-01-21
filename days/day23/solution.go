package main

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
)

func solvePart1(A []string) int {
    adj := map[string][]string{}
    for _, connection := range A {
        u := connection[0:2]
        v := connection[3:5]
        adj[u] = append(adj[u], v)
        adj[v] = append(adj[v], u)
    }

    tripleSet := make(map[string]bool)

    for a, neighbors := range adj {
        for _, b := range neighbors {
            if b <= a {
                continue
            }
            for _, c := range adj[b] {
                if c <= b || c == a {
                    continue
                }
                for _, neighbor := range adj[a] {
                    if neighbor == c {
                        if a[0] != 't' && b[0] != 't' && c[0] != 't' {
                            break
                        }
                        triangle := []string{a, b, c}
                        sort.Strings(triangle)
                        key := triangle[0] + triangle[1] + triangle[2]
                        tripleSet[key] = true
                        break
                    }
                }
            }
        }
    }

    return len(tripleSet)
}

func solvePart2(A []string) []string {
	adj := map[string][]string{}
    for _, connection := range A {
        u := connection[0:2]
        v := connection[3:5]
        adj[u] = append(adj[u], v)
        adj[v] = append(adj[v], u)
    }

    var largestClique []string
    var dfs func(current, candidates []string)
    dfs = func(current, candidates []string) {
        if len(candidates) == 0 {
            if len(largestClique) < len(current) {
                largestClique = append([]string{}, current...)
            }
            return
        }

        for i, node := range candidates {
            newCurrent := append(current, node)
            newCandidates := []string{}
            for _, neighbor := range candidates[i + 1:] {
                for _, n := range adj[node] {
                    if neighbor == n {
                        newCandidates = append(newCandidates, neighbor)
                        break
                    }
                }
            }
            dfs(newCurrent, newCandidates)
        }
    }

    var allNodes []string
    for node := range adj {
        allNodes = append(allNodes, node)
    }

    sort.Strings(allNodes)
    dfs([]string{}, allNodes)
    sort.Strings(largestClique)

    return largestClique
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
