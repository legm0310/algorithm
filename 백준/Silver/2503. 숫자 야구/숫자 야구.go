package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	N     int
	hints [][]int
	ans   int = 0
	s     *bufio.Scanner
	w     *bufio.Writer
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

func check(idx int, num int) bool {
	hintStr := strconv.Itoa(hints[idx][0])
	numStr := strconv.Itoa(num)
	if numStr[0] == numStr[1] || numStr[0] == numStr[2] || numStr[1] == numStr[2] || strings.Contains(numStr, "0") {
		return false
	}
	strike, ball := 0, 0
	for i := 0; i < 3; i++ {
		if numStr[i] == hintStr[i] {
			strike++
		} else if strings.Contains(hintStr, string(numStr[i])) {
			ball++
		}
	}
	if strike == hints[idx][1] && ball == hints[idx][2] {
		return true
	} else {
		return false
	}

}

func recursion(idx int, num int) {
	if num == 1000 {
		return
	}
	if idx == N {
		ans++
		recursion(0, num+1)
		return
	}
	if check(idx, num) {
		recursion(idx+1, num)
	} else {
		recursion(0, num+1)
	}

}

func main() {
	defer w.Flush()
	N, _ = scanInt()
	hints = make([][]int, N)
	for i := 0; i < N; i++ {
		num, _ := scanInt()
		strike, _ := scanInt()
		ball, _ := scanInt()
		hints[i] = []int{num, strike, ball}
	}
	recursion(0, 123)
	w.WriteString(strconv.Itoa(ans))
	w.WriteByte('\n')
}