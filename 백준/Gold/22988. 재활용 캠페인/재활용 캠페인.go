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
	x, _ := scanInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i], _ = scanInt()
	}
	slices.Sort(arr)

	start := 0
	end := n - 1

	remain := 0
	cnt := 0
	for start <= end {
		if arr[end] == x {
			cnt++
			end--
			continue
		}
		if start == end {
			remain++
			break
		}
		if float64(arr[start]+arr[end]) >= float64(x)/2 {
			cnt++
			start++
			end--
		} else {
			start++
			remain++
		}
	}
	w.WriteString(strconv.Itoa(cnt + (remain / 3)))
	w.WriteByte('\n')
}
