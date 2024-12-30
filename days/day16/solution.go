package main

import (
	"container/heap"
	"errors"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Item struct {
	node int
	cost int
}
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].cost < pq[j].cost }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func dijkstra(adj [][][]int, sources []int, N int) []int {
	dist := make([]int, N)
	for i := 0; i < N; i++ {
		dist[i] = math.MaxInt
	}
	pq := &PriorityQueue{}
	heap.Init(pq)
	for _, src := range sources {
		dist[src] = 0
		heap.Push(pq, &Item{node: src, cost: 0})
	}

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		u := item.node
		if item.cost > dist[u] {
			continue
		}
		for _, edge := range adj[u] {
			v, weight := edge[0], edge[1]
			if dist[u]+weight < dist[v] {
				dist[v] = dist[u] + weight
				heap.Push(pq, &Item{node: v, cost: dist[v]})
			}
		}
	}
	return dist
}

func solvePart1(board []string) int {
	n, m := len(board), len(board[0])

	//            U  R  D   L
	dirX := []int{0, 1, 0, -1}
	dirY := []int{-1, 0, 1, 0}

	N := n * m
	adj := make([][][]int, N * 4)

	for i := 0; i < N; i++ {
		for dir := 0; dir < 4; dir++ {
			node := N * dir + i
			adj[node] = append(adj[node], []int{N * ((dir + 1) % 4) + i, 1000})
			adj[node] = append(adj[node], []int{N * ((dir - 1 + 4) % 4) + i, 1000})
		}
	}

	src, dstX, dstY := -1, -1, -1

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'S' {
				src = N + (i * m + j) // east
			} else if board[i][j] == 'E' {
				dstX, dstY = j, i
			}
			for dir := 0; dir < 4; dir++ {
				node := N * dir + (i * m + j)
				ni, nj := i + dirY[dir], j + dirX[dir]
				if ni < 0 || ni >= n || nj < 0 || nj >= m || board[ni][nj] == '#' {
					continue
				}
				next := N * dir + (ni * m + nj)
				adj[node] = append(adj[node], []int{next, 1})
			}
		}
	}

	cost := math.MaxInt
	dist := dijkstra(adj, []int{src}, N * 4)

	for dir := 0; dir < 4; dir++ {
		cost = min(cost, dist[N * dir + (dstY * m + dstX)])
	}

	return cost
}

func solvePart2(board []string) int {
	n, m := len(board), len(board[0])

	//            U  R  D   L
	dirX := []int{0, 1, 0, -1}
	dirY := []int{-1, 0, 1, 0}

	N := n * m
	adj := make([][][]int, N * 4)

	for i := 0; i < N; i++ {
		for dir := 0; dir < 4; dir++ {
			node := N * dir + i
			adj[node] = append(adj[node], []int{N * ((dir + 1) % 4) + i, 1000})
			adj[node] = append(adj[node], []int{N * ((dir - 1 + 4) % 4) + i, 1000})
		}
	}

	src, dstX, dstY := -1, -1, -1

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'S' {
				src = N + (i * m + j) // east
			} else if board[i][j] == 'E' {
				dstX, dstY = j, i
			}
			for dir := 0; dir < 4; dir++ {
				node := N * dir + (i * m + j)
				ni, nj := i + dirY[dir], j + dirX[dir]
				if ni < 0 || ni >= n || nj < 0 || nj >= m || board[ni][nj] == '#' {
					continue
				}
				next := N * dir + (ni * m + nj)
				adj[node] = append(adj[node], []int{next, 1})
			}
		}
	}

	distFromSrc := dijkstra(adj, []int{src}, N * 4)

	reverseAdj := make([][][]int, N * 4)
	for u := 0; u < len(adj); u++ {
		for _, edge := range adj[u] {
			v, weight := edge[0], edge[1]
			reverseAdj[v] = append(reverseAdj[v], []int{u, weight})
		}
	}

	reverseStartNodes := []int{}
	for dir := 0; dir < 4; dir++ {
		reverseStartNodes = append(reverseStartNodes, N * dir + (dstY * m + dstX))
	}
	distToDst := dijkstra(reverseAdj, reverseStartNodes, N * 4)

	minCost := distToDst[src]

	B := make([][]int, n)
	for i := range B {
		B[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == '#' {
				continue
			}
			for dir := 0; dir < 4; dir++ {
				node := N * dir + i * m + j
				if distFromSrc[node] + distToDst[node] == minCost {
					B[i][j] = 1
				}
			}
		}
	}
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if B[i][j] == 1 {
				count++
			}
		}
	}
	return count
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

	var board []string = []string{}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		board = append(board, line)
	}

	return board, nil
}
