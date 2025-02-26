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
	if N == 0 {
		w.WriteString(strconv.Itoa(0))
		w.WriteByte('\n')
		return
	}
	dp := make([]int, N)

	for i := 0; i < N; i++ {
		if i == 0 {
			dp[i] = 1
		} else if i == 1 {
			dp[i] = 2
		} else {
			dp[i] = (dp[i-2] + dp[i-1]) % 10007
		}
	}
	w.WriteString(strconv.Itoa(dp[N-1]))
	w.WriteByte('\n')
}
