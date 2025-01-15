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

func recursion(dep int) {
	if dep == M {
		for i := 0; i < M; i++ {
			w.WriteString(strconv.Itoa(arr[i]))
			w.WriteByte(' ')
		}
		w.WriteByte('\n')
		return
	}
	for i := 1; i <= N; i++ {
		arr[dep] = i
		recursion(dep + 1)
	}
}

func main() {
	defer w.Flush()
	N, _ = scanInt()
	M, _ = scanInt()
	arr = make([]int, M+1)
	recursion(0)
}
