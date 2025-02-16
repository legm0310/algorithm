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
	N, _ := scanInt()
	for i := 0; i < 9; i++ {
		w.WriteString(strconv.Itoa(N))
		w.WriteString(" * ")
		w.WriteString(strconv.Itoa(i + 1))
		w.WriteString(" = ")
		w.WriteString(strconv.Itoa((i + 1) * N))
		w.WriteByte('\n')
	}
}