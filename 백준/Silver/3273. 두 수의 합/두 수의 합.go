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
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i], _ = scanInt()
	}
	x, _ := scanInt()
	slices.Sort(arr)

	start := 0
	end := n - 1
	cnt := 0
	for start < end {
		if arr[start]+arr[end] == x {
			cnt++
		}
		if arr[start]+arr[end] >= x {
			end--
		}
		if arr[start]+arr[end] < x {
			start++
		}
	}

	w.WriteString(strconv.Itoa(cnt))
	w.WriteByte('\n')

}
