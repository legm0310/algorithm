package main

import (
	"bufio"
	"os"
	"strconv"
)

const MAX = 100000

var (
	n, k int
	dist []int
	s    *bufio.Scanner
	w    *bufio.Writer
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

func bfs(start int) int {
	queue := make([]int, 0)
	queue = append(queue, start)
	dist[start] = 0

	for len(queue) > 0 {
		x := queue[0]
		queue = queue[1:]

		if x == k {
			return dist[x]
		}

		next := []int{x * 2, x - 1, x + 1}

		for _, nx := range next {
			if nx >= 0 && nx <= MAX && dist[nx] == -1 {
				if nx == x*2 {
					dist[nx] = dist[x]
					queue = append([]int{nx}, queue...)
				} else {
					dist[nx] = dist[x] + 1
					queue = append(queue, nx)
				}
			}
		}
	}
	return -1
}

func main() {
	defer w.Flush()
	n, _ = scanInt()
	k, _ = scanInt()
	dist = make([]int, MAX+1)
	for i := range dist {
		dist[i] = -1
	}
	w.WriteString(strconv.Itoa(bfs(n)))
	w.WriteByte('\n')
}
