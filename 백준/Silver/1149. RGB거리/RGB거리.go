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

//var (
//	s *bufio.Scanner
//	w *bufio.Writer
//)

func init() {
	s = bufio.NewScanner(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
	s.Split(bufio.ScanWords)
}

func scanInt() (int, error) {
	s.Scan()
	return strconv.Atoi(s.Text())
}

//	func scanValue() []int {
//		s.Scan()
//		v := strings.Split(s.Text(), " ")
//		nums := make([]int, len(v))
//		for i, str := range v {
//			nums[i], _ = strconv.Atoi(str)
//		}
//		return nums
//	}

func main() {
	defer w.Flush()
	N, _ := scanInt()
	input := make([][]int, N)
	dp := make([][]int, N)
	for i, _ := range dp {
		input[i] = make([]int, 3)
		dp[i] = make([]int, 3)
		input[i][0], _ = scanInt()
		input[i][1], _ = scanInt()
		input[i][2], _ = scanInt()
	}

	for idx, _ := range dp {
		if idx == 0 {
			dp[0][0] = input[0][0]
			dp[0][1] = input[0][1]
			dp[0][2] = input[0][2]
			continue
		}
		for rgb, _ := range dp[idx] {
			if rgb == 0 {
				dp[idx][rgb] = input[idx][rgb] + min(dp[idx-1][rgb+1], dp[idx-1][rgb+2])
			}
			if rgb == 1 {
				dp[idx][rgb] = input[idx][rgb] + min(dp[idx-1][rgb-1], dp[idx-1][rgb+1])
			}
			if rgb == 2 {
				dp[idx][rgb] = input[idx][rgb] + min(dp[idx-1][rgb-1], dp[idx-1][rgb-2])
			}
		}
	}
	w.WriteString(strconv.Itoa(min(min(dp[N-1][0], dp[N-1][1]), dp[N-1][2])))
	w.WriteByte('\n')
}
