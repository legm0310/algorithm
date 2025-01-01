package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	var a, b int
	fmt.Fscanf(r, "%d %d", &a, &b)
	w.WriteString(strconv.Itoa(a + b))
	w.WriteByte('\n')
	w.WriteString(strconv.Itoa(a - b))
	w.WriteByte('\n')
	w.WriteString(strconv.Itoa(a * b))
	w.WriteByte('\n')
	w.WriteString(strconv.Itoa(a / b))
	w.WriteByte('\n')
	w.WriteString(strconv.Itoa(a % b))
	w.WriteByte('\n')
	w.Flush()
}
