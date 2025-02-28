package main

import (
	"bufio"
	"container/list"
	"os"
	"strconv"
)

var (
	N        int
	M        int
	virusMap [][]int
	dx       []int = []int{-1, 1, 0, 0}
	dy       []int = []int{0, 0, -1, 1}
	result   int
	s        *bufio.Scanner
	w        *bufio.Writer
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

func spreadVirus() int {
	queue := list.New()
	tmpMap := make([][]int, N)
	for i := 0; i < N; i++ {
		tmpMap[i] = make([]int, M)
		copy(tmpMap[i], virusMap[i])
	}
	for i, _ := range tmpMap {
		for j, n := range tmpMap[i] {
			if n == 2 {
				queue.PushBack([2]int{i, j})
			}
		}
	}

	for queue.Len() > 0 {
		target := queue.Remove(queue.Front()).([2]int)
		x, y := target[0], target[1]

		for d := 0; d < 4; d++ {
			nx, ny := x+dx[d], y+dy[d]

			if nx >= 0 && nx < N && ny >= 0 && ny < M && tmpMap[nx][ny] == 0 {
				tmpMap[nx][ny] = 2
				queue.PushBack([2]int{nx, ny})
			}
		}
	}

	sum := 0
	for i, _ := range tmpMap {
		for _, n := range tmpMap[i] {
			if n == 0 {
				sum++
			}
		}
	}
	return sum
}

func recursion(idx, count int) {
	if count == 3 {
		sum := spreadVirus()
		if sum > result {
			result = sum
		}
		return
	}

	for i := idx; i < N*M; i++ {
		x, y := i/M, i%M
		if virusMap[x][y] == 0 {
			virusMap[x][y] = 1
			recursion(i+1, count+1)
			virusMap[x][y] = 0
		}
	}

}

func main() {
	defer w.Flush()
	N, _ = scanInt()
	M, _ = scanInt()
	virusMap = make([][]int, N)
	for i := 0; i < N; i++ {
		virusMap[i] = make([]int, M)
		for j := 0; j < M; j++ {
			virusMap[i][j], _ = scanInt()
		}
	}
	recursion(0, 0)
	w.WriteString(strconv.Itoa(result))
	w.WriteByte('\n')
}
