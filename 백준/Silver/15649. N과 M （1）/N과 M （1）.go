package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

var (
	n int
	m int
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

func recursion(num int, numArr []int) {
	if num == m {
		strArr := make([]string, len(numArr))
		for i, v := range numArr {
			strArr[i] = strconv.Itoa(v)
		}
		w.WriteString(strings.Join(strArr, " "))
		w.WriteString("\n")
		return
	}
	for i := 1; i < n+1; i++ {
		if slices.Contains(numArr, i) {
			continue
		}
		recursion(num+1, append(numArr, i))
	}
}

func main() {
	defer w.Flush()
	n, _ = scanInt()
	m, _ = scanInt()
	arr := make([]int, 0)
	recursion(0, arr)
}