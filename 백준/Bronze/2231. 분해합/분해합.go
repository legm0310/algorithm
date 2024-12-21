package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	for i := 1; i <= n; i++ {
		numString := strconv.Itoa(i)
		sum := i
		for _, num := range numString {
			sum += int(num - '0')
		}
		if sum == n {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(0)
}