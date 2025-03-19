package main

import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
)

var (
	vertexCnt, edgeCnt, k int
	edges                 map[int][]*Edge
	dist                  []int
	s                     *bufio.Scanner
	w                     *bufio.Writer
)

type Edge struct {
	v, w int
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
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
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

func main() {
	defer w.Flush()
	vertexCnt, edgeCnt, k = scanInt(), scanInt(), scanInt()

	edges = make(map[int][]*Edge)
	for i := 0; i < edgeCnt; i++ {
		u := scanInt()
		edges[u] = append(edges[u], &Edge{scanInt(), scanInt()})
	}

	dist = make([]int, vertexCnt+1)
	for i := 0; i < vertexCnt+1; i++ {
		dist[i] = 10000000
	}

	bfs(k)

	for i := 1; i < vertexCnt+1; i++ {
		if dist[i] == 10000000 {
			w.WriteString("INF\n")
		} else {
			w.WriteString(strconv.Itoa(dist[i]) + "\n")
		}
	}
}

func bfs(v int) {
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Edge{v, 0})
	dist[v] = 0

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(*Edge)
		for _, nextEdge := range edges[cur.v] {
			if dist[cur.v]+nextEdge.w < dist[nextEdge.v] {
				dist[nextEdge.v] = dist[cur.v] + nextEdge.w
				heap.Push(pq, &Edge{nextEdge.v, dist[nextEdge.v]})
			}
		}
	}
}
