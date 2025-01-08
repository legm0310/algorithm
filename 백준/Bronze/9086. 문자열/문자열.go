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
}

func scanInt() (int, error) {
	s.Scan()
	return strconv.Atoi(s.Text())
}

func main() {
	defer w.Flush()
	n, _ := scanInt()
	for i := 0; i < n; i++ {
		s.Scan()
		str := s.Text()
		w.WriteString(string(str[0]) + string(str[len(str)-1]))
		w.WriteString("\n")
	}
}