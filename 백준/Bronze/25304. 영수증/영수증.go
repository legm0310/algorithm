package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	result string
	s      *bufio.Scanner
	w      *bufio.Writer
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
	X, _ := scanInt()
	N, _ := scanInt()
	for i := 0; i < N; i++ {
		amount, _ := scanInt()
		cnt, _ := scanInt()
		X -= amount * cnt
	}
	if X == 0 {
		result = "Yes"
	} else {
		result = "No"
	}
	w.WriteString(result)
	w.WriteByte('\n')
}