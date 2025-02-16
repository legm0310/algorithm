package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	N         int
	I         [][]int
	MI        []int
	result    int = 1000000
	resultSeq []int
	s         *bufio.Scanner
	w         *bufio.Writer
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

func recursion(idx, P, F, S, V, C int, seq *[]int) {
	if MI[0] <= P && MI[1] <= F && MI[2] <= S && MI[3] <= V {
		if result > C {
			result = C
			resultSeq = append([]int{}, *seq...)
		}
		return
	}
	if idx == N {
		return
	}
	*seq = append(*seq, idx+1)
	recursion(idx+1, P+I[idx][0], F+I[idx][1], S+I[idx][2], V+I[idx][3], C+I[idx][4], seq)
	*seq = (*seq)[:len(*seq)-1]
	recursion(idx+1, P, F, S, V, C, seq)
}

func main() {
	defer w.Flush()
	N, _ = scanInt()
	MI = make([]int, 4)
	for i := 0; i < 4; i++ {
		MI[i], _ = scanInt()
	}
	I = make([][]int, N)
	for i := 0; i < N; i++ {
		I[i] = make([]int, 5)
		for j := 0; j < 5; j++ {
			I[i][j], _ = scanInt()
		}
	}

	seq := make([]int, 0)
	recursion(0, 0, 0, 0, 0, 0, &seq)

	if len(resultSeq) == 0 {
		w.WriteString("-1\n")
	} else {
		w.WriteString(strconv.Itoa(result))
		w.WriteByte('\n')
		for i := 0; i < len(resultSeq); i++ {
			w.WriteString(strconv.Itoa(resultSeq[i]))
			w.WriteString(" ")
		}
		w.WriteByte('\n')
	}
}
