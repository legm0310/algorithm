package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	N      int
	C      [][]int
	result int = 0
	s      *bufio.Scanner
	w      *bufio.Writer
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

func recursion(idx, P int) {
	if idx > N {
		return
	} else if idx == N {
		result = max(result, P)
		return
	}
	recursion(idx+C[idx][0], P+C[idx][1])
	recursion(idx+1, P)
}

func main() {
	defer w.Flush()
	N, _ = scanInt()
	C = make([][]int, N)
	for i := 0; i < N; i++ {
		C[i] = make([]int, 2)
		C[i][0], _ = scanInt()
		C[i][1], _ = scanInt()
	}
	recursion(0, 0)
	w.WriteString(strconv.Itoa(result))
	w.WriteByte('\n')
}
