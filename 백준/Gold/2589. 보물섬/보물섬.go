package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	X, Y    int
	visited [][]int
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
	queue := make([][]int, 0, X*Y)
	queue = append(queue, []int{x, y})
	visited[x][y] = 1

	for len(queue) > 0 {
		ex, ey := queue[0][0], queue[0][1]
		queue = queue[1:]

		for d := 0; d < 4; d++ {
			nx, ny := ex+dx[d], ey+dy[d]
			if 0 <= nx && nx < X && 0 <= ny && ny < Y {
				if graph[nx][ny] == 'L' && visited[nx][ny] == 0 {
					visited[nx][ny] = 1
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
	visited = make([][]int, X)
	dist = make([][]int, X)
	for i := 0; i < X; i++ {
		s.Scan()
		graph[i] = s.Bytes()
	}

	for i := 0; i < X; i++ {
		for j := 0; j < Y; j++ {
			if graph[i][j] == 'L' {
				for k := 0; k < X; k++ {
					visited[k] = make([]int, Y)
					dist[k] = make([]int, Y)
				}
				bfs(i, j)
			}
		}
	}
	w.WriteString(strconv.Itoa(cnt) + "\n")
}
