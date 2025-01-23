package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
)

var (
	N     int
	M     int
	arr   []int
	input []int
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

func recursion(num int, numArr *[]int) {
	if num == M {
		for _, v := range *numArr {
			w.WriteString(strconv.Itoa(v))
			w.WriteString(" ")
		}
		w.WriteByte('\n')
		return
	}

	for i := 0; i < N; i++ {
		if slices.Contains(*numArr, input[i]) {
			continue
		}
		*numArr = append(*numArr, input[i])
		recursion(num+1, numArr)
		*numArr = (*numArr)[:len(*numArr)-1]
	}
}

func main() {
	defer w.Flush()
	N, _ = scanInt()
	M, _ = scanInt()
	arr = make([]int, 0)
	input = make([]int, N)
	for i := 0; i < N; i++ {
		input[i], _ = scanInt()
	}
	slices.Sort(input)
	recursion(0, &arr)
}
