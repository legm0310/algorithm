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
	dp = make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, K+1)
	}

	for weight := 0; weight <= K; weight++ {
		if weight >= T[0][0] {
			dp[0][weight] = T[0][1]
		}
	}

	for idx := 1; idx < N; idx++ {
		for weight := 1; weight < K+1; weight++ {
			if idx == 0 {
				dp[idx][weight] = 0
			} else if weight < T[idx][0] {
				dp[idx][weight] = dp[idx-1][weight]
			} else {
				dp[idx][weight] = max(dp[idx-1][weight-T[idx][0]]+T[idx][1], dp[idx-1][weight])
			}
		}
	}
	w.WriteString(strconv.Itoa(dp[N-1][K]))
	w.WriteByte('\n')
}
