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

func recur(a, b, c int) int {
	if b == 0 {
		return 1
	}
	if b%2 == 0 {
		half := recur(a, b/2, c)
		return (half * half) % c
	} else {
		tmp := recur(a, b-1, c)
		return (a * tmp) % c
	}
}

func main() {
	defer w.Flush()
	a, _ := scanInt()
	b, _ := scanInt()
	c, _ := scanInt()

	w.WriteString(strconv.Itoa(recur(a, b, c)) + "\n")
}
