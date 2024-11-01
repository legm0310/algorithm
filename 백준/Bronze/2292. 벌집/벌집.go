package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	tmp := 1
	fmt.Fscanf(reader, "%d\n", &n)
	for i := 1; ; i++ {
		if tmp >= n {
			fmt.Println(i)
			break
		}
		tmp += 6 * i
	}
}
