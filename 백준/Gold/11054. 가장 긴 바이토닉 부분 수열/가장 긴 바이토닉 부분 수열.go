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
	var maxLen int
	N, _ := scanInt()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i], _ = scanInt()
	}
	lis := make([]int, N)
	lds := make([]int, N)

	//lis
	for i := 0; i < N; i++ {
		lis[i] = 1
		// i까지 연결 가능한 증가하는 부분 수열 찾기
		for j := 0; j < N; j++ {
			if A[j] < A[i] {
				lis[i] = max(lis[i], lis[j]+1)
			}
		}
	}

	//lds
	for i := N - 1; i >= 0; i-- {
		lds[i] = 1
		// i이후의 수열에서 감소하는 부분 수열 찾기
		for j := i + 1; j < N; j++ {
			if A[j] < A[i] {
				lds[i] = max(lds[i], lds[j]+1)
			}
		}
	}

	for i := 0; i < N; i++ {
		maxLen = max(maxLen, lis[i]+lds[i]-1)
	}
	w.WriteString(strconv.Itoa(maxLen))
	w.WriteByte('\n')
}
