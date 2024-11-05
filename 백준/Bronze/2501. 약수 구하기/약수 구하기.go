package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, seq int
	fmt.Fscanf(reader, "%d %d", &n, &seq)
	s := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			s = append(s, i)
		}
	}
	if len(s) >= seq {
		fmt.Println(s[seq-1])
	} else {
		fmt.Println(0)
	}
}