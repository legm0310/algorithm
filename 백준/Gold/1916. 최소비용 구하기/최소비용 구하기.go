package main

import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
)

var (
	n, m, start, end int
	maxAmount        int = 99999999
	edges            map[int][]*Edge
	dist             []int
	s                *bufio.Scanner
	w                *bufio.Writer
)

type Edge struct {
	v, w int
}

func main() {
	defer w.Flush()
	n, m = scanInt(), scanInt()
	edges = make(map[int][]*Edge)
	dist = make([]int, n+1)
	for i := 0; i < m; i++ {
		city := scanInt()
		edges[city] = append(edges[city], &Edge{v: scanInt(), w: scanInt()})
	}
	start, end = scanInt(), scanInt()

	bfs(start, 0)
	w.WriteString(strconv.Itoa(dist[end]) + "\n")
}

type PriorityQueue []*Edge

func (pq *PriorityQueue) Len() int { return len(*pq) }
func (pq *PriorityQueue) Less(i, j int) bool {
	return (*pq)[i].w < (*pq)[j].w
}
func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}
func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(*Edge))
}
func (pq *PriorityQueue) Pop() any {
	old := *pq
	l := len(old)
	x := old[l-1]
	*pq = old[0 : l-1]
	return x
}

func bfs(startCity, weight int) {
	for i := 0; i <= n; i++ {
		dist[i] = maxAmount
	}

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Edge{v: startCity, w: weight})
	dist[startCity] = weight

	for pq.Len() > 0 {
		e := heap.Pop(pq).(*Edge)

		if dist[e.v] < e.w {
			continue
		}

		for _, eg := range edges[e.v] {
			if dist[eg.v] > eg.w+e.w {
				dist[eg.v] = eg.w + e.w
				heap.Push(pq, &Edge{v: eg.v, w: dist[eg.v]})
			}
		}
	}
}

func init() {
	s = bufio.NewScanner(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
	s.Split(bufio.ScanWords)
}

func scanInt() int {
	s.Scan()
	x, _ := strconv.Atoi(s.Text())
	return x
}
