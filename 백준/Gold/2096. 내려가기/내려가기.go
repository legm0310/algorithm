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
	input := make([][]int, N)
	maxDp := make([][]int, N)
	minDp := make([][]int, N)
	for i, _ := range maxDp {
		input[i] = make([]int, 3)
		maxDp[i] = make([]int, 3)
		minDp[i] = make([]int, 3)
		input[i][0], _ = scanInt()
		input[i][1], _ = scanInt()
		input[i][2], _ = scanInt()
	}
	for idx, _ := range maxDp {
		if idx == 0 {
			maxDp[idx][0] = input[idx][0]
			maxDp[idx][1] = input[idx][1]
			maxDp[idx][2] = input[idx][2]
			minDp[idx][0] = input[idx][0]
			minDp[idx][1] = input[idx][1]
			minDp[idx][2] = input[idx][2]
		} else {
			maxDp[idx][0] = input[idx][0] + max(maxDp[idx-1][0], maxDp[idx-1][1])
			maxDp[idx][1] = input[idx][1] + max(max(maxDp[idx-1][0], maxDp[idx-1][1]), maxDp[idx-1][2])
			maxDp[idx][2] = input[idx][2] + max(maxDp[idx-1][1], maxDp[idx-1][2])
			minDp[idx][0] = input[idx][0] + min(minDp[idx-1][0], minDp[idx-1][1])
			minDp[idx][1] = input[idx][1] + min(min(minDp[idx-1][0], minDp[idx-1][1]), minDp[idx-1][2])
			minDp[idx][2] = input[idx][2] + min(minDp[idx-1][1], minDp[idx-1][2])
		}
	}
	w.WriteString(strconv.Itoa(max(max(maxDp[N-1][0], maxDp[N-1][1]), maxDp[N-1][2])))
	w.WriteString(" ")
	w.WriteString(strconv.Itoa(min(min(minDp[N-1][0], minDp[N-1][1]), minDp[N-1][2])))
	w.WriteByte('\n')
}
