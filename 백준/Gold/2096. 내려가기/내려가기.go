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
		for j := 0; j < 3; j++ {
			if idx == 0 {
				maxDp[idx][j] = input[idx][j]
				minDp[idx][j] = input[idx][j]
			} else {
				if j == 0 {
					maxDp[idx][j] = input[idx][j] + max(maxDp[idx-1][j], maxDp[idx-1][j+1])
					minDp[idx][j] = input[idx][j] + min(minDp[idx-1][j], minDp[idx-1][j+1])
				}
				if j == 1 {
					maxDp[idx][j] = input[idx][j] + max(max(maxDp[idx-1][j-1], maxDp[idx-1][j]), maxDp[idx-1][j+1])
					minDp[idx][j] = input[idx][j] + min(min(minDp[idx-1][j-1], minDp[idx-1][j]), minDp[idx-1][j+1])
				}
				if j == 2 {
					maxDp[idx][j] = input[idx][j] + max(maxDp[idx-1][j-1], maxDp[idx-1][j])
					minDp[idx][j] = input[idx][j] + min(minDp[idx-1][j-1], minDp[idx-1][j])
				}
			}

		}
	}
	w.WriteString(strconv.Itoa(max(max(maxDp[N-1][0], maxDp[N-1][1]), maxDp[N-1][2])))
	w.WriteString(" ")
	w.WriteString(strconv.Itoa(min(min(minDp[N-1][0], minDp[N-1][1]), minDp[N-1][2])))
	w.WriteByte('\n')
}
