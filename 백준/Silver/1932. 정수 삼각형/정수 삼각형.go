package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	n     int
	input [][]int
	dp    [][]int
	s     = bufio.NewScanner(os.Stdin)
	w     = bufio.NewWriter(os.Stdout)
)

func main() {
	defer w.Flush()
	s.Split(bufio.ScanWords)
	n = scanInt()
	input = make([][]int, n)
	dp = make([][]int, n)
	for i, _ := range input {
		dp[i] = make([]int, i+1)
		input[i] = make([]int, i+1)
		for j, _ := range input[i] {
			dp[i][j] = -1
			input[i][j] = scanInt()
		}
	}
	w.WriteString(strconv.Itoa(find(0, 0)) + "\n")
}

func find(x, y int) int {
	if x == n-1 {
		return input[x][y]
	}
	if dp[x][y] != -1 {
		return dp[x][y]
	}
	dp[x][y] = max(find(x+1, y), find(x+1, y+1)) + input[x][y]
	return dp[x][y]
}

func scanInt() int {
	s.Scan()
	x, _ := strconv.Atoi(s.Text())
	return x
}
