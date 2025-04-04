package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	n, p     int
	melodies [7][]int
	ans      = 0
	s        = bufio.NewScanner(os.Stdin)
	w        = bufio.NewWriter(os.Stdout)
)

func main() {
	defer w.Flush()
	s.Split(bufio.ScanWords)

	n, p = scanInt(), scanInt()
	for i := 0; i < n; i++ {
		line, fret := scanInt(), scanInt()
		stack := melodies[line]

		for len(stack) > 0 && stack[len(stack)-1] > fret {
			ans++
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 || stack[len(stack)-1] < fret {
			ans++
			stack = append(stack, fret)
		}

		melodies[line] = stack
	}
	fmt.Println(ans)
}

func scanInt() int {
	s.Scan()
	x, _ := strconv.Atoi(s.Text())
	return x
}
