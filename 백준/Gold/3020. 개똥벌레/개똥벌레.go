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
	h, _ := scanInt()
	slice := make([]int, h)
	for i := 0; i < n; i++ {
		b, _ := scanInt()
		if i%2 == 0 {
			slice[0]++
			if h > b {
				slice[b]--
			}
		} else {
			slice[h-b]++
		}
	}
	for i := 1; i < h; i++ {
		slice[i] += slice[i-1]
	}

	result := [2]int{0, 0}
	result[0] = slice[0]
	result[1] = 1
	for i := 1; i < h; i++ {
		if slice[i] < result[0] {
			result[0] = slice[i]
			result[1] = 1
		} else if slice[i] == result[0] {
			result[1]++
		}
	}
	w.WriteString(strconv.Itoa(result[0]) + " " + strconv.Itoa(result[1]))
	w.WriteString("\n")
}