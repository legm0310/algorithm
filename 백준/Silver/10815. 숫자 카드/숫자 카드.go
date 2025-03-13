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
	cards := make([]int, n)
	for i := 0; i < n; i++ {
		cards[i], _ = scanInt()
	}

	m, _ := scanInt()
	arr := make([]int, m)
	for i := 0; i < m; i++ {
		arr[i], _ = scanInt()
	}
	slices.Sort(cards)
	for _, num := range arr {
		start := 0
		end := n - 1
		flag := false
		for start <= end {
			mid := (start + end) / 2
			if cards[mid] == num {
				flag = true
				break
			} else if cards[mid] < num {
				start = mid + 1
			} else if cards[mid] > num {
				end = mid - 1
			}
		}
		if flag {
			w.WriteString(strconv.Itoa(1) + " ")
		} else {
			w.WriteString(strconv.Itoa(0) + " ")
		}
	}
	w.WriteByte('\n')
}
