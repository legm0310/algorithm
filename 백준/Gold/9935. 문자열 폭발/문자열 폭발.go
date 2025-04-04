package main

import (
	"bufio"
	"os"
	"strings"
)

var (
	w *bufio.Writer
	r *bufio.Reader
)

func init() {
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
}


func main() {
	defer w.Flush()
	str, _ := r.ReadString('\n')
	str = strings.TrimSpace(str)
	subStr, _ := r.ReadString('\n')
	subStr = strings.TrimSpace(subStr)

	if len(str) < len(subStr) {
		w.WriteString(str + "\n")
		return
	}

	stack := make([]byte, 0)
	for i := 0; i < len(str); i++ {
		stack = append(stack, str[i])
		if len(stack) >= len(subStr) && string(stack[len(stack)-len(subStr):]) == subStr {
			stack = stack[:len(stack)-len(subStr)]
		}
	}
	if len(stack) == 0 {
		w.WriteString("FRULA\n")
	} else {
		w.WriteString(string(stack) + "\n")
	}
}
