package main

import (
	"bufio"
	"os"
	"slices"
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
		if slices.IsSorted(arr) == false {
			return
		}
		for i := 0; i < M; i++ {
			w.WriteString(strconv.Itoa(arr[i]))
			w.WriteByte(' ')
		}
		w.WriteByte('\n')
		return
	}
	for i := num; i <= N; i++ {
		arr[dep] = i
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
