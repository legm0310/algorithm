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
	k, _ := strconv.Atoi(parts[1])
	nums := make([]int, n)
	line, _ = r.ReadString('\n')
	parts = strings.Fields(line)
	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(parts[i])
		nums[i] = num
	}
	var tmp, result int
	for i := 0; i < n; i++ {
		if i == 0 {
			tmp = nums[i]
		} else {
			tmp += nums[i]
		}
		if i+1 == k {
			result = tmp
		}
		if i+1 > k {
			tmp -= nums[i-k]
		}
		if tmp > result {
			result = tmp
		}
	}
	w.WriteString(strconv.Itoa(result))
	w.WriteByte('\n')
	w.Flush()
}
