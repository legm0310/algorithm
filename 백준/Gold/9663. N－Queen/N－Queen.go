package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	n     int
	board [][]int
	ans   = 0
	s     = bufio.NewScanner(os.Stdin)
	w     = bufio.NewWriter(os.Stdout)
)

func main() {
	defer w.Flush()
	n = scanInt()
	board = make([][]int, n)
	for i := 0; i < n; i++ {
		board[i] = make([]int, n)
	}
	find(0)
	w.WriteString(strconv.Itoa(ans) + "\n")
}

func find(row int) {
	if row == n {
		ans++
		return
	}

	for col := 0; col < n; col++ {
		if isSafe(row, col) {
			board[row][col] = 1
			find(row + 1)
			board[row][col] = 0
		}
	}
}

func isSafe(row, col int) bool {
	for i := 0; i < row; i++ {
		if board[i][col] == 1 {
			return false
		}
		if col-(row-i) >= 0 && board[i][col-(row-i)] == 1 {
			return false
		}
		if col+(row-i) < n && board[i][col+(row-i)] == 1 {
			return false
		}
	}
	return true
}

func scanInt() int {
	s.Scan()
	x, _ := strconv.Atoi(s.Text())
	return x
}
