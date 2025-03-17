package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	X, Y    int
	vc      int
	visited [51][51]int
	dist    [][]int
	graph   [][]byte
	cnt     int = 0
	s       *bufio.Scanner
	w       *bufio.Writer
)

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

func bfs(x, y int) {
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	vc++
	dist = make([][]int, X)
	for i := range dist {
		dist[i] = make([]int, Y)
	}
	queue := make([][]int, 0, X*Y)
	queue = append(queue, []int{x, y})
	visited[x][y] = vc
	for len(queue) > 0 {
		ex, ey := queue[0][0], queue[0][1]
		queue = queue[1:]

		for d := 0; d < 4; d++ {
			nx, ny := ex+dx[d], ey+dy[d]
			if 0 <= nx && nx < X && 0 <= ny && ny < Y {
				if visited[nx][ny] == vc {
					continue
				}
				if graph[nx][ny] == 'L' {
					visited[nx][ny] = vc
					dist[nx][ny] = dist[ex][ey] + 1
					cnt = max(cnt, dist[nx][ny])
					queue = append(queue, []int{nx, ny})
				}
			}
		}
	}
}

func main() {
	defer w.Flush()
	X = scanInt()
	Y = scanInt()
	graph = make([][]byte, X)
	for i := 0; i < X; i++ {
		s.Scan()
		graph[i] = s.Bytes()
	}

	for i := 0; i < X; i++ {
		for j := 0; j < Y; j++ {
			if graph[i][j] == 'L' {
				bfs(i, j)
			}
		}
	}
	w.WriteString(strconv.Itoa(cnt) + "\n")
}
