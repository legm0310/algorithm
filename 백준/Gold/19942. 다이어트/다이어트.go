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

func recursion(idx int, seqArr *[]int, ingrArr *[]int) {
	if MI[0] <= (*ingrArr)[0] && MI[1] <= (*ingrArr)[1] && MI[2] <= (*ingrArr)[2] && MI[3] <= (*ingrArr)[3] {
		if (*ingrArr)[4] < result {
			resultSeq = append([]int{}, *seqArr...)
			result = (*ingrArr)[4]
		}
		return
	}
	for i := idx; i < N; i++ {
		*seqArr = append(*seqArr, i+1)
		for j := 0; j < 5; j++ {
			(*ingrArr)[j] += I[i][j]
		}
		recursion(i+1, seqArr, ingrArr)
		*seqArr = (*seqArr)[:len(*seqArr)-1]
		for j := 0; j < 5; j++ {
			(*ingrArr)[j] -= I[i][j]
		}
	}
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
	seqArr := make([]int, 0)
	ingrArr := make([]int, 5)
	recursion(0, &seqArr, &ingrArr)
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
