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
}

func scanInt() int {
	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	return n
}

func main() {
	s.Split(bufio.ScanWords)
	m := make(map[int]int)
	n := scanInt()
	maxLength, minLength := 0, 0
	for i := 0; i < n; i++ {
		l := scanInt()
		h := scanInt()
		if i == 0 {
			maxLength, minLength = l, l
		}
		if maxLength < l {
			maxLength = l
		}
		if minLength > l {
			minLength = l
		}
		m[l] = h
	}
	length := maxLength - minLength + 1
	heights := make([]int, length)
	for i := minLength; i <= maxLength; i++ {
		if v, exists := m[i]; exists {
			heights[i-minLength] = v
		} else {
			heights[i-minLength] = 0
		}
	}
	leftMax, rightMax := make([]int, length), make([]int, length)
	leftMax[0] = heights[0]
	rightMax[length-1] = heights[length-1]
	for i := 1; i < length; i++ {
		leftMax[i] = max(leftMax[i-1], heights[i])
	}
	for i := length - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], heights[i])
	}
	sum := 0
	for i := 0; i < length; i++ {
		maxHeight := min(leftMax[i], rightMax[i])
		sum += maxHeight
	}
	w.WriteString(strconv.Itoa(sum))
	w.WriteByte('\n')
	w.Flush()
}
