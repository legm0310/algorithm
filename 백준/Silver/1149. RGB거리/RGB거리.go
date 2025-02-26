package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	s *bufio.Scanner
	w *bufio.Writer
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
	N, _ := scanInt()
	input := make([][]int, N)
	dp := make([][]int, N)
	for i, _ := range dp {
		input[i] = make([]int, 3)
		dp[i] = make([]int, 3)
		input[i][0], _ = scanInt()
		input[i][1], _ = scanInt()
		input[i][2], _ = scanInt()
	}

	dp[0][0] = input[0][0]
	dp[0][1] = input[0][1]
	dp[0][2] = input[0][2]
	
	for idx, _ := range dp {
		if idx == 0 {
			continue
		}
		dp[idx][0] = input[idx][0] + min(dp[idx-1][1], dp[idx-1][2])
		dp[idx][1] = input[idx][1] + min(dp[idx-1][0], dp[idx-1][2])
		dp[idx][2] = input[idx][2] + min(dp[idx-1][0], dp[idx-1][1])
	}
	w.WriteString(strconv.Itoa(min(min(dp[N-1][0], dp[N-1][1]), dp[N-1][2])))
	w.WriteByte('\n')
}
