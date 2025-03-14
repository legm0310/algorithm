package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	visited []int
	graph   [][]int
	s       *bufio.Scanner
	w       *bufio.Writer
)

func init() {
	s = bufio.NewScanner(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
	s.Split(bufio.ScanWords)
}

func scanInt() (int, error) {
	s.Scan()
	return strconv.Atoi(s.Text())
}

func dfs(node int) {
	visited[node] = 1
	for _, next := range graph[node] {
		if visited[next] == 1 {
			continue
		}
		dfs(next)
	}
}

func main() {
	defer w.Flush()
	n, _ := scanInt()
	m, _ := scanInt()
	visited = make([]int, n+1)
	graph = make([][]int, n+1)

	for i, _ := range graph {
		visited[i] = 0
		graph[i] = make([]int, 0)
	}

	for i := 0; i < m; i++ {
		a, _ := scanInt()
		b, _ := scanInt()
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	dfs(1)

	sum := 0
	for _, cnt := range visited {
		sum += cnt
	}
	w.WriteString(strconv.Itoa(sum-1) + "\n")
}
