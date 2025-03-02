package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
)

var (
	n     int
	m     int
	input []int
	used  []bool
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

func recursion(arr *[]int) {
	if len(*arr) == m {
		for _, num := range *arr {
			w.WriteString(strconv.Itoa(num))
			w.WriteString(" ")
		}
		w.WriteByte('\n')
		return
	}
	usedCurDepth := make(map[int]bool)
	for i := 0; i < n; i++ {
		if used[i] || usedCurDepth[input[i]] {
			continue
		}
		used[i] = true
		usedCurDepth[input[i]] = true
		*arr = append(*arr, input[i])
		recursion(arr)
		used[i] = false
		*arr = (*arr)[:len(*arr)-1]
	}
}

func main() {
	defer w.Flush()
	n, _ = scanInt()
	m, _ = scanInt()
	input = make([]int, n)
	used = make([]bool, n)
	for i := range input {
		input[i], _ = scanInt()
	}
	slices.Sort(input)
	arr := make([]int, 0)
	recursion(&arr)
}
