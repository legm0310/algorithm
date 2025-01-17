package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	N   int
	M   int
	arr []int
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

func recursion(num, dep int) {
	if dep == M {
		for i := 0; i < M; i++ {
			w.WriteString(strconv.Itoa(arr[i]))
			w.WriteByte(' ')
		}
		w.WriteByte('\n')
		return
	}
	for i := 0; i < N; i++ {
		if dep > 0 && arr[dep-1] > i+1 {
			continue
		}
		arr[dep] = i + 1
		recursion(num, dep+1)
	}
}

func main() {
	defer w.Flush()
	N, _ = scanInt()
	M, _ = scanInt()
	arr = make([]int, M)
	recursion(1, 0)
}
