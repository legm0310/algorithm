package main

import (
	"bufio"
	"container/list"
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

func scanInt() (int, error) {
	s.Scan()
	return strconv.Atoi(s.Text())
}

func bfs(x, y int) {
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	queue := list.New()
	queue.PushBack([]int{x, y})
	visited[x][y] = 1

	for queue.Len() > 0 {
		target := queue.Remove(queue.Front()).([]int)
		ex, ey := target[0], target[1]

		for d := 0; d < 4; d++ {
			nx, ny := ex+dx[d], ey+dy[d]
			if 0 <= nx && nx < X && 0 <= ny && ny < Y {
				if string(graph[nx][ny]) == "L" && visited[nx][ny] == 0 {
					visited[nx][ny] = 1
					dist[nx][ny] = dist[ex][ey] + 1
					cnt = max(cnt, dist[nx][ny])
					queue.PushBack([]int{nx, ny})
				}
			}
		}
	}
}

func main() {
	defer w.Flush()
	X, _ = scanInt()
	Y, _ = scanInt()
	graph = make([][]byte, X)
	for i := 0; i < X; i++ {
		graph[i] = make([]byte, Y)
		s.Scan()
		graph[i] = s.Bytes()
	}

	for i := 0; i < X; i++ {
		for j := 0; j < Y; j++ {
			if string(graph[i][j]) == "L" {
				visited = make([][]int, X)
				dist = make([][]int, X)
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
