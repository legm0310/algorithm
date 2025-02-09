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

func recursion(numArr *[]int) {
	if len(*numArr) == M {
		for _, num := range *numArr {
			w.WriteString(strconv.Itoa(num))
			w.WriteString(" ")
		}
		w.WriteByte('\n')
		return
	}

	for i := 0; i < N; i++ {
		*numArr = append(*numArr, input[i])
		recursion(numArr)
		*numArr = (*numArr)[:len(*numArr)-1]
	}
}

func main() {
	defer w.Flush()
	N, _ = scanInt()
	M, _ = scanInt()
	input = make([]int, N)
	for i := 0; i < N; i++ {
		input[i], _ = scanInt()
	}
	slices.Sort(input)
	arr := []int{}
	recursion(&arr)
}
