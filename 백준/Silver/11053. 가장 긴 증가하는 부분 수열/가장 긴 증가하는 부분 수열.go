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
	n, _ := scanInt()
	seq := make([]int, n)
	for i := 0; i < n; i++ {
		seq[i], _ = scanInt()
	}

	lis := make([]int, n)

	for i := 0; i < n; i++ {
		lis[i] = 1
		for j := 0; j < i; j++ {
			if seq[j] < seq[i] {
				lis[i] = max(lis[i], lis[j]+1)
			}
		}
	}
	result := 0
	for _, num := range lis {
		if result < num {
			result = num
		}
	}

	w.WriteString(strconv.Itoa(result))
	w.WriteByte('\n')
}