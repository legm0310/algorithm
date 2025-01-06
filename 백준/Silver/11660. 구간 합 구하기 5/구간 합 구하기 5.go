package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	s *bufio.Scanner
	w *bufio.Writer
)

func init() {
	s = bufio.NewScanner(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
}

func scanValue() []int {
	s.Scan()
	v := strings.Split(s.Text(), " ")
	nums := make([]int, len(v))
	for i, str := range v {
		nums[i], _ = strconv.Atoi(str)
	}
	return nums
}

func main() {
	defer w.Flush()
	nums := scanValue()
	n := nums[0]
	m := nums[1]
	matrix := make([][]int, n)
	prefix := make([][]int, n+1)
	for i := 0; i < n; i++ {
		row := scanValue()
		matrix[i] = row
		prefix[i] = make([]int, n+1)
	}
	prefix[n] = make([]int, n+1)
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			prefix[x+1][y+1] = prefix[x][y+1] + prefix[x+1][y] - prefix[x][y] + matrix[x][y]
		}
	}
	for i := 0; i < m; i++ {
		q := scanValue()
		result := prefix[q[2]][q[3]] - prefix[q[2]][q[1]-1] - prefix[q[0]-1][q[3]] + prefix[q[0]-1][q[1]-1]
		w.WriteString(strconv.Itoa(result))
		w.WriteByte('\n')
	}
}
