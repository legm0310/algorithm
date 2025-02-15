package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

var (
	N   int
	T   [][]int
	ans int = 9999999999
	s   *bufio.Scanner
	w   *bufio.Writer
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

func recursion(idx, S, B, use int) {
	if idx == N {
		if use == 0 {
			return
		}
		result := int(math.Abs(float64(S - B)))
		ans = min(ans, result)
		return
	}
	recursion(idx+1, S*T[idx][0], B+T[idx][1], use+1)
	recursion(idx+1, S, B, use)
}

func main() {
	defer w.Flush()
	N, _ = scanInt()
	T = make([][]int, N)
	for i := 0; i < N; i++ {
		T[i] = make([]int, 2)
		T[i][0], _ = scanInt()
		T[i][1], _ = scanInt()
	}
	recursion(0, 1, 0, 0)
	w.WriteString(strconv.Itoa(ans))
	w.WriteByte('\n')
}
