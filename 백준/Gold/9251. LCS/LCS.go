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

func main() {
	defer w.Flush()
	s.Scan()
	str1 := s.Text()
	s.Scan()
	str2 := s.Text()
	n, m := len(str1), len(str2)
	lcs := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		lcs[i] = make([]int, m+1)
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if str1[i-1] == str2[j-1] {
				lcs[i][j] = lcs[i-1][j-1] + 1
			} else {
				lcs[i][j] = max(lcs[i-1][j], lcs[i][j-1])
			}
		}
	}

	w.WriteString(strconv.Itoa(lcs[n][m]))
	w.WriteByte('\n')
}
