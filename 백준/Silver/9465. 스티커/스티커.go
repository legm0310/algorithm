package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	t, n  int
	graph [2][]int
	dist  [2][]int
    r *bufio.Reader
	w *bufio.Writer
)

func init() {
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
}

func recur(x, y int) int {
	if y == n-1 {
		dist[x][y] = graph[x][y]
		return dist[x][y]
	}
	if dist[x][y] != -1 {
		return dist[x][y]
	}

	nx := (x + 1) % 2
	dist[x][y] = max(dist[x][y], recur(nx, y+1)+graph[x][y])

	dist[x][y] = max(dist[x][y], recur(x, y+1))
	return dist[x][y]
}

func main() {
	defer w.Flush()
	fmt.Fscan(r, &t)

	for i := 0; i < t; i++ {
		fmt.Fscan(r, &n)
		for j := 0; j < 2; j++ {
			graph[j] = make([]int, n)
			dist[j] = make([]int, n+1)
			for k := 0; k < n; k++ {
				fmt.Fscan(r, &graph[j][k])
				dist[j][k] = -1
			}
			dist[j][n] = -1
		}

		fmt.Fprintln(w, max(recur(0, 0), recur(1, 0)))
	}
}
