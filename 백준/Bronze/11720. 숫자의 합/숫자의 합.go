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
	s.Scan()
	nums := s.Text()
	sum := 0
	for i := 0; i < n; i++ {
		sum += int(nums[i] - '0')
	}
	w.WriteString(strconv.Itoa(sum))
	w.WriteByte('\n')
}