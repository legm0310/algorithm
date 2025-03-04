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
	lcs := make([]int, max(n, m))

	for i := 0; i < n; i++ {
		ml := 0
		for j := 0; j < m; j++ {
			if ml < lcs[j] {
				ml = lcs[j]
			} else if str1[i] == str2[j] {
				lcs[j] = ml + 1
			}
		}
	}
	result := 0
	for i := 0; i < m; i++ {
		if lcs[i] > result {
			result = lcs[i]
		}
	}
	w.WriteString(strconv.Itoa(result))
	w.WriteByte('\n')
}
