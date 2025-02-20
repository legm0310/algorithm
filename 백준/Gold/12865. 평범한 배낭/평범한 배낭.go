package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	N  int
	K  int
	T  [][]int
	dp [][]int
	s  *bufio.Scanner
	w  *bufio.Writer
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

func recursion(idx, weight int) int {
	if weight > K {
		return -9999999
	}
	if idx == N {
		return 0
	}
	if dp[idx][weight] != -1 {
		return dp[idx][weight]
	}
	dp[idx][weight] = max(recursion(idx+1, weight+T[idx][0])+T[idx][1], recursion(idx+1, weight))
	return dp[idx][weight]
}

func main() {
	defer w.Flush()
	N, _ = scanInt()
	K, _ = scanInt()
	T = make([][]int, N)
	for i := 0; i < N; i++ {
		T[i] = make([]int, 2)
		T[i][0], _ = scanInt()
		T[i][1], _ = scanInt()
	}
	dp = make([][]int, N)
	for i := range dp {
		dp[i] = make([]int, 100001)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	w.WriteString(strconv.Itoa(recursion(0, 0)))
	w.WriteByte('\n')
}
