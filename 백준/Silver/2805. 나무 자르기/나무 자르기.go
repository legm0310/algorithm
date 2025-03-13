package main

import (
	"bufio"
	"os"
	"slices"
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
	m, _ := scanInt()
	trees := make([]int, n)
	for i := 0; i < n; i++ {
		trees[i], _ = scanInt()
	}

	start := 1
	end := slices.Max(trees)

	for start <= end {
		mid := (start + end) / 2
		wood := 0

		for _, tree := range trees {
			if tree >= mid {
				wood += tree - mid
			}
		}
		if wood >= m {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	w.WriteString(strconv.Itoa(end) + "\n")
}
