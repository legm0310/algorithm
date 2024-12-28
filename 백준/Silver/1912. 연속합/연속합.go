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

func main() {
	line, _ := r.ReadString('\n')
	parts := strings.Fields(line)
	n, _ := strconv.Atoi(parts[0])
	nums := make([]int, n)
	line, _ = r.ReadString('\n')
	parts = strings.Fields(line)
	for i := 0; i < n; i++ {
		nums[i], _ = strconv.Atoi(parts[i])
	}
	tmp, result := nums[0], nums[0]
	for i := 1; i < n; i++ {
		if tmp+nums[i] < nums[i] {
			tmp = nums[i]
		} else {
			tmp += nums[i]
		}
		if result < tmp {
			result = tmp
		}
	}
	w.WriteString(strconv.Itoa(result))
	w.WriteByte('\n')
	w.Flush()
}
