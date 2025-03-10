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
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i], _ = scanInt()
	}
	v, _ := scanInt()
	result := 0
	for i := 0; i < n; i++ {
		if arr[i] == v {
			result++
		}
	}
	w.WriteString(strconv.Itoa(result) + "\n")
}
