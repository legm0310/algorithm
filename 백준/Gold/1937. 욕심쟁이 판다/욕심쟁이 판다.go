package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	N       int
	treeMap [][]int
	dp      [][]int
	answer  int = 0
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

func recursion(x, y int) int {
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

    if dp[x][y] != 0 {
		return dp[x][y]
	}
    
	for d := 0; d < 4; d++ {
		nx := dx[d] + x
		ny := dy[d] + y
		if 0 <= nx && N > nx && 0 <= ny && N > ny {
			if treeMap[x][y] < treeMap[nx][ny] {
				dp[x][y] = max(dp[x][y], recursion(nx, ny)+1)
			}
		}
	}
	return dp[x][y]
}

func main() {
	defer w.Flush()
	N, _ = scanInt()
	treeMap = make([][]int, N)
	for i := 0; i < N; i++ {
		treeMap[i] = make([]int, N)
		for j := 0; j < N; j++ {
			treeMap[i][j], _ = scanInt()
		}
	}
	dp = make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, N)
	}

	for x := 0; x < N; x++ {
		for y := 0; y < N; y++ {
			recursion(x, y)
		}
	}
	for x := 0; x < N; x++ {
		for y := 0; y < N; y++ {
			if dp[x][y] > answer {
				answer = dp[x][y]
			}
		}
	}
	w.WriteString(strconv.Itoa(answer + 1))
	w.WriteByte('\n')
}
