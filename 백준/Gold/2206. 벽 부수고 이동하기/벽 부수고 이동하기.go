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
	ans     = 1 << 20
	table   [][]int
	visited [][][2]int
	r       = bufio.NewReader(os.Stdin)
	w       = bufio.NewWriter(os.Stdout)
)

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
	q := make([][]int, 0)
	q = append(q, []int{x, y, 0, 1})
	visited[x][y][0] = 1
	for len(q) > 0 {
		ex, ey, isBreak, dist := q[0][0], q[0][1], q[0][2], q[0][3]
		q = q[1:]
		if ex == n-1 && ey == m-1 {
			return dist
		}
		for i := 0; i < 4; i++ {
			nx, ny := ex+dx[i], ey+dy[i]
			//좌표 이동 기본 조건
			if nx >= 0 && ny >= 0 && nx < n && ny < m {
				//for j := 0; j < 2; j++ {
				// 1 벽인데 깬적 있을 때
				// 2 벽인데 깬적 없을 때
				// 3 벽이 아닐 떄
				// 4 방문 여부
				if table[nx][ny] == 1 && isBreak == 1 {
					continue
				}
				if table[nx][ny] == 1 && isBreak == 0 {
					if visited[nx][ny][1] == 1 {
						continue
					}
					visited[nx][ny][1] = 1
					q = append(q, []int{nx, ny, 1, dist + 1})
				}
				if table[nx][ny] == 0 && visited[nx][ny][isBreak] == 0 {
					visited[nx][ny][isBreak] = 1
					q = append(q, []int{nx, ny, isBreak, dist + 1})
				}
			}
		}
	}
	return -1
}

func parse(input byte) int {
	return int(input - '0')
}
