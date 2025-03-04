package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
)

var (
	n, m int
	seq  []int
	s    *bufio.Scanner
	w    *bufio.Writer
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

func recur(idx int, arr *[]int) {
	if len(*arr) == m {
		for _, num := range *arr {
			w.WriteString(strconv.Itoa(num) + " ")
		}
		w.WriteByte('\n')
		return
	}
	used := make(map[int]bool)
	for i := idx; i < n; i++ {
		if used[seq[i]] {
			continue
		}
		used[seq[i]] = true
		*arr = append(*arr, seq[i])
		recur(i, arr)
		*arr = (*arr)[:len(*arr)-1]
	}

}

func main() {
	defer w.Flush()
	n, _ = scanInt()
	m, _ = scanInt()
	seq = make([]int, n)
	for i := 0; i < n; i++ {
		seq[i], _ = scanInt()
	}
	slices.Sort(seq)
	arr := make([]int, 0)
	recur(0, &arr)
}
