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
	maxNum, minNum := -1000000, 1000000
	for i := 0; i < n; i++ {
		if arr[i] > maxNum {
			maxNum = arr[i]
		}
		if arr[i] < minNum {
			minNum = arr[i]
		}

	}
	w.WriteString(strconv.Itoa(minNum) + " " + strconv.Itoa(maxNum))
	w.WriteByte('\n')
}
