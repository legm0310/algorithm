package main

import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
)

// X가 출발점
var (
	n, m, x int
	graph   [][]*Edge
	reverse [][]*Edge
	dist    []int
	ans     = 0
	s       = bufio.NewScanner(os.Stdin)
	w       = bufio.NewWriter(os.Stdout)
)

type Edge struct {
	dest, weight int
}

type PQ []*Edge

func (pq *PQ) Len() int { return len(*pq) }

func (pq *PQ) Less(i, j int) bool {
	return (*pq)[i].weight < (*pq)[j].weight
}

func (pq *PQ) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PQ) Push(x any) {
	*pq = append(*pq, x.(*Edge))
}

func (pq *PQ) Pop() any {
	old := *pq
	length := len(old)
	pop := old[length-1]
	*pq = old[:length-1]
	return pop
}

func main() {
	defer w.Flush()
	s.Split(bufio.ScanWords)
	n, m, x = scanInt(), scanInt(), scanInt()
	graph = make([][]*Edge, m+1)
	reverse = make([][]*Edge, m+1)
	for i := 0; i < m; i++ {
		start, end, t := scanInt(), scanInt(), scanInt()
		graph[start] = append(graph[start], &Edge{end, t})
		reverse[end] = append(reverse[end], &Edge{start, t})
	}

	//역방향 그래프 다익스트라 --> 여러 점에서(i) x까지의 최단거리
	itox := find(x, reverse)
	xtoi := find(x, graph)
	for i := 1; i < n+1; i++ {
		ans = max(ans, itox[i]+xtoi[i])
	}
	w.WriteString(strconv.Itoa(ans))
	w.WriteByte('\n')
}

func find(node int, input [][]*Edge) []int {
	dist = make([]int, n+1)
	for i := range dist {
		dist[i] = 1 << 28
	}
	dist[node] = 0
	pq := &PQ{}
	heap.Init(pq)
	heap.Push(pq, &Edge{node, 0})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(*Edge)
		if cur.weight > dist[cur.dest] {
			continue
		}
		for _, nextEdge := range input[cur.dest] {
			cost := cur.weight + nextEdge.weight
			if cost < dist[nextEdge.dest] {
				dist[nextEdge.dest] = cost
				heap.Push(pq, &Edge{nextEdge.dest, cost})
			}
		}
	}
	return dist
}

func scanInt() int {
	s.Scan()
	num, _ := strconv.Atoi(s.Text())
	return num
}
