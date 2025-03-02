package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	m    int
	n    int
	load [][]int
	dp   [][]int
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

func recursion(x, y int) int {
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	route := 0

	if x == m-1 && y == n-1 {
		return 1
	}

	if dp[x][y] != -1 {
		return dp[x][y]
	}

	for d := 0; d < 4; d++ {
		nx := dx[d] + x
		ny := dy[d] + y
		if 0 <= nx && m > nx && 0 <= ny && n > ny {
			if load[x][y] > load[nx][ny] {
				route += recursion(nx, ny)
			}
		}
	}
	dp[x][y] = route
	return dp[x][y]
}

func main() {
	defer w.Flush()
	m, _ = scanInt()
	n, _ = scanInt()
	load = make([][]int, m)
	dp = make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		load[i] = make([]int, n)
		for j := 0; j < n; j++ {
			load[i][j], _ = scanInt()
			dp[i][j] = -1
		}
	}

	w.WriteString(strconv.Itoa(recursion(0, 0)))
	w.WriteByte('\n')
}
