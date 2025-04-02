package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	tc, n, m, wh int
	dist         []int
	edges        []*Edge
	s            = bufio.NewScanner(os.Stdin)
	w            = bufio.NewWriter(os.Stdout)
)

type Edge struct {
	s, e, t int
}

func main() {
	defer w.Flush()
	s.Split(bufio.ScanWords)

	tc = scanInt()
	for i := 0; i < tc; i++ {
		n, m, wh = scanInt(), scanInt(), scanInt()
		edges = nil
		for j := 0; j < m; j++ {
			start, end, time := scanInt(), scanInt(), scanInt()
			edges = append(edges, &Edge{start, end, time})
			edges = append(edges, &Edge{end, start, time})
		}
		for j := 0; j < wh; j++ {
			start, end, time := scanInt(), scanInt(), scanInt()
			edges = append(edges, &Edge{start, end, -time})
		}

		for j := 1; j <= n; j++ {
			edges = append(edges, &Edge{0, j, 0}) // 가상 정점에서 모든 정점으로 연결
		}

		dist = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dist[j] = 0
		}

		hasCycle := false
		for j := 0; j < n; j++ { // 최단 경로는 최대 (정점 수 - 1)개의 간선을 지남 , n-1번 순회 이후에도 갱신되는지 체크하기 위해 n번까지 순회
			for _, edge := range edges {
				if dist[edge.s]+edge.t < dist[edge.e] { // s -> e의 간선으로 더 짧은 경로를 발견했는지 여부,
					dist[edge.e] = dist[edge.s] + edge.t
					if j == n-1 { // n-1개의 간선 이후 또 갱신 되었을 경우
						hasCycle = true
					}
				}
			}
		}

		if hasCycle {
			w.WriteString("YES")
			w.WriteByte('\n')
		} else {
			w.WriteString("NO")
			w.WriteByte('\n')
		}
	}
}

func scanInt() int {
	s.Scan()
	x, _ := strconv.Atoi(s.Text())
	return x
}
