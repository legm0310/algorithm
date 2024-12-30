package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	r *bufio.Reader
	w *bufio.Writer
)

func init() {
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
}

func maxNum(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func minNum(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	line, _ := r.ReadString('\n')
	parts := strings.Fields(line)
	wi, _ := strconv.Atoi(parts[1])
	nums := make([]int, wi)
	line, _ = r.ReadString('\n')
	parts = strings.Fields(line)
	for i, part := range parts {
		num, _ := strconv.Atoi(part)
		nums[i] = num
	}
	leftMax := make([]int, wi)
	leftMax[0] = nums[0]
	rightMax := make([]int, wi)
	rightMax[wi-1] = nums[wi-1]
	for i := 1; i < wi; i++ {
		leftMax[i] = maxNum(leftMax[i-1], nums[i])
	}
	for i := wi - 2; i >= 0; i-- {
		rightMax[i] = maxNum(rightMax[i+1], nums[i])
	}
	result := 0
	for i := 0; i < wi; i++ {
		sum := minNum(leftMax[i], rightMax[i]) - nums[i]
		if sum > 0 {
			result += sum
		}
	}
	w.WriteString(strconv.Itoa(result))
	w.WriteByte('\n')
	w.Flush()
}