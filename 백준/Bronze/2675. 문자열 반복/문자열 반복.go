package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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
	for i := 0; i < n; i++ {
		cnt, _ := scanInt()
		s.Scan()
		str := s.Text()
		for j := 0; j < len(str); j++ {
			w.WriteString(strings.Repeat(string(str[j]), cnt))
		}
		w.WriteByte('\n')
	}
}