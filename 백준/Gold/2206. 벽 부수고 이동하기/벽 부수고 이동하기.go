package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	dx      = []int{-1, 1, 0, 0}
	dy      = []int{0, 0, -1, 1}
	n, m    int
	table   [][]int
	visited [][][2]int
	r       = bufio.NewReader(os.Stdin)
	w       = bufio.NewWriter(os.Stdout)
)

type State struct {
	x, y    int
	isBreak int
	dist    int
}

func main() {
	defer w.Flush()
	line, _ := r.ReadString('\n')
	parts := strings.Fields(line)
	n, _ = strconv.Atoi(parts[0])
	m, _ = strconv.Atoi(parts[1])

	table = make([][]int, n)
	visited = make([][][2]int, n)
	for i := 0; i < n; i++ {
		line, _ = r.ReadString('\n')
		table[i] = make([]int, m)
		visited[i] = make([][2]int, m)
		for j := 0; j < m; j++ {
			table[i][j] = parse(line[j])
		}
	}
	w.WriteString(strconv.Itoa(find(0, 0)))
	w.WriteByte('\n')
}

func find(x, y int) int {
	q := make([]*State, 0)
	q = append(q, &State{x, y, 0, 1})
	visited[x][y][0] = 1

	for len(q) > 0 {
		ex, ey, isBreak, dist := q[0].x, q[0].y, q[0].isBreak, q[0].dist
		q = q[1:]

		if ex == n-1 && ey == m-1 {
			return dist
		}

		for i := 0; i < 4; i++ {
			nx, ny := ex+dx[i], ey+dy[i]
			if nx >= 0 && ny >= 0 && nx < n && ny < m {
				if table[nx][ny] == 1 {
					//이미 깼거나 벽을 깨고 온적 있으면 패스
					if isBreak == 1 || visited[nx][ny][1] == 1 {
						continue
					}
					//벽을 깨고 이동할 수 있는 조건 충족
					visited[nx][ny][1] = 1
					q = append(q, &State{nx, ny, 1, dist + 1})
				} else {
					if visited[nx][ny][isBreak] == 0 {
						visited[nx][ny][isBreak] = 1
						q = append(q, &State{nx, ny, isBreak, dist + 1})
					}
				}
			}
		}
	}
	return -1
}

func parse(input byte) int {
	return int(input - '0')
}