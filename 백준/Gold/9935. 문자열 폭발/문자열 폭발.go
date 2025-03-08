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

	stack := make([]rune, 0)
	subRunes := []rune(subStr)
	for _, v := range str {
		stack = append(stack, v)
	    for len(stack) >= len(subRunes) {
			if string(stack[len(stack)-len(subRunes):]) == string(subRunes) {
				stack = stack[:len(stack)-len(subRunes)]
			} else {
				break
			}
		}
	}
	if len(stack) == 0 {
		w.WriteString("FRULA\n")
	} else {
		w.WriteString(string(stack) + "\n")
	}
}
